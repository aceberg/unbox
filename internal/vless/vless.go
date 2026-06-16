package vless

import (
	"errors"
	"net/url"
	"slices"
	"strconv"
)

// ParseVLESS converts VLESS URL to struct
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

	var tls *TLS
	if q.Get("sni") != "" {
		tls = &TLS{
			Enabled: true,
			SNI:     q.Get("sni"),
		}
	}

	res := &VLESS{
		Type:    "vless",
		Tag:     u.Fragment,
		Server:  u.Hostname(),
		Port:    portInt,
		UUID:    u.User.Username(),
		Flow:    q.Get("flow"),
		TLS:     tls,
		PackEnc: "xudp",
	}

	if res.Flow != "" && res.Flow != "xtls-rprx-vision" {
		return nil, errors.New("Unsupported flow: " + res.Flow)
	}

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

	allowedTransport := []string{"http", "ws", "quic", "grpc", "httpupgrade"}

	if !slices.Contains(allowedTransport, res.Trans.Type) {
		res.Trans = nil
	}

	sec := q.Get("security")
	if sec == "reality" {
		tls := res.TLS

		if tls == nil {
			tls = &TLS{
				Enabled: true,
			}
		}

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
