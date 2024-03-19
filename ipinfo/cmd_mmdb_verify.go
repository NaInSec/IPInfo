package main

import (
	"fmt"

	"github.com/ipinfo/cli/lib/complete"
	"github.com/ipinfo/cli/lib/complete/predict"
	"github.com/ipinfo/mmdbctl/lib"
	"github.com/spf13/pflag"
)

var completionsMmdbVerify = &complete.Command{
	Flags: map[string]complete.Predictor{
		"-h":     predict.Nothing,
		"--help": predict.Nothing,
	},
}

func printHelpMmdbVerify() {
	fmt.Printf(
		`Usage: %s mmdb verify [<opts>] <mmdb_file>

Options:
  General:
    --help, -h
      show help.
`, progBase)
}

func cmdMmdbVerify() error {
	f := lib.CmdVerifyFlags{}
	f.Init()
	pflag.Parse()
	if pflag.NArg() <= 2 && pflag.NFlag() == 0 {
		f.Help = true
	}

	return lib.CmdVerify(f, pflag.Args()[2:], printHelpMmdbVerify)
}
