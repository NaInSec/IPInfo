package main

import (
	"compress/gzip"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/ipinfo/cli/lib/complete"
	"github.com/ipinfo/cli/lib/complete/predict"
	"github.com/spf13/pflag"
)

const dbDownloadURL = "https://ipinfo.io/data/free/"

type ChecksumResponse struct {
	Checksums struct {
		MD5    string `json:"md5"`
		SHA1   string `json:"sha1"`
		SHA256 string `json:"sha256"`
	} `json:"checksums"`
}

var completionsDownload = &complete.Command{
	Flags: map[string]complete.Predictor{
		"-c":         predict.Nothing,
		"--compress": predict.Nothing,
		"-f":         predict.Nothing,
		"--format":   predict.Nothing,
		"-t":         predict.Nothing,
		"--token":    predict.Nothing,
		"-h":         predict.Nothing,
		"--help":     predict.Nothing,
	},
}

func printHelpDownload() {
	fmt.Printf(
		`Usage: %s download [<opts>] <database> [<output>]

Description:
    Download the free ipinfo databases.

Examples:
    # Download country database in csv format.
    $ %[1]s download country -f csv > country.csv
    $ %[1]s download country-asn country_asn.mmdb

Databases:
    asn            free ipinfo asn database.
    country        free ipinfo country database.
    country-asn    free ipinfo country-asn database.

Options:
  General:
    --token <tok>, -t <tok>
      use <tok> as API token.
    --help, -h
      show help.

Outputs:
    --compress, -c
      save the file in compressed format.
      default: false.
    --format, -f <mmdb | json | csv>
      output format of the database file.
      default: mmdb.
`, progBase)
}

func cmdDownload() error {
	var fTok string
	var fFmt string
	var fZip bool
	var fHelp bool

	pflag.StringVarP(&fTok, "token", "t", "", "the token to use.")
	pflag.StringVarP(&fFmt, "format", "f", "mmdb", "the output format to use.")
	pflag.BoolVarP(&fZip, "compress", "c", false, "compressed output.")
	pflag.BoolVarP(&fHelp, "help", "h", false, "show help.")
	pflag.Parse()

	args := pflag.Args()[1:]
	if fHelp || len(args) > 2 || len(args) < 1 {
		printHelpDownload()
		return nil
	}

	token := fTok
	if token == "" {
		token = gConfig.Token
	}

	// require token for download.
	if token == "" {
		return errors.New("downloading requires a token; login via `ipinfo init` or pass the `--token` argument")
	}

	// get download format and extension.
	var format string
	var fileExtension string
	switch strings.ToLower(fFmt) {
	case "mmdb":
		format = "mmdb"
		fileExtension = "mmdb"
	case "csv":
		format = "csv.gz"
		fileExtension = "csv"
	case "json":
		format = "json.gz"
		fileExtension = "json"
	default:
		return errors.New("unknown download format")
	}

	if fZip {
		fileExtension = fmt.Sprintf("%s.%s", fileExtension, "gz")
	}

	// download the db.
	var dbName string
	switch strings.ToLower(args[0]) {
	case "asn":
		dbName = "asn"
	case "country":
		dbName = "country"
	case "country-asn":
		dbName = "country_asn"
	default:
		return fmt.Errorf("database '%v' is invalid", args[0])
	}

	// get file name.
	var fileName string
	if len(pflag.Args()) > 2 {
		fileName = pflag.Args()[2]
	} else {
		fileName = fmt.Sprintf("%s.%s", dbName, fileExtension)
	}

	url := fmt.Sprintf("%s%s.%s?token=%s", dbDownloadURL, dbName, format, token)
	err := downloadDb(url, fileName, format, fZip)
	if err != nil {
		return err
	}

	// fetch checksums from API and check if they match.
	checksumUrl := fmt.Sprintf("%s%s.%s/checksums?token=%s", dbDownloadURL, dbName, format, token)
	checksumResponse, err := fetchChecksums(checksumUrl)
	if err != nil {
		return err
	}

	// compute checksum of downloaded file.
	localChecksum, err := computeSHA256(fileName)
	if err != nil {
		return err
	}

	// compare checksums.
	if localChecksum != checksumResponse.Checksums.SHA256 {
		return errors.New("checksums do not match. File might be corrupted")
	}

	return nil
}

func downloadDb(url string, fileName string, format string, zip bool) error {

	// make API req to download the file.
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// if output not terminal write to stdout.
	if fileInfo, _ := os.Stdout.Stat(); (fileInfo.Mode() & os.ModeCharDevice) == 0 {
		if zip {
			err := zipWriter(os.Stdout, res.Body)
			if err != nil {
				return err
			}
		} else {
			err := unzipWrite(os.Stdout, res.Body)
			if err != nil {
				return err
			}
		}

	} else {
		// create file.
		file, err := os.Create(fileName)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		// save compressed file.
		if zip {
			if format == "mmdb" {
				err := zipWriter(file, res.Body)
				if err != nil {
					return err
				}
			} else {
				_, err = io.Copy(file, res.Body)
				if err != nil {
					return err
				}
			}
		} else {
			if format == "mmdb" {
				_, err = io.Copy(file, res.Body)
				if err != nil {
					return err
				}
			} else {
				err := unzipWrite(file, res.Body)
				if err != nil {
					return err
				}
			}
		}

		fmt.Printf("Database %s saved successfully.\n", fileName)
	}

	return nil
}

func zipWriter(file *os.File, data io.Reader) error {
	writer := gzip.NewWriter(file)
	defer writer.Close()

	body, err := io.ReadAll(data)
	if err != nil {
		return err
	}

	_, err = writer.Write(body)
	if err != nil {
		return err
	}

	return nil
}

func unzipWrite(file *os.File, data io.Reader) error {
	unzipData, err := gzip.NewReader(data)
	if err != nil {
		return err
	}
	defer unzipData.Close()

	_, err = io.Copy(file, unzipData)
	if err != nil {
		return err
	}

	return nil
}

func computeSHA256(filepath string) (string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hasher := sha256.New()
	if _, err := io.Copy(hasher, file); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hasher.Sum(nil)), nil
}

func fetchChecksums(url string) (*ChecksumResponse, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var checksumResponse ChecksumResponse
	if err := json.Unmarshal(body, &checksumResponse); err != nil {
		return nil, err
	}

	return &checksumResponse, nil
}
