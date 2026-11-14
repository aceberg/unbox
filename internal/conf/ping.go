package conf

import (
	"fmt"
	"log"

	"github.com/aceberg/unbox/internal/api"
	"github.com/aceberg/unbox/internal/check"
	"github.com/aceberg/unbox/internal/share"
)

// Ping checks TCP reachability of outbounds
func Ping() {
	config := getConfigFromFile(share.Settings.OutPath)
	outbounds, ok := config["outbounds"].([]any)
	if !ok {
		log.Println("No outbounds found in", share.Settings.OutPath)
		return
	}

	var online []api.ProxyServer

	l := len(outbounds)

	for i, ob := range outbounds {
		m, ok := ob.(map[string]any)
		if !ok {
			continue
		}

		mType := fmt.Sprint(m["type"])
		if mType != "vless" && mType != "vmess" && mType != "trojan" {
			online = append(online, api.ProxyServer{Tag: m["tag"].(string)})
			continue
		}

		if pingOne(m, i, l) {
			online = append(online, api.ProxyServer{Tag: m["tag"].(string)})
		}
	}

	editConfig(online)
}

func pingOne(m map[string]any, i, l int) bool {
	strOK := share.Col.Ok + "[OK]" + share.Col.Reset
	strX := share.Col.Err + "[X] " + share.Col.Reset

	mType := fmt.Sprint(m["type"])
	mTag := share.Col.Main + "\"" + fmt.Sprint(m["tag"]) + "\"" + share.Col.Reset

	server := fmt.Sprint(m["server"])
	port := fmt.Sprint(m["server_port"])

	str := mType + " (" + server + ":" + port + ") " + mTag
	pre := "INFO [" + fmt.Sprint(i) + "-" + fmt.Sprint(l) + "]"

	if check.PortReachable(server, port) {
		log.Println(pre, strOK, str)
		return true
	} else {
		log.Println(pre, strX, str)
		return false
	}
}
