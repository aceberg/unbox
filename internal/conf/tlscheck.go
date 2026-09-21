package conf

import (
	"log"
	"strconv"

	"github.com/aceberg/unbox/internal/api"
	"github.com/aceberg/unbox/internal/check"
	"github.com/aceberg/unbox/internal/share"
)

func checkAllTLS() (int, []api.ProxyServer) {
	var alive []api.ProxyServer
	var msg string

	tags := api.GetAllTags()
	l := len(tags)
	lStr := strconv.Itoa(l)

	host := share.Settings.HostForTLS

	selName := api.GetSelectorName()
	if selName == "" {
		log.Println(share.Col.Err + "ERROR" + share.Col.Reset + "Can't get Selector tag name to switch proxy")
		return l, alive
	}

	log.Println("INFO: Started TLS connection check for host", host)

	for i, tag := range tags {

		err := api.SwitchProxy(selName, tag)
		if check.IfError(err) {
			continue
		}

		err = check.HandshakeTLS(share.Settings.ProxyForTLS, host)
		if err == nil {
			alive = append(alive, api.ProxyServer{Tag: tag, Delay: 0})
			msg = share.Col.Ok + "\"ok\"" + share.Col.Reset
		} else {
			msg = err.Error()
		}

		logPref := "[" + strconv.Itoa(i+1) + "-" + lStr + "]"
		log.Print("INFO "+logPref+" \""+tag+"\": ", msg)
	}

	return l, alive
}
