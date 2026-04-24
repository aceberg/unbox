package main

import (
	"flag"

	"github.com/aceberg/unbox/pkg/file"
)

func main() {
	filePtr := flag.String("f", "VLESS.txt", "Path to file with vless:// links")
	tmplPtr := flag.String("t", "", "Path to template sing-box config")
	namePtr := flag.String("n", "no", "Rename tags (yes|no)")

	flag.Parse()

	file.Config = file.Conf{
		FilePath:     *filePtr,
		TemplatePath: *tmplPtr,
	}

	if *namePtr == "yes" {
		file.Config.RenameTags = true
	}

	file.Parse()
}
