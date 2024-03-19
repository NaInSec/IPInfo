package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/fatih/color"
	"github.com/ipinfo/cli/lib/complete"
	"github.com/ipinfo/cli/lib/complete/predict"
	"github.com/ipinfo/go/v2/ipinfo"
	"github.com/spf13/pflag"
)

var completionsASNSingle = &complete.Command{
	Flags: map[string]complete.Predictor{
		"-t":        predict.Nothing,
		"--token":   predict.Nothing,
		"--nocache": predict.Nothing,
		"-h":        predict.Nothing,
		"--help":    predict.Nothing,
		"-f":        predict.Set(asnFields),
		"--field":   predict.Set(asnFields),
		"--nocolor": predict.Nothing,
		"-p":        predict.Nothing,
		"--pretty":  predict.Nothing,
		"-j":        predict.Nothing,
		"--json":    predict.Nothing,
	},
}

func printHelpASNSingle(asn string) {
	fmt.Printf(
		`Usage: %s %s [<opts>]

Options:
  General:
    --token <tok>, -t <tok>
      use <tok> as API token.
    --nocache
      do not use the cache.
    --help, -h
      show help.

  Outputs:
    --field <field>, -f <field>
      lookup only specific fields in the output.
      field names correspond to JSON keys, e.g. 'registry' or 'allocated'.
      multiple field names must be separated by commas.
    --nocolor
      disable colored output.

  Formats:
    --json, -j
      output JSON format. (default)
    --yaml, -y
      output YAML format.
`, progBase, asn)
}

func cmdASNSingle(asn string) error {
	var fTok string
	var fField []string
	var fJSON bool
	var fYAML bool

	pflag.StringVarP(&fTok, "token", "t", "", "the token to use.")
	pflag.BoolVar(&fNoCache, "nocache", false, "disable the cache.")
	pflag.BoolVarP(&fHelp, "help", "h", false, "show help.")
	pflag.StringSliceVarP(&fField, "field", "f", nil, "specific field to lookup.")
	pflag.BoolVarP(&fJSON, "json", "j", true, "output JSON format. (default)")
	pflag.BoolVarP(&fYAML, "yaml", "y", false, "output YAML format.")
	pflag.BoolVar(&fNoColor, "nocolor", false, "disable color output.")
	pflag.Parse()

	if fNoColor {
		color.NoColor = true
	}

	if fHelp {
		printHelpASNSingle(asn)
		return nil
	}

	ii = prepareIpinfoClient(fTok)

	// require token for ASN API.
	if ii.Token == "" {
		return errors.New("ASN lookups require a token; login via `ipinfo init`.")
	}

	data, err := ii.GetASNDetails(asn)
	if err != nil {
		iiErr, ok := err.(*ipinfo.ErrorResponse)
		if ok && (iiErr.Response.StatusCode == http.StatusUnauthorized) {
			return errors.New("Token does not have access to ASN API")
		}
		return err
	}

	if len(fField) > 0 {
		d := make(ipinfo.BatchASNDetails, 1)
		d[data.ASN] = data
		return outputFieldBatchASNDetails(d, fField, false, false)
	}
	if fYAML {
		return outputYAML(data)
	}

	return outputJSON(data)
}
