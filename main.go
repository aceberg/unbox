package main

import (
	"flag"
	"fmt"
	"os"
	_ "time/tzdata"

	"github.com/aceberg/unbox/internal/check"
	"github.com/aceberg/unbox/internal/conf"
	"github.com/aceberg/unbox/internal/keep"
	"github.com/aceberg/unbox/internal/parse"
	"github.com/aceberg/unbox/internal/share"
)

func printUsage() {
	fmt.Println("Usage: unbox <command>")

	fmt.Println("\nCommands:" +
		"\n  examples  Show usage examples" +
		"\n  conf      Works with sing-box config file" +
		"\n  keep      Keep connection alive, auto switch proxy" +
		"\n  parse     Parse a file with URLs to sing-box config")
}

func main() {
	if len(os.Args) < 2 {
		printUsage()

		return
	}

	switch os.Args[1] {
	case "examples":
		showExamples()
	case "conf":
		confCmd := flag.NewFlagSet("unbox conf", flag.ExitOnError)

		dedupPtr := confCmd.Bool("d", false, "Deduplicate")

		apiPtr := confCmd.String("a", "", "URL of sing-box Clash API")
		secPtr := confCmd.String("as", "", "Clash API secret")
		inpPtr := confCmd.String("i", "", "Path to sing-box config file to get URLs from")
		outPtr := confCmd.String("o", "", "Path to output sing-box config file")
		urlPtr := confCmd.String("u", "", "URL to test proxies")

		limPtr := confCmd.Int("l", 3000, "Timeout for proxy delay (latency) check (ms)")
		benPtr := confCmd.Int("n", 0, "Number of best servers to save (0 - save all)")

		err := confCmd.Parse(os.Args[2:])
		check.IfError(err)

		share.Settings = share.SettingsType{
			APIPath:      *apiPtr,
			APISecret:    *secPtr,
			TestURL:      *urlPtr,
			LimitTimeout: *limPtr,
			OutPath:      *outPtr,
			InputPath:    *inpPtr,
			Deduplicate:  *dedupPtr,
			BestN:        *benPtr,
		}

		if *dedupPtr {
			conf.Deduplicate()
		} else if *inpPtr != "" {
			conf.Invert()
		} else {
			conf.RemoveUnreachable()
		}

	case "keep":
		keepCmd := flag.NewFlagSet("unbox keep", flag.ExitOnError)

		apiPtr := keepCmd.String("a", "", "URL of sing-box Clash API")
		secPtr := keepCmd.String("as", "", "Clash API secret")
		urlPtr := keepCmd.String("u", "", "URL to test proxies")

		allPtr := keepCmd.Int("da", 5*60, "Delay between checks of all proxy servers (s). Use 0 to disable")
		bkpPtr := keepCmd.Int("db", 30, "Delay between checks of backup proxy servers (s). Use 0 to disable")
		mainPtr := keepCmd.Int("dm", 5, "Delay between checks of main proxy server (s). Use 0 to disable")
		switchPtr := keepCmd.Int("ds", 5*60, "Delay between auto switch to a faster proxy attempts (s). Use 0 to disable")
		stepPtr := keepCmd.Int("st", 50, "Switch to a faster proxy if it is at least this many ms faster than the current one (ms).")
		limPtr := keepCmd.Int("l", 3000, "Timeout for proxy delay (latency) check (ms)")
		benPtr := keepCmd.Int("n", 5, "Number of backup servers")

		err := keepCmd.Parse(os.Args[2:])
		check.IfError(err)

		share.Settings = share.SettingsType{
			APIPath:      *apiPtr,
			APISecret:    *secPtr,
			TestURL:      *urlPtr,
			LimitTimeout: *limPtr,
			DelayMain:    *mainPtr,
			DelayBkp:     *bkpPtr,
			DelayAll:     *allPtr,
			DelaySwitch:  *switchPtr,
			SwitchStep:   *stepPtr,
			BackupN:      *benPtr,
		}

		keep.Alive()

	case "parse":
		parseCmd := flag.NewFlagSet("unbox parse", flag.ExitOnError)

		jsonPtr := parseCmd.Bool("j", false, "Validate and Indent JSON output")
		namePtr := parseCmd.Bool("n", false, "Rename tags")

		filePtr := parseCmd.String("f", "VLESS.txt", "Path to file with URLs")
		outPtr := parseCmd.String("o", "", "Path to output sing-box config file")
		tmplPtr := parseCmd.String("t", "", "Path to template sing-box config")

		err := parseCmd.Parse(os.Args[2:])
		check.IfError(err)

		parse.Settings = parse.SettingsType{
			FilePath:     *filePtr,
			TemplatePath: *tmplPtr,
			OutPath:      *outPtr,
			RenameTags:   *namePtr,
			ValidateJSON: *jsonPtr,
		}

		parse.Start()

	default:
		printUsage()
	}
}

func showExamples() {

	fmt.Println("conf")
	fmt.Println("  Remove unreachable nodes from sing-box config:" +
		"\n    unbox conf -a 'http://127.0.0.1:9090' -o sing-box.json")
	fmt.Println("  Remove duplicate outbounds from sing-box config:" +
		"\n    unbox conf -d -o sing-box.json")
	fmt.Println("  Convert sing-box config outbounds to URLs:" +
		"\n    unbox conf -i sing-box.json > URLs.txt")

	fmt.Println("\nkeep")
	fmt.Println("  Keep proxy alive with default intervals between checks:" +
		"\n    unbox keep -a 'http://127.0.0.1:9090'")

	fmt.Println("\nparse")
	fmt.Println("  Parse VLESS.txt file with URLs and print result to stdout:" +
		"\n    unbox parse")
	fmt.Println("  Parse myfile.txt, insert result in template tmpl.json, output to sing-box.json and check/format JSON (-j):" +
		"\n    unbox parse -f myfile.txt -t tmpl.json -o sing-box.json -j")
}
