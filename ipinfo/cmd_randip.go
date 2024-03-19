package main

import (
	"fmt"

	"github.com/ipinfo/cli/lib"
	"github.com/ipinfo/cli/lib/complete"
	"github.com/ipinfo/cli/lib/complete/predict"
	"github.com/spf13/pflag"
)

var completionsRandIP = &complete.Command{
	Flags: map[string]complete.Predictor{
		"-h":                 predict.Nothing,
		"--help":             predict.Nothing,
		"-n":                 predict.Nothing,
		"--num":              predict.Nothing,
		"-4":                 predict.Nothing,
		"--ipv4":             predict.Nothing,
		"-6":                 predict.Nothing,
		"--ipv6":             predict.Nothing,
		"-s":                 predict.Nothing,
		"--start":            predict.Nothing,
		"-e":                 predict.Nothing,
		"--end":              predict.Nothing,
		"-x":                 predict.Nothing,
		"--exclude-reserved": predict.Nothing,
		"-u":                 predict.Nothing,
		"--unique":           predict.Nothing,
	},
}

func printHelpRandIP() {
	fmt.Printf(
		`Usage: %s randip [<opts>]

Description:
  Generate random IPs.

  By default, generates 1 random IPv4 address with starting range 0.0.0.0 and 
  ending range 255.255.255.255, but can be configured to generate any number of 
  a combination of IPv4/IPv6 addresses within any range.

  Using --ipv6 or -6 without any starting or ending range will generate a IP 
  between range of :: to ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff.

  Note that only IPv4 or IPv6 IPs can be generated, but not both.

  When generating unique IPs, the range size must not be less than the desired 
  number of random IPs.
 
  $ %[1]s randip
  $ %[1]s randip --ipv6 --num 5
  $ %[1]s randip -4 -n 10
  $ %[1]s randip -4 -s 1.1.1.1 -e 10.10.10.10
  $ %[1]s randip -4 -s 1.1.1.1 -e 10.10.10.10 -n 5 -u
  $ %[1]s randip -6 --start 9c61:f71e:656d:d12e:98a3:9814:38cf:5592
  $ %[1]s randip -6 --end eedd:8977:56d9:aac3:947b:29cc:78ea:deab

Options:
  --help, -h
    show help.
  --num, -n 
    number of IPs to generate.
  --ipv4, -4
    generate IPv4 IPs.
  --ipv6, -6
    generate IPv6 IPs.
  --start, -s 
    starting range of IPs.
    default: minimum IP possible for IP type selected.
  --end, -e
    ending range of IPs.
    default: maximum IP possible for IP type selected.
  --exclude-reserved, -x
    exclude reserved/bogon IPs.
    full list can be found at https://ipinfo.io/bogon.
  --unique, -u 
    generate unique IPs.
`, progBase)
}

func cmdRandIP() error {
	f := lib.CmdRandIPFlags{}
	f.Init()
	pflag.Parse()

	return lib.CmdRandIP(f, pflag.Args()[1:], printHelpRandIP)
}
