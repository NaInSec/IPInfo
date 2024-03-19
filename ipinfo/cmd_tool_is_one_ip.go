package main

import (
	"fmt"

	"github.com/ipinfo/cli/lib"
	"github.com/ipinfo/cli/lib/complete"
	"github.com/ipinfo/cli/lib/complete/predict"
	"github.com/spf13/pflag"
)

var completionsToolIsOneIp = &complete.Command{
	Flags: map[string]complete.Predictor{
		"-h":     predict.Nothing,
		"--help": predict.Nothing,
	},
}

func printHelpToolIsOneIp() {
	fmt.Printf(
		`Usage: %s tool is_one_ip [<opts>] <cidr | ip | ip-range | filepath>

Description:
  Checks whether a CIDR or IP Range contains exactly one IP.

Examples:
  # Check CIDR.
  $ %[1]s tool is_one_ip 1.1.1.0/30

  # Check IP.
  $ %[1]s tool is_one_ip 1.1.1.1

  # Check IP range.
  $ %[1]s tool is_one_ip 1.1.1.1-2.2.2.2
  
  # Check for file.
  $ %[1]s tool is_one_ip /path/to/file.txt 

  # Check entries from stdin.
  $ cat /path/to/file1.txt | %[1]s tool is_one_ip

Options:
  --help, -h
    show help.
`, progBase)
}

func cmdToolIsOneIp() (err error) {
	f := lib.CmdToolIsOneIpFlags{}
	f.Init()
	pflag.Parse()

	return lib.CmdToolIsOneIp(f, pflag.Args()[2:], printHelpToolIsOneIp)
}
