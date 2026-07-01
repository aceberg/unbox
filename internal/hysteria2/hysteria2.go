package hysteria2

import (
	"errors"
	"net/url"
	"strconv"

	"github.com/aceberg/unbox/internal/tls"
)

// ParseHyst2 converts Hysteria2 URL string to struct
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
	}

	t, ok := tls.Get(q)
	if ok {
		res.TLS = &t
	}

	if !ok || res.Server == "" || res.Port == 0 {
		return nil, errors.New("required field empty in " + raw)
	}

	return res, nil
}
