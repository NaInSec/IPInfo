package lib

import (
	"fmt"
	"net/netip"

	"github.com/ipinfo/cli/lib/iputil"
	"github.com/spf13/pflag"
)

// CmdToolIsV4In6Flags are flags expected by CmdToolIsV4In6
type CmdToolIsV4In6Flags struct {
	Help bool
}

// Init initializes the common flags available to CmdToolIsV4In6 with sensible
func (f *CmdToolIsV4In6Flags) Init() {
	pflag.BoolVarP(
		&f.Help,
		"help", "h", false,
		"show help.",
	)
}

// CmdToolIsV4In6 checks if given ip is an IPv4-mapped IPv6 address.
func CmdToolIsV4In6(f CmdToolIsV4In6Flags, args []string, printHelp func()) error {
	if f.Help {
		printHelp()
		return nil
	}

	op := func(input string, inputType iputil.INPUT_TYPE) error {
		switch inputType {
		case iputil.INPUT_TYPE_IP:
			addr, err := netip.ParseAddr(input)
			if err != nil {
				return iputil.ErrInvalidInput
			}

			fmt.Printf("%s,%t\n", input, addr.Is4In6())
		default:
			return iputil.ErrInvalidInput
		}
		return nil
	}

	return iputil.GetInputFrom(args, true, true, op)
}
