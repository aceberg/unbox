package keep

import (
	"log"
	"strings"
	"time"

	"github.com/aceberg/unbox/internal/api"
	"github.com/aceberg/unbox/internal/check"
	"github.com/aceberg/unbox/internal/share"
)

var (
	currentProxy string
	alive        bool
	aliveTags    []api.ProxyServer
)

// Alive - keep alive and auto switch
func Alive() {

	if share.Settings.DelayMain > 0 {
		go testCurrentProxy()
	}
	if share.Settings.DelayBkp > 0 {
		go testBackup()
	}
	if share.Settings.DelayAll > 0 {
		go testAllTagsRoutine()
	}
	if share.Settings.DelaySwitch > 0 {
		go testFasterProxy()
	}

	for {
		if !alive {
			tag, ok := chooseTag()
			if ok {
				alive = true
				switchProxy(tag)
			} else {
				log.Println(share.Col.Err + "ERROR" + share.Col.Main + "[MAIN] " + share.Col.Reset + "No proxies online!")
			}
		}

		time.Sleep(time.Duration(1) * time.Second)
	}
}

func switchProxy(tag string) {

	selName := api.GetSelectorName()
	if selName == "" {
		log.Println(share.Col.Err + "ERROR" + share.Col.Main + "[MAIN] " + share.Col.Reset + "Can't get Selector tag name to select new proxy")
		return
	}

	log.Println(share.Col.Warn+"WARN "+share.Col.Main+"[MAIN] "+share.Col.Reset+"Selecting proxy:", tag)
	currentProxy = tag

	body := strings.NewReader(`{"name":"` + tag + `"}`)

	_, err := api.Request("PUT", "/proxies/"+selName, body)
	check.IfError(err)
}

func chooseTag() (string, bool) {

	aliveTags = api.GetAliveServers()

	for _, tag := range aliveTags {
		if tag.Tag != currentProxy {
			return tag.Tag, true
		}
	}

	return "", false
}
