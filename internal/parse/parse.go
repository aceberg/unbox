package parse

import (
	"bufio"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/aceberg/unbox/internal/check"
	"github.com/aceberg/unbox/pkg/anytls"
	"github.com/aceberg/unbox/pkg/hysteria2"
	"github.com/aceberg/unbox/pkg/shadowsocks"
	"github.com/aceberg/unbox/pkg/trojan"
	"github.com/aceberg/unbox/pkg/tuic"
	"github.com/aceberg/unbox/pkg/vless"
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

		lowLine := strings.ToLower(line)

		// keep only vless links (case-insensitive)
		if strings.HasPrefix(lowLine, "vless://") {
			v, err := vless.Parse(line)
			if !check.IfError(err) {
				v.Tag = renameTag(v.Tag)
				addResult(v, v.Tag)
			}
		}

		// keep only Hysteria2 links (case-insensitive)
		if strings.HasPrefix(lowLine, "hysteria2://") || strings.HasPrefix(lowLine, "hy2://") {
			h, err := hysteria2.Parse(line)
			if !check.IfError(err) {
				h.Tag = renameTag(h.Tag)
				addResult(h, h.Tag)
			}
		}

		// keep only Trojan links (case-insensitive)
		if strings.HasPrefix(lowLine, "trojan://") {
			t, err := trojan.Parse(line)
			if !check.IfError(err) {
				t.Tag = renameTag(t.Tag)
				addResult(t, t.Tag)
			}
		}

		// keep only Shadowsocks links (case-insensitive)
		if strings.HasPrefix(lowLine, "ss://") {
			s, err := shadowsocks.Parse(line)
			if !check.IfError(err) {
				s.Tag = renameTag(s.Tag)
				addResult(s, s.Tag)
			}
		}

		// keep only AnyTLS links (case-insensitive)
		if strings.HasPrefix(lowLine, "anytls://") {
			a, err := anytls.Parse(line)
			if !check.IfError(err) {
				a.Tag = renameTag(a.Tag)
				addResult(a, a.Tag)
			}
		}

		// keep only TUIC links (case-insensitive)
		if strings.HasPrefix(lowLine, "tuic://") {
			c, err := tuic.Parse(line)
			if !check.IfError(err) {
				c.Tag = renameTag(c.Tag)
				addResult(c, c.Tag)
			}
		}
	}

	err := scanner.Err()
	check.IfError(err)
}

func renameTag(in string) (out string) {

	if Settings.RenameTags {
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
