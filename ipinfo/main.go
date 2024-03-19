package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/ipinfo/cli/lib/iputil"
)

var progBase = filepath.Base(os.Args[0])
var version = "3.3.1"

// global flags.
var fHelp bool
var fHelpDetailed bool
var fNoCache bool
var fNoColor bool

func main() {
	var err error
	var cmd string

	// obey NO_COLOR env var.
	if os.Getenv("NO_COLOR") != "" {
		color.NoColor = true
	}

	handleCompletions()

	if len(os.Args) > 1 {
		cmd = os.Args[1]
	}

	switch {
	case iputil.StrIsIPStr(cmd):
		err = cmdIP(cmd)
	case iputil.StrIsASNStr(cmd):
		asn := strings.ToUpper(cmd)
		err = cmdASNSingle(asn)
	case len(cmd) >= 3 && strings.IndexByte(cmd, '.') != -1:
		err = cmdDomain(cmd)
	case cmd == "myip":
		err = cmdMyIP()
	case cmd == "bulk":
		err = cmdBulk()
	case cmd == "asn":
		err = cmdASN()
	case cmd == "summarize" || cmd == "sum":
		err = cmdSum()
	case cmd == "map":
		err = cmdMap()
	case cmd == "prips":
		err = cmdPrips()
	case cmd == "grepip":
		err = cmdGrepIP()
	case cmd == "matchip":
		err = cmdMatchIP()
	case cmd == "grepdomain":
		err = cmdGrepDomain()
	case cmd == "cidr2range":
		err = cmdCIDR2Range()
	case cmd == "cidr2ip":
		err = cmdCIDR2IP()
	case cmd == "range2cidr":
		err = cmdRange2CIDR()
	case cmd == "range2ip":
		err = cmdRange2IP()
	case cmd == "randip":
		err = cmdRandIP()
	case cmd == "splitcidr":
		err = cmdSplitCIDR()
	case cmd == "mmdb":
		err = cmdMmdb()
	case cmd == "calc":
		err = cmdCalc()
	case cmd == "download":
		err = cmdDownload()
	case cmd == "tool":
		err = cmdTool()
	case cmd == "cache":
		err = cmdCache()
	case cmd == "config":
		err = cmdConfig()
	case cmd == "quota":
		err = cmdQuota()
	case cmd == "init":
		err = cmdInit()
	case cmd == "logout":
		err = cmdLogout()
	case cmd == "completion":
		err = cmdCompletion()
	case cmd == "version" || cmd == "vsn" || cmd == "v":
		err = cmdVersion()
	default:
		err = cmdDefault()
	}

	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}
