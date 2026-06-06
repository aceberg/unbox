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
	colorRed     = "\033[31m"
	colorGreen   = "\033[32m"
	colorYellow  = "\033[33m"
	colorBlue    = "\033[34m"
	colorMagenta = "\033[35m"
	colorCyan    = "\033[36m"
	colorReset   = "\033[0m"
)

func keepConnectionAlive() {

	go testCurrentProxy()
	go updateAliveTags()
	go testAllTags()

	for {
		if !alive {
			tag, ok := chooseTag()
			if ok {
				alive = true
				switchProxy(tag)
			} else {
				log.Println(colorRed + "ERROR" + colorCyan + "[MAIN] " + colorReset + "No proxies online!")
			}
		}

		time.Sleep(time.Duration(1) * time.Second)
	}
}

func testOneProxy(tag string, logPref string) bool {

	resp, err := http.Get(Config.ApiPath + "/proxies/" + tag + "/delay?timeout=3000")
	check.IfError(err)

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	check.IfError(err)

	msg := string(body)

	log.Print("INFO "+logPref+" \""+tag+"\":", msg)

	return !strings.Contains(msg, "message")
}

func switchProxy(tag string) {

	selName := getSelectorName()
	if selName == "" {
		log.Println(colorRed + "ERROR" + colorCyan + "[MAIN] " + colorReset + "Can't get Selector tag name to select new proxy")
		return
	}

	log.Println(colorYellow+"WARN "+colorCyan+"[MAIN] "+colorReset+"Selecting proxy:", tag)
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

func testCurrentProxy() {

	for {
		if currentProxy != "" {
			ok := testOneProxy(currentProxy, colorCyan+"[MAIN]"+colorReset)
			if !ok {
				alive = false
			}
		}

		time.Sleep(time.Duration(5) * time.Second)
	}
}

func updateAliveTags() {

	for {
		for i, tag := range aliveTags {

			if tag.Tag == currentProxy {
				continue
			}

			if i > 3 {
				break
			}

			testOneProxy(tag.Tag, colorGreen+"[BKP] "+colorReset)
		}

		aliveTags = getAliveTags()

		time.Sleep(time.Duration(30) * time.Second)
	}
}

func testAllTags() {
	for {
		for _, tag := range getAllTags() {

			if tag == currentProxy {
				continue
			}

			testOneProxy(tag, "[ALL] ")
		}

		aliveTags = getAliveTags()

		time.Sleep(time.Duration(5*60) * time.Second)
	}
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
