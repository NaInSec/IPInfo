package main

import (
	"os"

	"github.com/ipinfo/cli/lib/complete"
	"github.com/ipinfo/cli/lib/complete/predict"
	"github.com/ipinfo/cli/lib/iputil"
)

var completions = &complete.Command{
	Sub: map[string]*complete.Command{
		"myip":       completionsMyIP,
		"bulk":       completionsBulk,
		"asn":        completionsASN,
		"summarize":  completionsSummarize,
		"map":        completionsMap,
		"prips":      completionsPrips,
		"grepip":     completionsGrepIP,
		"matchip":    completionsMatchIP,
		"grepdomain": completionsGrepDomain,
		"cidr2range": completionsCIDR2Range,
		"cidr2ip":    completionsCIDR2IP,
		"range2cidr": completionsRange2CIDR,
		"range2ip":   completionsRange2IP,
		"randip":     completionsRandIP,
		"splitcidr":  completionsSplitCIDR,
		"mmdb":       completionsMmdb,
		"calc":       completionsCalc,
		"tool":       completionsTool,
		"download":   completionsDownload,
		"cache":      completionsCache,
		"quota":      completionsQuota,
		"init":       completionsInit,
		"logout":     completionsLogout,
		"config":     completionsConfig,
		"completion": completionsCompletion,
		"version":    completionsVersion,
	},
	Flags: map[string]complete.Predictor{
		"-v":        predict.Nothing,
		"--version": predict.Nothing,
		"-h":        predict.Nothing,
		"--help":    predict.Nothing,
	},
}

func handleCompletions() {
	line := os.Getenv("COMP_LINE")
	args := complete.Parse(line)
	if len(args) > 1 {
		cmdSecondArg := args[1].Text
		if iputil.StrIsIPStr(cmdSecondArg) {
			completions.Sub[cmdSecondArg] = completionsIP
		} else if iputil.StrIsASNStr(cmdSecondArg) {
			completions.Sub[cmdSecondArg] = completionsASNSingle
		}
	}

	completions.Complete(progBase)
}
