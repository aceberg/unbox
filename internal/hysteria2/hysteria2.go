package hysteria2

import (
	"net/url"
	"strconv"
)

// ParseHyst2 converts Hysteria2 URL to struct
func ParseHyst2(raw string) (*Hysteria2, error) {
	u, err := url.Parse(raw)
	if err != nil {
		return nil, err
	}

	portInt, err := strconv.Atoi(u.Port())
	if err != nil {
		return nil, err
	}

	q := u.Query()

	res := &Hysteria2{
		Type:     "hysteria2",
		Tag:      u.Fragment,
		Server:   u.Hostname(),
		Port:     portInt,
		Password: u.User.Username(),
		TLS: TLS{
			Enabled:  true,
			SNI:      q.Get("sni"),
			Insecure: true,
		},
	}

	return res, nil
}
