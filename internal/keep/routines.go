package keep

import (
	"log"
	"strconv"
	"time"

	"github.com/aceberg/unbox/internal/api"
	"github.com/aceberg/unbox/internal/share"
)

func testCurrentProxy() {

	for {
		currentProxy = api.GetCurrntProxy()
		if currentProxy != "" {
			ok := api.CheckOneProxy(currentProxy, colorMain+"[MAIN]"+colorReset)
			if !ok {
				alive = false
			}
		}

		time.Sleep(time.Duration(share.Settings.DelayMain) * time.Second)
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

			api.CheckOneProxy(tag.Tag, colorBkp+"[BKP] "+colorReset)
		}

		aliveTags = api.GetAliveServers()

		time.Sleep(time.Duration(share.Settings.DelayBkp) * time.Second)
	}
}

func testAllTagsRoutine() {
	for {
		testAllTags()

		aliveTags = api.GetAliveServers()

		time.Sleep(time.Duration(share.Settings.DelayAll) * time.Second)
	}
}

func testAllTags() {

	tags := api.GetAllTags()
	l := len(tags)

	for i, tag := range tags {

		if tag == currentProxy {
			continue
		}

		api.CheckOneProxy(tag, "["+strconv.Itoa(i+1)+"-"+strconv.Itoa(l)+"]")
	}
}

func checkSwitch() {
	var curDelay int
	var betterProxy api.ProxyServer
	var found bool

	for {
		found = false
		for _, tag := range aliveTags {

			if tag.Tag == currentProxy {
				curDelay = tag.Delay
				break
			} else if !found {
				betterProxy = tag
				found = true
			}
		}

		if found && betterProxy.Delay < (curDelay+50) {
			log.Println(colorWarn+"WARN "+colorMain+"[MAIN] "+colorReset+"Switching to faster proxy:", betterProxy.Tag)
			switchProxy(betterProxy.Tag)
		}

		time.Sleep(time.Duration(share.Settings.DelaySwitch) * time.Second)
	}
}
