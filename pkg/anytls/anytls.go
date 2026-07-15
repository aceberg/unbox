package anytls

import (
	"errors"
	"net/url"
	"strconv"

	"github.com/aceberg/unbox/pkg/tls"
)

// Parse converts AnyTLS URL string to struct
func Parse(raw string) (*AnyTLS, error) {
	u, err := url.Parse(raw)
	if err != nil {
		return nil, err
	}

	portInt, err := strconv.Atoi(u.Port())
	if err != nil {
		return nil, err
	}

	q := u.Query()

	res := &AnyTLS{
		Type:     "anytls",
		Tag:      u.Fragment,
		Server:   u.Hostname(),
		Port:     portInt,
		Password: u.User.Username(),
	}

	t, ok := tls.Get(q)
	if ok {
		res.TLS = &t
	} else {
		res.TLS = &tls.TLS{
			Enabled:    true,
			DisableSNI: true,
		}
	}

	if res.Server == "" || res.Port == 0 || res.Password == "" {
		return nil, errors.New("required field empty in " + raw)
	}

	return res, nil
}
