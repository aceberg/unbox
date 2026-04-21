package vless

import (
	"net/url"
	"strconv"
)

type UTLS struct {
	Enabled bool   `json:"enabled"`
	Finger  string `json:"fingerprint"`
}

type Reality struct {
	Enabled bool   `json:"enabled"`
	Key     string `json:"public_key"`
	ID      string `json:"short_id"`
}

type TLS struct {
	Enabled bool     `json:"enabled"`
	SNI     string   `json:"server_name"`
	Utls    *UTLS    `json:"utls,omitempty"`
	Real    *Reality `json:"reality,omitempty"`
}

type Headers struct {
	Host string `json:"Host"`
}

type Transport struct {
	Type     string   `json:"type"`
	Path     string   `json:"path,omitempty"`
	Head     *Headers `json:"headers,omitempty"`
	ServName string   `json:"service_name,omitempty"`
}

type VLESS struct {
	Type    string     `json:"type"`
	Tag     string     `json:"tag"`
	Server  string     `json:"server"`
	Port    int        `json:"server_port"`
	UUID    string     `json:"uuid"`
	Flow    string     `json:"flow,omitempty"`
	Tls     TLS        `json:"tls"`
	Trans   *Transport `json:"transport,omitempty"`
	PackEnc string     `json:"packet_encoding"`
}

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
		Tls: TLS{
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
	}
	if sec == "reality" {
		tls := res.Tls
		tls.Real = &Reality{
			Enabled: true,
			Key:     q.Get("pbk"),
			ID:      q.Get("sid"),
		}
		tls.Utls = &UTLS{
			Enabled: true,
			Finger:  q.Get("fp"),
		}
		res.Tls = tls
	}

	return res, nil
}
