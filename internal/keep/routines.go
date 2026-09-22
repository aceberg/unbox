package keep

import (
	"log"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/aceberg/unbox/internal/api"
	"github.com/aceberg/unbox/internal/share"
)

var testingAll atomic.Bool

func waitForCurrentProxy() {

	for {
		tag, ok := chooseTag()
		if ok {
			switchProxy(tag)
			break
		} else {
			log.Println(share.Col.Err + "ERROR" + share.Col.Main + "[MAIN] " + share.Col.Reset + "No proxies online!")
			go testAllTags()
		}

		time.Sleep(time.Duration(1) * time.Second)
	}
}

func testCurrentProxy() {

	for {
		currentProxy := api.GetCurrntProxy()
		if currentProxy != "" {
			ok := api.CheckOneProxy(currentProxy, share.Col.Main+"[MAIN]"+share.Col.Reset)
			if !ok {
				waitForCurrentProxy()
			}
		}

		time.Sleep(time.Duration(share.Settings.DelayMain) * time.Second)
	}
}

func testBackup() {

	for {
		aliveTags := api.GetAliveServers()
		currentProxy := api.GetCurrntProxy()
		n := share.Settings.BackupN

		for i, tag := range aliveTags {

			if i >= n {
				break
			}

			if tag.Tag == currentProxy {
				n = n + 1
				continue
			}

			api.CheckOneProxy(tag.Tag, share.Col.Bkp+"[BKP] "+share.Col.Reset)
		}

		time.Sleep(time.Duration(share.Settings.DelayBkp) * time.Second)
	}
}

func testAllTagsRoutine() {
	for {
		testAllTags()

		time.Sleep(time.Duration(share.Settings.DelayAll) * time.Second)
	}
}

func testAllTags() {

	if !testingAll.CompareAndSwap(false, true) {
		return
	}
	defer testingAll.Store(false)

	tags := api.GetAllTags()
	total := strconv.Itoa(len(tags))

	for i, tag := range tags {

		api.CheckOneProxy(tag, "["+strconv.Itoa(i+1)+"-"+total+"]")
	}
}

func testFasterProxy() {
	var curDelay int
	var betterProxy api.ProxyServer
	var found bool

	for {
		aliveTags := api.GetAliveServers()
		currentProxy := api.GetCurrntProxy()
		found = false
		curDelay = 0
		betterProxy = api.ProxyServer{}

		for _, tag := range aliveTags {

			if tag.Tag == currentProxy {
				curDelay = tag.Delay
				break
			} else if !found {
				betterProxy = tag
				found = true
			}
		}
		diff := curDelay - betterProxy.Delay

		if found && diff > share.Settings.SwitchStep {
			log.Println(share.Col.Warn+"WARN "+share.Col.Main+"[MAIN] "+share.Col.Reset+"Switching to faster proxy:", betterProxy.Tag)
			switchProxy(betterProxy.Tag)
		}

		time.Sleep(time.Duration(share.Settings.DelaySwitch) * time.Second)
	}
}
