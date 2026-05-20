package file

import (
	"bufio"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/aceberg/unbox/internal/check"
	"github.com/aceberg/unbox/internal/hysteria2"
	"github.com/aceberg/unbox/internal/trojan"
	"github.com/aceberg/unbox/internal/vless"
)

// Parse file with links
func parseFile() {

	file, ok := getLinksFromFile()
	if !ok {
		return
	}

	i = 1

	scanner := bufio.NewScanner(strings.NewReader(file))
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
				v.Tag = renameTag(v.Tag)
				addResult(v, v.Tag)
			}
		}

		// keep only Hysteria2 links (case-insensitive)
		if strings.HasPrefix(strings.ToLower(line), "hysteria2://") {
			h, err := hysteria2.ParseHyst2(line)
			if !check.IfError(err) {
				h.Tag = renameTag(h.Tag)
				addResult(h, h.Tag)
			}
		}

		// keep only Trojan links (case-insensitive)
		if strings.HasPrefix(strings.ToLower(line), "trojan://") {
			t, err := trojan.ParseTrojan(line)
			if !check.IfError(err) {
				t.Tag = renameTag(t.Tag)
				addResult(t, t.Tag)
			}
		}
	}

	err := scanner.Err()
	check.IfError(err)
}

func renameTag(in string) (out string) {

	if Config.RenameTags {
		out = fmt.Sprint("tag", i)
	} else {
		out = in + fmt.Sprint(" ", i)
	}
	i = i + 1

	return out
}

func addResult(a any, t string) {

	data, _ := json.MarshalIndent(a, "", "  ")
	result = append(result, string(data))
	tags = append(tags, "\""+t+"\"")
}
