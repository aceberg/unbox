package api

import (
	"encoding/json"
	"io"
	"net/http"
	"sort"

	"github.com/aceberg/unbox/internal/check"
)

type ProxyServer struct {
	Tag   string
	Delay int
}

func getAliveTags() []ProxyServer {
	var aliveTags []ProxyServer

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
		if ptype == "URLTest" || ptype == "Fallback" || ptype == "Selector" {
			continue
		}

		history, ok := proxy["history"].([]any)
		if !ok || len(history) == 0 {
			continue
		}

		last := history[len(history)-1].(map[string]any)
		delay := last["delay"].(float64)

		aliveTags = append(aliveTags, ProxyServer{Tag: tag, Delay: int(delay)})
	}

	sort.Slice(aliveTags, func(i, j int) bool {
		return aliveTags[i].Delay < aliveTags[j].Delay
	})

	return aliveTags
}

func getAllTags() []string {
	var allTags []string

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
		if ptype == "URLTest" || ptype == "Fallback" || ptype == "Selector" {
			continue
		}

		allTags = append(allTags, tag)
	}

	return allTags
}
