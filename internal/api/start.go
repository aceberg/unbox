package api

import "log"

// Conf contains command-line options for unbox
type Conf struct {
	ApiPath     string
	OutPath     string
	KeepAlive   bool
	TestURL     string
	DelayMain   uint
	DelayBkp    uint
	DelayAll    uint
	DelaySwitch uint
	Deduplicate bool
}

// Config - app config
var Config Conf

func Start() {

	if Config.KeepAlive {
		keepConnectionAlive()
	} else if Config.Deduplicate {
		deduplicate()
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
