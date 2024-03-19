package lib

import (
	"errors"
	"strings"

	"github.com/ipinfo/cli/lib/iputil"
	"github.com/ipinfo/go/v2/ipinfo"
	"github.com/spf13/pflag"
)

// CmdASNBulkFlags are flags expected by CmdASNBulk
type CmdASNBulkFlags struct {
	Token   string
	nocache bool
	help    bool
	Field   []string
	json    bool
	Yaml    bool
}

// Init initializes the common flags available to CmdASNBulk with sensible
// defaults.
// pflag.Parse() must be called to actually use the final flag values.
func (f *CmdASNBulkFlags) Init() {
	_h := "see description in --help"
	pflag.StringVarP(
		&f.Token,
		"token", "t", "",
		_h,
	)
	pflag.BoolVarP(
		&f.nocache,
		"nocache", "", false,
		_h,
	)
	pflag.BoolVarP(
		&f.help,
		"help", "h", false,
		_h,
	)
	pflag.StringSliceVarP(
		&f.Field,
		"field", "f", []string{},
		_h,
	)
	pflag.BoolVarP(
		&f.json,
		"json", "j", false,
		_h,
	)
	pflag.BoolVarP(
		&f.Yaml,
		"yaml", "y", false,
		_h,
	)
}

// CmdASNBulk is the entrypoint for the `ipinfo asn-bulk` command.
func CmdASNBulk(f CmdASNBulkFlags, ii *ipinfo.Client, args []string, printHelp func()) (ipinfo.BatchASNDetails, error) {
	if f.help {
		printHelp()
		return nil, nil
	}

	var asns []string

	op := func(string string, inputType iputil.INPUT_TYPE) error {
		switch inputType {
		case iputil.INPUT_TYPE_ASN:
			asns = append(asns, strings.ToUpper(string))
		default:
			return iputil.ErrInvalidInput
		}
		return nil
	}
	err := iputil.GetInputFrom(args, true, true, op)
	if err != nil {
		return nil, err
	}

	if ii.Token == "" {
		return nil, errors.New("bulk lookups require a token; login via `ipinfo init`")
	}

	return ii.GetASNDetailsBatch(asns, ipinfo.BatchReqOpts{
		TimeoutPerBatch:              60 * 30, // 30min
		ConcurrentBatchRequestsLimit: 20,
	})
}
