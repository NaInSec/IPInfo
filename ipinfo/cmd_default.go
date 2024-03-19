package main

import (
	"errors"
	"fmt"
	"net"
	"os"

	"github.com/fatih/color"
	"github.com/ipinfo/cli/lib"
	"github.com/ipinfo/cli/lib/iputil"
	"github.com/ipinfo/go/v2/ipinfo"
	"github.com/spf13/pflag"
)

func printHelpDefault() {
	fmt.Printf(
		`Usage: %s <cmd> [<opts>] [<args>]

Commands:
  <ip>        look up details for an IP address, e.g. 8.8.8.8.
  <asn>       look up details for an ASN, e.g. AS123 or as123.
  myip        get details for your IP.
  bulk        get details for multiple IPs in bulk.
  asn         tools related to ASNs.
  summarize   get summarized data for a group of IPs.
  map         open a URL to a map showing the locations of a group of IPs.
  prips       print IP list from CIDR or range.
  grepip      grep for IPs matching criteria from any source.
  matchip     print the overlapping IPs and subnets.
  grepdomain  grep for domains matching criteria from any source.
  cidr2range  convert CIDRs to IP ranges.
  cidr2ip     convert CIDRs to individual IPs within those CIDRs.
  range2cidr  convert IP ranges to CIDRs.
  range2ip    convert IP ranges to individual IPs within those ranges.
  randip      Generates random IPs.
  splitcidr   splits a larger CIDR into smaller CIDRs.
  mmdb        read, import and export mmdb files.
  calc 	      evaluates a mathematical expression that may contain IP addresses.
  tool        misc. tools related to IPs, IP ranges and CIDRs.
  download    download free ipinfo database files.
  cache       manage the cache.
  config      manage the config.
  quota       print the request quota of your account.
  init        login or signup account.
  logout      delete your current API token session.
  completion  install or output shell auto-completion script.
  version     show current version.

Options:
  General:
    --token <tok>, -t <tok>
      use <tok> as API token.
    --nocache
      do not use the cache.
    --version, -v
      show binary release number.
    --help, -h
      show help.

  Outputs:
    --field <field>, -f <field>
      lookup only specific fields in the output.
      field names correspond to JSON keys, e.g. 'hostname' or 'company.type'.
      multiple field names must be separated by commas.
    --nocolor
      disable colored output.

  Formats:
    --pretty, -p
      output pretty format.
    --json, -j
      output JSON format.
    --csv, -c
      output CSV format.
    --yaml, -y
      output YAML format.
`, progBase)
}

var detailedHelp = fmt.Sprintf(`Usage: %s <cmd> [<opts>] [<args>]

Commands:
  <ip>        look up details for an IP address, e.g. 8.8.8.8.
  <asn>       look up details for an ASN, e.g. AS123 or as123.
  myip        get details for your IP.
  bulk        get details for multiple IPs in bulk.
  asn         tools related to ASNs.
  summarize   get summarized data for a group of IPs.
  map         open a URL to a map showing the locations of a group of IPs.
  prips       print IP list from CIDR or range.
  grepip      grep for IPs matching criteria from any source.
  matchip     print the overlapping IPs and subnets.
  grepdomain  grep for domains matching criteria from any source.
  cidr2range  convert CIDRs to IP ranges.
  cidr2ip     convert CIDRs to individual IPs within those CIDRs.
  range2cidr  convert IP ranges to CIDRs.
  range2ip    convert IP ranges to individual IPs within those ranges.
  randip      Generates random IPs.
  splitcidr   splits a larger CIDR into smaller CIDRs.
  mmdb        read, import and export mmdb files.
  calc 	      evaluates a mathematical expression that may contain IP addresses.
  tool        misc. tools related to IPs, IP ranges and CIDRs.
  download    download free ipinfo database files.
  cache       manage the cache.
  config      manage the config.
  quota       print the request quota of your account.
  init        login or signup account.
  logout      delete your current API token session.
  completion  install or output shell auto-completion script.
  version     show current version.

Options:
  General:
    --token <tok>, -t <tok>
      use <tok> as API token.
    --nocache
      do not use the cache.
    --version, -v
      show binary release number.
    --help, -h
      show help.

  Outputs:
    --field <field>, -f <field>
      lookup only specific fields in the output.
      field names correspond to JSON keys, e.g. 'hostname' or 'company.type'.
      multiple field names must be separated by commas.
    --nocolor
      disable colored output.

  Formats:
    --pretty, -p
      output pretty format.
    --json, -j
      output JSON format.
    --csv, -c
      output CSV format.
    --yaml, -y
      output YAML format.
`, progBase)

func cmdDefault() (err error) {
	var ips []net.IP
	var fTok string
	var fVsn bool
	var fField []string
	var fPretty bool
	var fJSON bool
	var fCSV bool
	var fYAML bool

	pflag.StringVarP(&fTok, "token", "t", "", "the token to use.")
	pflag.BoolVar(&fNoCache, "nocache", false, "disable the cache.")
	pflag.BoolVarP(&fVsn, "version", "v", false, "print binary release number.")
	pflag.BoolVarP(&fHelp, "", "h", false, "show help.")
	pflag.BoolVar(&fHelpDetailed, "help", false, "show detailed help")
	pflag.StringSliceVarP(&fField, "field", "f", nil, "specific field to lookup.")
	pflag.BoolVarP(&fPretty, "pretty", "p", true, "output pretty format.")
	pflag.BoolVarP(&fJSON, "json", "j", true, "output JSON format. (default)")
	pflag.BoolVarP(&fCSV, "csv", "c", false, "output CSV format.")
	pflag.BoolVarP(&fYAML, "yaml", "y", false, "output YAML format.")
	pflag.BoolVar(&fNoColor, "nocolor", false, "disable colored output.")
	pflag.Parse()

	if fNoColor {
		color.NoColor = true
	}

	if fHelp {
		printHelpDefault()
		return nil
	}

	if fHelpDetailed {
		// Read the string and display it using a pager
		lib.HelpDetailed(detailedHelp, printHelpDefault)
		return nil
	}

	if fVsn {
		fmt.Println(version)
		return nil
	}

	args := pflag.Args()
	if len(args) != 0 && args[0] != "-" {
		fmt.Printf("err: \"%s\" is not a command.\n", os.Args[1])
		fmt.Println()
		printHelpDefault()
		return nil
	}

	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) != 0 {
		printHelpDefault()
		return nil
	}

	ips = iputil.IPListFromStdin()
	if len(ips) == 0 {
		fmt.Println("no input ips")
		return nil
	}

	ii = prepareIpinfoClient(fTok)

	// require token for bulk.
	if ii.Token == "" {
		return errors.New("bulk lookups require a token; login via `ipinfo init`.")
	}

	data, err := ii.GetIPInfoBatch(ips, ipinfo.BatchReqOpts{
		TimeoutPerBatch:              60 * 30, // 30min
		ConcurrentBatchRequestsLimit: 20,
	})
	if err != nil {
		return err
	}

	if len(fField) > 0 {
		return outputFieldBatchCore(data, fField, true, true)
	}

	if fCSV {
		return outputCSVBatchCore(data)
	}
	if fYAML {
		return outputYAML(data)
	}

	return outputJSON(data)
}
