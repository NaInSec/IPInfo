package main

import (
	"fmt"

	"github.com/ipinfo/cli/lib/complete"
	"github.com/ipinfo/cli/lib/complete/predict"
	"github.com/ipinfo/mmdbctl/lib"
	"github.com/spf13/pflag"
)

var predictMetadataFmts = []string{"pretty", "json"}

var completionsMmdbMetadata = &complete.Command{
	Flags: map[string]complete.Predictor{
		"--nocolor": predict.Nothing,
		"-h":        predict.Nothing,
		"--help":    predict.Nothing,
		"-f":        predict.Set(predictMetadataFmts),
		"--format":  predict.Set(predictMetadataFmts),
	},
}

func printHelpMmdbMetadata() {
	fmt.Printf(
		`Usage: %s mmdb metadata [<opts>] <mmdb_file>

Options:
  General:
    --nocolor
      disable colored output.
    --help, -h
      show help.

  Format:
    -f <format>, --format <format>
      the metadata output format.
      can be "pretty" or "json".
      default: pretty.
`, progBase)
}

func cmdMmdbMetadata() error {
	f := lib.CmdMetadataFlags{}
	f.Init()
	pflag.Parse()
	if pflag.NArg() <= 2 && pflag.NFlag() == 0 {
		f.Help = true
	}

	return lib.CmdMetadata(f, pflag.Args()[2:], printHelpMmdbMetadata)
}
