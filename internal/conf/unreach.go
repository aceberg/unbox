package conf

import (
	"log"
	"strconv"

	"github.com/aceberg/unbox/internal/api"
)

// RemoveUnreachable removes unreachable nodes from sing-box config
func RemoveUnreachable() {

	checkAllTags()
	aliveServers := api.GetAliveServers()
	if len(aliveServers) == 0 {
		log.Println("No proxies online. Exiting")
		return
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
