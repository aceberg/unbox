package trojan

import (
	"net/url"
	"strconv"
)

// ParseTrojan converts Trojan URL to struct
func ParseTrojan(raw string) (*Trojan, error) {
	u, err := url.Parse(raw)
	if err != nil {
		return nil, err
	}

	portInt, err := strconv.Atoi(u.Port())
	if err != nil {
		return nil, err
	}

	q := u.Query()

	res := &Trojan{
		Type:     "trojan",
		Tag:      u.Fragment,
		Server:   u.Hostname(),
		Port:     portInt,
		Password: u.User.Username(),
		TLS: TLS{
			Enabled: true,
			SNI:     q.Get("sni"),
		},
		Trans: Transport{
			Type: q.Get("type"),
			Path: q.Get("path"),
			Head: &Headers{
				Host: q.Get("host"),
			},
		},
	}

	return res, nil
}
