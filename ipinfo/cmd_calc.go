package main

import (
	"fmt"
	"github.com/ipinfo/cli/lib"
	"github.com/ipinfo/cli/lib/complete"
	"github.com/ipinfo/cli/lib/complete/predict"
	"github.com/spf13/pflag"
)

var completionsCalc = &complete.Command{
	Flags: map[string]complete.Predictor{
		"-h":     predict.Nothing,
		"--help": predict.Nothing,
	},
}

// printHelpCalc prints the help message for the "calc" command.
func printHelpCalc() {
	fmt.Printf(
		`Usage: %s calc <expression> [<opts>]

Description:
  Evaluate a mathematical expression and print the result.

Examples:
  %[1]s calc "2*2828-1"
  %[1]s calc "190.87.89.1*2"
  %[1]s calc "2001:0db8:85a3:0000:0000:8a2e:0370:7334*6"

Options:
  General:
    --help, -h
      show help.
`, progBase)
}

// cmdCalc is the handler for the "calc" command.
func cmdCalc() error {
	f := lib.CmdCalcFlags{}
	f.Init()
	pflag.Parse()

	return lib.CmdCalc(f, pflag.Args()[1:], printHelpCalc)
}
