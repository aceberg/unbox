package keep

import (
	"log"

	"github.com/aceberg/unbox/internal/api"
	"github.com/aceberg/unbox/internal/check"
	"github.com/aceberg/unbox/internal/share"
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

	select {}
}

func switchProxy(tag string) {

	selName := api.GetSelectorName()
	if selName == "" {
		log.Println(share.Col.Err + "ERROR" + share.Col.Reset + "Can't get Selector tag name to switch proxy")
		return
	}

	log.Println(share.Col.Warn+"WARN "+share.Col.Main+"[MAIN] "+share.Col.Reset+"Selecting proxy:", tag)

	err := api.SwitchProxy(selName, tag)
	check.IfError(err)
}

func chooseTag() (string, bool) {

	aliveTags := api.GetAliveServers()
	currentProxy := api.GetCurrntProxy()

	for _, tag := range aliveTags {
		if tag.Tag != currentProxy {
			return tag.Tag, true
		}
	}

	return "", false
}
