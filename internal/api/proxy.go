package api

import (
	"io"
	"log"
	"strconv"
	"strings"

	"github.com/aceberg/unbox/internal/check"
	"github.com/aceberg/unbox/internal/share"
)

// CheckOneProxy returns true if proxy is alive
func CheckOneProxy(tag string, logPref string) bool {

	url := "https://www.gstatic.com/generate_204"
	if share.Settings.TestURL != "" {
		url = share.Settings.TestURL
	}
	l := strconv.FormatUint(uint64(share.Settings.LimitTimeout), 10)

	resp, err := Request("GET", "/proxies/"+tag+"/delay?timeout="+l+"&url="+url, nil)
	if check.IfError(err) {
		return false
	}

	body, err := io.ReadAll(resp.Body)
	if check.IfError(err) {
		return false
	}

	err = resp.Body.Close()
	check.IfError(err)

	msg := string(body)

	log.Print("INFO "+logPref+" \""+tag+"\":", msg)

	return !strings.Contains(msg, "message")
}
