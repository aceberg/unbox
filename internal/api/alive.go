package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/aceberg/unbox/internal/check"
)

func getAliveTags() []string {
	var aliveTags []string

	resp, err := http.Get(Config.ApiPath + "/proxies")
	check.IfError(err)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	check.IfError(err)

	var data map[string]any
	err = json.Unmarshal(body, &data)
	check.IfError(err)

	proxies := data["proxies"].(map[string]any)

	for tag, p := range proxies {
		proxy := p.(map[string]any)

		ptype := proxy["type"]
		if ptype == "URLTest" {
			continue
		}

		history, ok := proxy["history"].([]any)
		if !ok || len(history) == 0 {
			continue
		}

		aliveTags = append(aliveTags, tag)
	}

	return aliveTags
}
