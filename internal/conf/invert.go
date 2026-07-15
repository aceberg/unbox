package conf

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/aceberg/unbox/internal/share"
	"github.com/aceberg/unbox/pkg/anytls"
	"github.com/aceberg/unbox/pkg/hysteria2"
	"github.com/aceberg/unbox/pkg/shadowsocks"
	"github.com/aceberg/unbox/pkg/trojan"
	"github.com/aceberg/unbox/pkg/tuic"
	"github.com/aceberg/unbox/pkg/vless"
)

// Invert converts sing-box config outbounds to URLs
func Invert() {

	config := getConfigFromFile(share.Settings.InputPath)

	outbounds, ok := config["outbounds"].([]any)
	if !ok {
		log.Println("No outbounds found in", share.Settings.InputPath)
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

		if m["type"] == "anytls" {
			var an anytls.AnyTLS

			if err := json.Unmarshal(b, &an); err != nil {
				continue
			}
			fmt.Println(anytls.ToURL(an))
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

		if m["type"] == "shadowsocks" {
			var ss shadowsocks.Shadowsocks

			if err := json.Unmarshal(b, &ss); err != nil {
				continue
			}
			fmt.Println(shadowsocks.ToURL(ss))
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

		if m["type"] == "tuic" {
			var tu tuic.TUIC

			if err := json.Unmarshal(b, &tu); err != nil {
				continue
			}
			fmt.Println(tuic.ToURL(tu))
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
