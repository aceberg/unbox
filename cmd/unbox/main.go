package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/aceberg/unbox/pkg/check"
	"github.com/aceberg/unbox/pkg/vless"
)

const filePath = "VLESS.txt"

var nameTags bool

func main() {
	filePtr := flag.String("f", filePath, "Path to file with vless:// links")
	namePtr := flag.String("n", "yes", "Rename tags (yes|no)")
	// t - template of sing-box.json
	// o - output file
	flag.Parse()

	if *namePtr == "yes" {
		nameTags = true
	}

	data, err := os.ReadFile(*filePtr)
	check.IfError(err)

	extractStrings(string(data))
}

func extractStrings(input string) {
	var res string
	var tags string

	r := strings.NewReader(input)
	i := 0

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// skip empty lines and comments
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// keep only vless links (case-insensitive)
		if strings.HasPrefix(strings.ToLower(line), "vless://") {
			conf, err := vless.ParseVLESS(line)
			if !check.IfError(err) {
				if nameTags {
					i = i + 1
					conf.Tag = fmt.Sprint("tag", i)
					tags = tags + "\"" + conf.Tag + "\","
				}

				data, _ := json.MarshalIndent(conf, "", "  ")
				res = res + "\n" + string(data) + ","
			}
		}
	}

	err := scanner.Err()
	check.IfError(err)

	tags = tags[:len(tags)-1]
	fmt.Printf(`{
	"type": "urltest",
	"tag": "auto",
	"outbounds": [%s],
	"url": "https://www.gstatic.com/generate_204",
	"interval": "10m"
},`, tags)

	res = res[:len(res)-1]
	fmt.Println(res)
}
