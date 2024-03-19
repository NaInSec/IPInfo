package lib

import (
	"fmt"
	"net"

	"github.com/ipinfo/cli/lib/iputil"
	"github.com/spf13/pflag"
)

type CmdToolIsV4Flags struct {
	Help  bool
	Quiet bool
}

func (f *CmdToolIsV4Flags) Init() {
	pflag.BoolVarP(
		&f.Help,
		"help", "h", false,
		"show help.",
	)
	pflag.BoolVarP(
		&f.Quiet,
		"quiet", "q", false,
		"quiet mode; suppress additional output.",
	)
}

func CmdToolIsV4(
	f CmdToolIsV4Flags,
	args []string,
	printHelp func(),
) error {
	if f.Help {
		printHelp()
		return nil
	}

	actionFunc := func(input string, inputType iputil.INPUT_TYPE) error {
		switch inputType {
		case iputil.INPUT_TYPE_IP:
			ActionForIsV4(input)
		case iputil.INPUT_TYPE_IP_RANGE:
			ActionForIsV4Range(input)
		case iputil.INPUT_TYPE_CIDR:
			ActionForIsV4CIDR(input)
		}
		return nil
	}
	err := iputil.GetInputFrom(args, true, true, actionFunc)
	if err != nil {
		fmt.Println(err)
	}

	return nil
}

func ActionForIsV4(input string) {
	ip := net.ParseIP(input)
	isIPv4 := IsIPv4(ip)

	fmt.Printf("%s,%v\n", input, isIPv4)
}

func ActionForIsV4Range(input string) {
	ipRange, err := iputil.IPRangeStrFromStr(input)
	if err != nil {
		return
	}

	startIP := net.ParseIP(ipRange.Start)
	isIPv4 := IsIPv4(startIP)

	fmt.Printf("%s,%v\n", input, isIPv4)
}

func ActionForIsV4CIDR(input string) {
	_, ipnet, err := net.ParseCIDR(input)
	if err == nil {
		isCIDRIPv4 := IsCIDRIPv4(ipnet)
		fmt.Printf("%s,%v\n", input, isCIDRIPv4)
	} else {
		return
	}
}

func IsIPv4(ip net.IP) bool {
	return ip != nil && ip.To4() != nil
}

func IsCIDRIPv4(ipnet *net.IPNet) bool {
	return ipnet != nil && ipnet.IP.To4() != nil
}
