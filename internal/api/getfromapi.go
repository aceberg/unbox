package api

import (
	"encoding/json"
	"io"
	"sort"

	"github.com/aceberg/unbox/internal/check"
)

func getProxyList() map[string]any {
	var proxies map[string]any

	resp, err := apiRequest("GET", "/proxies", nil)
	if check.IfError(err) {
		return proxies
	}

	body, err := io.ReadAll(resp.Body)
	check.IfError(err)

	err = resp.Body.Close()
	check.IfError(err)

	var data map[string]any
	err = json.Unmarshal(body, &data)
	check.IfError(err)

	proxies = data["proxies"].(map[string]any)

	return proxies
}

func getAliveTags() []ProxyServer {
	var aliveTags []ProxyServer

	proxies := getProxyList()

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

	allTags = getSelectorTags()
	if len(allTags) > 0 {
		return allTags
	}

	proxies := getProxyList()

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

func getSelectorTags() []string {
	var selTags []string

	proxies := getProxyList()

	for _, p := range proxies {
		proxy := p.(map[string]any)

		if proxy["type"] == "Selector" {

			all, ok := proxy["all"].([]any)
			if ok {
				for _, v := range all {
					if s, ok := v.(string); ok {
						selTags = append(selTags, s)
					}
				}
			}
			break
		}
	}

	return selTags
}

func getSelectorName() string {
	var selName string

	proxies := getProxyList()

	for _, p := range proxies {
		proxy := p.(map[string]any)

		if proxy["type"] == "Selector" {

			selName = proxy["name"].(string)
			break
		}
	}

	return selName
}

func getCurrntProxy() string {
	var cur string

	proxies := getProxyList()

	for _, p := range proxies {
		proxy := p.(map[string]any)

		if proxy["type"] == "Selector" {

			cur = proxy["now"].(string)
			break
		}
	}

	return cur
}
