package api

import "log"

// Conf contains command-line options for unbox
type Conf struct {
	ApiPath string
	OutPath string
}

// Config - app config
var Config Conf

func Start() {

	aliveTags := getAliveTags()
	if len(aliveTags) == 0 {
		log.Println("No proxies online. Exiting")
		return
	}

	editConfig(aliveTags)
}
