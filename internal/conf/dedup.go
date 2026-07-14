package conf

import (
	"encoding/json"
	"log"

	"github.com/aceberg/unbox/internal/api"
	"github.com/aceberg/unbox/internal/check"
	"github.com/aceberg/unbox/internal/share"
)

// Deduplicate removes duplicate outbounds from sing-box config
func Deduplicate() {

	config := getConfigFromFile(share.Settings.OutPath)

	outbounds, ok := config["outbounds"].([]any)
	if !ok {
		log.Println("No outbounds found in", share.Settings.OutPath)
		return
	}

	seen := make(map[string]struct{})
	var unique []api.ProxyServer

	for _, ob := range outbounds {
		m, ok := ob.(map[string]any)
		if !ok {
			continue
		}

		if m["type"] == "urltest" || m["type"] == "selector" {
			continue
		}

		// Copy without tag
		cmp := make(map[string]any, len(m))
		for k, v := range m {
			if k != "tag" {
				cmp[k] = v
			}
		}

		keyBytes, err := json.Marshal(cmp)
		check.IfError(err)

		key := string(keyBytes)

		if _, exists := seen[key]; exists {
			log.Println("Dup", m["tag"].(string))
			continue // duplicate except for tag
		}

		seen[key] = struct{}{}
		unique = append(unique, api.ProxyServer{Tag: m["tag"].(string)})
	}

	editConfig(unique)
}
