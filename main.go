package main

import (
	"flag"
	_ "time/tzdata"

	"github.com/aceberg/unbox/internal/api"
	"github.com/aceberg/unbox/internal/file"
)

func main() {
	dedupPtr := flag.Bool("d", false, "Deduplicate")
	jsonPtr := flag.Bool("j", false, "Validate and Indent json output")
	namePtr := flag.Bool("n", false, "Rename tags")
	keepPtr := flag.Bool("k", false, "Keep alive")

	apiPtr := flag.String("a", "", "Path to sing-box Clash API")
	filePtr := flag.String("f", "VLESS.txt", "Path to file with links")
	tmplPtr := flag.String("t", "", "Path to template sing-box config")
	outPtr := flag.String("o", "", "Path to output file")
	urlPtr := flag.String("u", "", "URL to test proxies")

	allPtr := flag.Uint("da", 5*60, "Delay between checks of all proxy servers (s). Use 0 to disable")
	bkpPtr := flag.Uint("db", 30, "Delay between checks of backup proxy servers (s). Use 0 to disable")
	mainPtr := flag.Uint("dm", 5, "Delay between checks of main proxy server (s). Use 0 to disable")
	switchPtr := flag.Uint("ds", 5*60, "Delay between auto switch to a faster proxy attempts (s). Use 0 to disable")

	flag.Parse()

	if *apiPtr != "" || *dedupPtr {

		api.Config = api.Conf{
			ApiPath:     *apiPtr,
			OutPath:     *outPtr,
			Deduplicate: *dedupPtr,
			KeepAlive:   *keepPtr,
			TestURL:     *urlPtr,
			DelayMain:   *mainPtr,
			DelayBkp:    *bkpPtr,
			DelayAll:    *allPtr,
			DelaySwitch: *switchPtr,
		}

		api.Start()
		return
	}

	file.Config = file.Conf{
		FilePath:     *filePtr,
		TemplatePath: *tmplPtr,
		OutPath:      *outPtr,
		RenameTags:   *namePtr,
		ValidateJSON: *jsonPtr,
	}

	file.Start()
}
