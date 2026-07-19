package conf

import (
	"log"
	"strconv"

	"github.com/aceberg/unbox/internal/api"
	"github.com/aceberg/unbox/internal/share"
)

// RemoveUnreachable removes unreachable nodes from sing-box config
func RemoveUnreachable() {

	checkAllTags()
	aliveServers := api.GetAliveServers()
	if len(aliveServers) == 0 {
		log.Println("No proxies online. Exiting")
		return
	}

	if share.Settings.BestN != 0 && len(aliveServers) > share.Settings.BestN {
		aliveServers = aliveServers[0:share.Settings.BestN]
	}

	editConfig(aliveServers)
}

func checkAllTags() {

	tags := api.GetAllTags()
	l := len(tags)

	for i, tag := range tags {

		api.CheckOneProxy(tag, "["+strconv.Itoa(i+1)+"-"+strconv.Itoa(l)+"]")
	}
}
