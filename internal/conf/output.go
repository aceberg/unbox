package conf

import (
	"encoding/json"
	"log"
	"os"
	"slices"

	"github.com/aceberg/unbox/internal/api"
	"github.com/aceberg/unbox/internal/check"
	"github.com/aceberg/unbox/internal/share"
)

func getConfigFromFile(path string) map[string]any {
	f, err := os.ReadFile(path)
	check.IfError(err)

	var config map[string]any
	err = json.Unmarshal(f, &config)
	check.IfError(err)

	return config
}

func editConfig(saveTags []api.ProxyServer) {

	config := getConfigFromFile(share.Settings.OutPath)

	outbounds, ok := config["outbounds"].([]any)
	if !ok {
		log.Println("No outbounds found in", share.Settings.OutPath)
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
			isAlive := slices.ContainsFunc(saveTags, func(t api.ProxyServer) bool {
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

	err = os.WriteFile(share.Settings.OutPath, updated, 0644)
	if check.IfError(err) {
		log.Println("ERROR: Output file error")
	}
}
