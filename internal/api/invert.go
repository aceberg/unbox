package api

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/aceberg/unbox/pkg/hysteria2"
	"github.com/aceberg/unbox/pkg/trojan"
	"github.com/aceberg/unbox/pkg/vless"
)

func invertConf() {

	config := getConfigFromFile(Config.InputPath)

	outbounds, ok := config["outbounds"].([]any)
	if !ok {
		log.Println("No outbounds found in", Config.InputPath)
		return
	}

	for _, ob := range outbounds {
		m, ok := ob.(map[string]any)
		if !ok {
			continue
		}

		if m["type"] == "urltest" || m["type"] == "selector" {
			continue
		}

		b, err := json.Marshal(m)
		if err != nil {
			continue
		}

		if m["type"] == "hysteria2" {
			var hy hysteria2.Hysteria2

			if err := json.Unmarshal(b, &hy); err != nil {
				continue
			}
			fmt.Println(hysteria2.ToURL(hy))
			continue
		}

		if m["type"] == "trojan" {
			var tr trojan.Trojan

			if err := json.Unmarshal(b, &tr); err != nil {
				continue
			}
			fmt.Println(trojan.ToURL(tr))
			continue
		}

		if m["type"] == "vless" {
			var vl vless.VLESS

			if err := json.Unmarshal(b, &vl); err != nil {
				continue
			}

			fmt.Println(vless.ToURL(vl))
			continue
		}
	}
}
