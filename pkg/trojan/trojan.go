package trojan

import (
	"errors"
	"net/url"
	"strconv"

	"github.com/aceberg/unbox/pkg/tls"
	"github.com/aceberg/unbox/pkg/transport"
)

// Parse converts Trojan URL string to struct
func Parse(raw string) (*Trojan, error) {
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
	}

	if res.Server == "" || res.Port == 0 || res.Password == "" {
		return nil, errors.New("required field empty in " + raw)
	}

	tr, ok := transport.Get(q)
	if ok {
		res.Trans = &tr
	}

	t, ok := tls.Get(q)
	if ok {
		res.TLS = &t
	}

	return res, nil
}
