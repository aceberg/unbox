package api

import (
	"encoding/json"
	"log"
	"os"
	"slices"

	"github.com/aceberg/unbox/internal/check"
)

func getConfigFromFile() map[string]any {
	f, err := os.ReadFile(Config.OutPath)
	check.IfError(err)

	var config map[string]any
	err = json.Unmarshal(f, &config)
	check.IfError(err)

	return config
}

func editConfig(saveTags []ProxyServer) {

	config := getConfigFromFile()

	outbounds, ok := config["outbounds"].([]any)
	if !ok {
		log.Println("No outbounds found in", Config.OutPath)
		return
	}

	var newOutbounds []any

	for _, o := range outbounds {
		ob := o.(map[string]any)

		if ob["type"] == "urltest" || ob["type"] == "selector" {

			newUrltest := make([]any, len(saveTags))
			for j, tag := range saveTags {
				newUrltest[j] = tag.Tag
			}
			ob["outbounds"] = newUrltest
			newOutbounds = append(newOutbounds, ob)

		} else {
			tag := ob["tag"].(string)
			isAlive := slices.ContainsFunc(saveTags, func(t ProxyServer) bool {
				return t.Tag == tag
			})
			if isAlive {
				newOutbounds = append(newOutbounds, ob)
			}
		}
	}

	config["outbounds"] = newOutbounds

	updated, err := json.MarshalIndent(config, "", "  ")
	check.IfError(err)

	err = os.WriteFile(Config.OutPath, updated, 0644)
	if check.IfError(err) {
		log.Println("ERROR: Output file error")
	}
}
