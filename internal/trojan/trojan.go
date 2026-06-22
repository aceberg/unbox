package trojan

import (
	"net/url"
	"strconv"

	"github.com/aceberg/unbox/internal/transport"
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
	}

	tr, ok := transport.Get(q)
	if ok {
		res.Trans = &tr
	}

	return res, nil
}
