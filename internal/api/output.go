package api

import (
	"encoding/json"
	"log"
	"os"

	"github.com/aceberg/unbox/internal/check"
)

func editConfig(aliveTags []ProxyServer) {

	f, err := os.ReadFile(Config.OutPath)
	check.IfError(err)

	var config map[string]any
	err = json.Unmarshal(f, &config)
	check.IfError(err)

	outbounds, ok := config["outbounds"].([]any)
	if !ok {
		log.Println("No outbounds found in", Config.OutPath)
		return
	}

	newOutbounds := make([]any, len(outbounds))

	for i, o := range outbounds {
		ob := o.(map[string]any)

		if ob["type"] == "urltest" {

			newUrltest := make([]any, len(aliveTags))
			for j, tag := range aliveTags {
				newUrltest[j] = tag.Tag
			}
			ob["outbounds"] = newUrltest
			newOutbounds[i] = ob

		} else {
			newOutbounds[i] = o
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
