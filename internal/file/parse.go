package file

import (
	"bufio"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/aceberg/unbox/internal/check"
	"github.com/aceberg/unbox/internal/hysteria2"
	"github.com/aceberg/unbox/internal/vless"
)

// Conf contains command-line options for unbox
type Conf struct {
	FilePath     string
	OutPath      string
	TemplatePath string
	RenameTags   bool
	ValidateJSON bool
}

// Config - app config
var Config Conf

// Parse file with VLESS links
func Parse() {
	var res, tags string

	file, ok := getLinksFromFile()
	if !ok {
		return
	}

	r := strings.NewReader(file)
	i := 1

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// skip empty lines and comments
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// keep only vless links (case-insensitive)
		if strings.HasPrefix(strings.ToLower(line), "vless://") {
			v, err := vless.ParseVLESS(line)
			if !check.IfError(err) {
				if Config.RenameTags {
					v.Tag = fmt.Sprint("tag", i)
				} else {
					v.Tag = v.Tag + fmt.Sprint(" ", i)
				}
				i = i + 1
				data, _ := json.MarshalIndent(v, "", "  ")

				tags = tags + "\"" + v.Tag + "\","
				res = res + string(data) + ",\n"
			}
		}

		// keep only Hysteria2 links (case-insensitive)
		if strings.HasPrefix(strings.ToLower(line), "hysteria2://") {
			h, err := hysteria2.ParseHyst2(line)
			if !check.IfError(err) {
				if Config.RenameTags {
					h.Tag = fmt.Sprint("tag", i)
				} else {
					h.Tag = h.Tag + fmt.Sprint(" ", i)
				}
				i = i + 1
				data, _ := json.MarshalIndent(h, "", "  ")

				tags = tags + "\"" + h.Tag + "\","
				res = res + string(data) + ",\n"
			}
		}
	}

	err := scanner.Err()
	check.IfError(err)

	if i > 1 {
		tags = tags[:len(tags)-1]
		res = res[:len(res)-2]
	}

	var out string

	if Config.TemplatePath != "" {
		out = insertToTemplate(res, tags)
	} else {
		out = res
	}

	if Config.ValidateJSON {
		out = valIndent(out)
	}

	if Config.OutPath != "" {
		outToFile(out)
	} else {
		fmt.Println(out)
	}
}
