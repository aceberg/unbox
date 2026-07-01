package api

import "log"

// Conf contains command-line options for unbox
type Conf struct {
	APIPath     string
	APISecret   string
	OutPath     string
	InputPath   string
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

// ProxyServer - proxy name and delay
type ProxyServer struct {
	Tag   string
	Delay int
}

// Start choses func to run
func Start() {

	if Config.KeepAlive {
		keepConnectionAlive()
	} else if Config.Deduplicate {
		deduplicate()
	} else if Config.InputPath != "" {
		invertConf()
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
