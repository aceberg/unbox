package file

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/aceberg/unbox/pkg/check"
	"github.com/aceberg/unbox/pkg/vless"
)

// Conf - app config type
type Conf struct {
	FilePath     string
	TemplatePath string
	RenameTags   bool
}

// Config - app config
var Config Conf

// Parse file with links
func Parse() {
	var res, tags string

	f, err := os.ReadFile(Config.FilePath)
	if check.IfError(err) {
		log.Println("File error. Exiting...")
	}

	r := strings.NewReader(string(f))
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
	}

	err = scanner.Err()
	check.IfError(err)

	if i > 1 {
		tags = tags[:len(tags)-1]
		res = res[:len(res)-2]
	}

	if Config.TemplatePath != "" {
		insertToTemplate(res, tags)
	} else {
		fmt.Println(res)
		fmt.Println("\nTAGS:", tags)
	}
}
