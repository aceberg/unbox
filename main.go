package main

import (
	"flag"
	_ "time/tzdata"

	"github.com/aceberg/unbox/internal/api"
	"github.com/aceberg/unbox/internal/file"
)

func main() {
	jsonPtr := flag.Bool("j", false, "Validate and Indent json output")
	namePtr := flag.Bool("n", false, "Rename tags")
	keepPtr := flag.Bool("k", false, "Keep alive")

	apiPtr := flag.String("a", "", "Path to sing-box Clash API")
	filePtr := flag.String("f", "VLESS.txt", "Path to file with links")
	tmplPtr := flag.String("t", "", "Path to template sing-box config")
	outPtr := flag.String("o", "", "Path to output file")
	urlPtr := flag.String("u", "", "URL to test proxies")

	flag.Parse()

	if *apiPtr != "" {

		api.Config = api.Conf{
			ApiPath:   *apiPtr,
			OutPath:   *outPtr,
			KeepAlive: *keepPtr,
			TestURL:   *urlPtr,
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
