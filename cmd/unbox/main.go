// Unbox converts a list of `vless://` links to a sing-box config file.
//
// By default, it reads VLESS.txt in the current directory.
//
// Usage:
//
//	unbox -f VLESS.txt
//
// With template and output file:
//
//	unbox -f VLESS.txt -t sing-box.tmpl.json -o sing-box.json

package main

import (
	"flag"

	"github.com/aceberg/unbox/pkg/file"
)

func main() {
	filePtr := flag.String("f", "VLESS.txt", "Path to file with vless:// links")
	tmplPtr := flag.String("t", "", "Path to template sing-box config")
	outPtr := flag.String("o", "", "Path to output file")
	namePtr := flag.String("n", "no", "Rename tags (yes|no)")
	jsonPtr := flag.String("j", "no", "Validate and Indent json output (yes|no)")

	flag.Parse()

	file.Config = file.Conf{
		FilePath:     *filePtr,
		TemplatePath: *tmplPtr,
		OutPath:      *outPtr,
	}

	if *namePtr == "yes" {
		file.Config.RenameTags = true
	}

	if *jsonPtr == "yes" {
		file.Config.ValidateJSON = true
	}

	file.Parse()
}
