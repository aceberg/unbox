package api

import (
	"io"
	"net/http"

	"github.com/aceberg/unbox/internal/check"
	"github.com/aceberg/unbox/internal/share"
)

// ProxyServer - proxy name and delay
type ProxyServer struct {
	Tag   string
	Delay int
}

// Request sends request to sing-box Clash API
func Request(method, path string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, share.Settings.APIPath+path, body)
	if check.IfError(err) {
		return nil, err
	}

	if share.Settings.APISecret != "" {
		req.Header.Set("Authorization", "Bearer "+share.Settings.APISecret)
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	return http.DefaultClient.Do(req)
}
