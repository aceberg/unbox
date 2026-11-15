package conf

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/aceberg/unbox/internal/api"
	"github.com/aceberg/unbox/internal/check"
	"github.com/aceberg/unbox/internal/share"
)

// Ping checks TCP reachability of outbounds
func Ping() {
	start := time.Now()

	config := getConfigFromFile(share.Settings.OutPath)
	outbounds, ok := config["outbounds"].([]any)
	if !ok {
		log.Println("No outbounds found in", share.Settings.OutPath)
		return
	}

	l := len(outbounds)

	type job struct {
		index int
		data  map[string]any
	}

	jobs := make(chan job)
	online := make([]bool, l)

	var wg sync.WaitGroup

	// Start workers
	for range share.Settings.Workers {
		wg.Go(func() {
			for j := range jobs {
				online[j.index] = pingOne(j.data, j.index+1, l)
			}
		})
	}

	// Submit jobs
	for i, ob := range outbounds {
		m, ok := ob.(map[string]any)
		if !ok {
			continue
		}

		mType := fmt.Sprint(m["type"])
		if mType == "urltest" || mType == "selector" {
			continue
		}

		if mType != "vless" && mType != "vmess" && mType != "trojan" {
			online[i] = true
			continue
		}

		jobs <- job{index: i, data: m}
	}

	close(jobs)
	wg.Wait()

	// Collect results in original order
	var result []api.ProxyServer

	for i, ob := range outbounds {
		if !online[i] {
			continue
		}

		m, _ := ob.(map[string]any)

		result = append(result, api.ProxyServer{
			Tag: fmt.Sprint(m["tag"]),
		})
	}

	fmt.Println()
	log.Println("INFO Scanned servers:", l)
	log.Println("INFO Found online:", len(result))
	log.Println("INFO Ping finished in", time.Since(start).Round(time.Millisecond))

	editConfig(result)
}

func pingOne(m map[string]any, i, l int) bool {
	strOK := share.Col.Ok + "[OK]" + share.Col.Reset
	strX := share.Col.Err + "[NO]" + share.Col.Reset

	mType := fmt.Sprint(m["type"])
	mTag := share.Col.Main + "\"" + fmt.Sprint(m["tag"]) + "\"" + share.Col.Reset

	server := fmt.Sprint(m["server"])
	port := fmt.Sprint(m["server_port"])

	str := mType + " (" + server + ":" + port + ") " + mTag
	pre := "INFO [" + fmt.Sprint(i) + "-" + fmt.Sprint(l) + "]"

	if check.PortReachable(server, port) {
		log.Println(pre, strOK, str)
		return true
	}

	log.Println(pre, strX, str)
	return false
}
