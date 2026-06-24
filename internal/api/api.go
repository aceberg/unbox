package api

import (
	"io"
	"net/http"

	"github.com/aceberg/unbox/internal/check"
)

func apiRequest(method, path string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, Config.ApiPath+path, body)
	if check.IfError(err) {
		return nil, err
	}

	if Config.ApiSecret != "" {
		req.Header.Set("Authorization", "Bearer "+Config.ApiSecret)
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	return http.DefaultClient.Do(req)
}
