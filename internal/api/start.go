package api

import "log"

// Conf contains command-line options for unbox
type Conf struct {
	ApiPath   string
	OutPath   string
	KeepAlive bool
	TestURL   string
}

// Config - app config
var Config Conf

func Start() {

	if Config.KeepAlive {
		keepConnectionAlive()

	} else {
		testAllTags()
		aliveTags := getAliveTags()
		if len(aliveTags) == 0 {
			log.Println("No proxies online. Exiting")
			return
		}

		editConfig(aliveTags)
	}
}
