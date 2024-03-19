package main

import (
	"fmt"
	"os"

	"github.com/ipinfo/cli/lib/complete"
	"github.com/ipinfo/cli/lib/complete/predict"
	"github.com/spf13/pflag"
)

var completionsTool = &complete.Command{
	Sub: map[string]*complete.Command{
		"aggregate":                    completionsToolAggregate,
		"next":                         completionsToolNext,
		"prev":                         completionsToolPrev,
		"is_v4":                        completionsToolIsV4,
		"is_v6":                        completionsToolIsV6,
		"is_valid":                     completionsToolIsValid,
		"is_one_ip":                    completionsToolIsOneIp,
		"unmap":                        completionsToolUnmap,
		"lower":                        completionsToolLower,
		"upper":                        completionsToolUpper,
		"is_v4in6":                     completionsToolIs4In6,
		"ip2n":                         completionsToolIP2n,
		"n2ip":                         completionsToolN2IP,
		"n2ip6":                        completionsToolN2IP6,
		"prefix":                       completionsToolPrefix,
		"is_loopback":                  completionsToolIsLoopBack,
		"is_multicast":                 completionsToolIsMulticast,
		"is_unspecified":               completionsToolIsUnspecified,
		"is_global_unicast":            completionsToolIsGlobalUnicast,
		"is_link_local_unicast":        completionsToolIsLinkLocalUnicast,
		"is_link_local_multicast":      completionsToolIsLinkLocalMulticast,
		"is_interface_local_multicast": completionsToolIsInterfaceLocalMulticast,
	},
	Flags: map[string]complete.Predictor{
		"-h":     predict.Nothing,
		"--help": predict.Nothing,
	},
}

func printHelpTool() {

	fmt.Printf(
		`Usage: %s tool <cmd> [<opts>] [<args>]

Commands:
  aggregate                    aggregate IPs, IP ranges, and CIDRs.
  next                         get the next IP of the input IP
  prev                         get the previous IP of the input IP
  is_v4                        reports whether input is an IPv4 address.
  is_v6                        reports whether input is an IPv6 address.
  is_valid                     reports whether an IP is valid.
  is_one_ip                    checks whether a CIDR or IP Range contains exactly one IP.
  unmap                        returns ip with any IPv4-mapped IPv6 address prefix removed.	
  lower                        get start IP of IPs, IP ranges, and CIDRs.
  upper                        get end IP of IPs, IP ranges, and CIDRs.
  is_v4in6                     get whether the IP is an IPv4-mapped IPv6 address.
  ip2n                         converts an IPv4 or IPv6 address to its decimal representation.
  n2ip	                       evaluates a mathematical expression and converts it to an IPv4 or IPv6.
  n2ip6	                       evaluates a mathematical expression and converts it to an IPv6.
  prefix                       misc. prefix tools related to CIDRs.
  is_loopback                  reports whether an IP is a valid loopback address.
  is_multicast                 reports whether an IP is a valid multicast address.
  is_unspecified               reports whether an IP is an unspecified address.
  is_global_unicast            reports whether an IP is a global unicast address.
  is_link_local_unicast        reports whether IP is a link local unicast.
  is_link_local_multicast      reports whether IP is a link local multicast address.
  is_interface_local_multicast reports whether IP is an interface local multicast.

Options:
  --help, -h
    show help.
`, progBase)
}

func toolHelp() (err error) {
	pflag.BoolVarP(&fHelp, "help", "h", false, "show help.")
	pflag.Parse()

	if fHelp {
		printHelpTool()
		return nil
	}

	printHelpTool()
	return nil
}

func cmdTool() error {
	var err error
	cmd := ""
	if len(os.Args) > 2 {
		cmd = os.Args[2]
	}

	switch {
	case cmd == "aggregate":
		err = cmdToolAggregate()
	case cmd == "next":
		err = cmdToolNext()
	case cmd == "prev":
		err = cmdToolPrev()
	case cmd == "is_v4":
		err = cmdToolIsV4()
	case cmd == "is_v6":
		err = cmdToolIsV6()
	case cmd == "is_valid":
		err = cmdToolIsValid()
	case cmd == "is_one_ip":
		err = cmdToolIsOneIp()
	case cmd == "unmap":
		err = cmdToolUnmap()
	case cmd == "lower":
		err = cmdToolLower()
	case cmd == "upper":
		err = cmdToolUpper()
	case cmd == "is_v4in6":
		err = cmdToolIsV4In6()
	case cmd == "ip2n":
		err = cmdToolIP2n()
	case cmd == "n2ip":
		err = cmdToolN2IP()
	case cmd == "n2ip6":
		err = cmdToolN2IP6()
	case cmd == "prefix":
		err = cmdToolPrefix()
	case cmd == "is_loopback":
		err = cmdToolIsLoopBack()
	case cmd == "is_multicast":
		err = cmdToolIsMultiCast()
	case cmd == "is_unspecified":
		err = cmdToolIsUnspecified()
	case cmd == "is_global_unicast":
		err = cmdToolisGlobalUnicast()
	case cmd == "is_link_local_unicast":
		err = cmdToolIsLinkLocalUnicast()
	case cmd == "is_link_local_multicast":
		err = cmdToolIsLinkLocalMulticast()
	case cmd == "is_interface_local_multicast":
		err = cmdToolIsInterfaceLocalMulticast()
	default:
		err = toolHelp()
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "err: %v\n", err)
	}

	return nil
}
