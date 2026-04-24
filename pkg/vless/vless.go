package vless

import (
	"net/url"
	"strconv"
)

// ParseVLESS - string to struct
func ParseVLESS(raw string) (*VLESS, error) {
	u, err := url.Parse(raw)
	if err != nil {
		return nil, err
	}

	portInt, err := strconv.Atoi(u.Port())
	if err != nil {
		return nil, err
	}

	q := u.Query()

	res := &VLESS{
		Type:   "vless",
		Tag:    u.Fragment,
		Server: u.Hostname(),
		Port:   portInt,
		UUID:   u.User.Username(),
		Flow:   q.Get("flow"),
		TLS: TLS{
			Enabled: true,
			SNI:     q.Get("sni"),
		},
		PackEnc: "xudp",
	}

	sec := q.Get("security")
	if sec == "tls" {
		var head *Headers
		if q.Get("host") != "" {
			head = &Headers{
				Host: q.Get("host"),
			}
		}

		res.Trans = &Transport{
			Type:     q.Get("type"),
			Path:     q.Get("path"),
			Head:     head,
			ServName: q.Get("serviceName"),
		}

		if res.Trans.Type == "xhttp" || res.Trans.Type == "tcp" {
			res.Trans = nil
		}
	}
	if sec == "reality" {
		tls := res.TLS
		tls.Real = &Reality{
			Enabled: true,
			Key:     q.Get("pbk"),
			ID:      q.Get("sid"),
		}
		tls.Utls = &UTLS{
			Enabled: true,
			Finger:  q.Get("fp"),
		}
		res.TLS = tls
	}

	return res, nil
}
