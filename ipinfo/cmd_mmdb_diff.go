package main

import (
	"fmt"

	"github.com/ipinfo/cli/lib/complete"
	"github.com/ipinfo/cli/lib/complete/predict"
	"github.com/ipinfo/mmdbctl/lib"
	"github.com/spf13/pflag"
)

var completionsMmdbDiff = &complete.Command{
	Flags: map[string]complete.Predictor{
		"-h":        predict.Nothing,
		"--help":    predict.Nothing,
		"-s":        predict.Nothing,
		"--subnets": predict.Nothing,
		"-r":        predict.Nothing,
		"--records": predict.Nothing,
	},
}

func printHelpMmdbDiff() {
	fmt.Printf(
		`Usage: %s mmdb diff [<opts>] <old> <new>

Description:
  Print subnet and record differences between two mmdb files (i.e. do set
  difference `+"`"+"(new - old) U (old - new)"+"`"+`).

Options:
  General:
    --help, -h
      show help.
    --subnets, -s
      show subnets difference.
    --records, -r
      show records difference.
`, progBase)
}

func cmdMmdbDiff() error {
	f := lib.CmdDiffFlags{}
	f.Init()
	pflag.Parse()
	if pflag.NArg() <= 2 && pflag.NFlag() == 0 {
		f.Help = true
	}

	return lib.CmdDiff(f, pflag.Args()[2:], printHelpMmdbDiff)
}
