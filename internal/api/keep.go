package api

import (
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/aceberg/unbox/internal/check"
)

var (
	currentProxy string
	alive        bool
	aliveTags    []ProxyServer
)

const (
	colorErr   = "\033[31m" // red
	colorBkp   = "\033[32m" // green
	colorWarn  = "\033[33m" // yellow
	colorMain  = "\033[36m" // cyan
	colorReset = "\033[0m"  // reset
)

func keepConnectionAlive() {

	if Config.DelayMain > 0 {
		go testCurrentProxy()
	}
	if Config.DelayBkp > 0 {
		go updateAliveTags()
	}
	if Config.DelayAll > 0 {
		go testAllTagsRoutine()
	}
	if Config.DelaySwitch > 0 {
		go checkSwitch()
	}

	for {
		if !alive {
			tag, ok := chooseTag()
			if ok {
				alive = true
				switchProxy(tag)
			} else {
				log.Println(colorErr + "ERROR" + colorMain + "[MAIN] " + colorReset + "No proxies online!")
			}
		}

		time.Sleep(time.Duration(1) * time.Second)
	}
}

func testOneProxy(tag string, logPref string) bool {

	url := "https://www.gstatic.com/generate_204"
	if Config.TestURL != "" {
		url = Config.TestURL
	}

	resp, err := http.Get(Config.ApiPath + "/proxies/" + tag + "/delay?timeout=3000&url=" + url)
	if check.IfError(err) {
		return false
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if check.IfError(err) {
		return false
	}

	msg := string(body)

	log.Print("INFO "+logPref+" \""+tag+"\":", msg)

	return !strings.Contains(msg, "message")
}

func switchProxy(tag string) {

	selName := getSelectorName()
	if selName == "" {
		log.Println(colorErr + "ERROR" + colorMain + "[MAIN] " + colorReset + "Can't get Selector tag name to select new proxy")
		return
	}

	log.Println(colorWarn+"WARN "+colorMain+"[MAIN] "+colorReset+"Selecting proxy:", tag)
	currentProxy = tag

	body := strings.NewReader(`{"name":"` + tag + `"}`)

	req, _ := http.NewRequest(
		"PUT",
		Config.ApiPath+"/proxies/"+selName,
		body,
	)
	req.Header.Set("Content-Type", "application/json")

	_, err := http.DefaultClient.Do(req)
	check.IfError(err)
}

func chooseTag() (string, bool) {

	aliveTags = getAliveTags()

	for _, tag := range aliveTags {
		if tag.Tag != currentProxy {
			return tag.Tag, true
		}
	}

	return "", false
}
