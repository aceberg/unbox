package api

import (
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/aceberg/unbox/internal/check"
	"github.com/aceberg/unbox/internal/share"
)

// SwitchProxy switches proxy
func SwitchProxy(selName, tag string) error {

	body := strings.NewReader(`{"name":"` + tag + `"}`)

	resp, err := Request("PUT", "/proxies/"+selName, body)
	if err != nil {
		return err
	}
	defer check.IfError(resp.Body.Close())

	if resp.StatusCode != http.StatusNoContent {
		return errors.New("unexpected status:" + strconv.Itoa(resp.StatusCode))
	}

	return nil
}

// CheckOneProxy returns true if proxy is alive
func CheckOneProxy(tag string, logPref string) bool {
	var online bool

	l := strconv.FormatUint(uint64(share.Settings.LimitTimeout), 10)

	resp, err := Request("GET", "/proxies/"+url.PathEscape(tag)+"/delay?timeout="+l+"&url="+url.QueryEscape(share.Settings.TestURL), nil)
	if check.IfError(err) {
		return false
	}

	body, err := io.ReadAll(resp.Body)
	if check.IfError(err) {
		return false
	}

	check.IfError(resp.Body.Close())

	msg := strings.TrimRight(string(body), "\r\n")

	if !strings.Contains(msg, "message") {
		online = true
		msg = share.Col.Ok + msg + share.Col.Reset
	}

	log.Print("INFO "+logPref+" \""+tag+"\":", msg)

	return online
}
