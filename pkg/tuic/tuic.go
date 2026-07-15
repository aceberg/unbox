package tuic

import (
	"errors"
	"net/url"
	"strconv"

	"github.com/aceberg/unbox/pkg/tls"
)

// Parse converts TUIC URL string to struct
func Parse(raw string) (*TUIC, error) {
	u, err := url.Parse(raw)
	if err != nil {
		return nil, err
	}

	portInt, err := strconv.Atoi(u.Port())
	if err != nil {
		return nil, err
	}

	q := u.Query()

	res := &TUIC{
		Type:     "tuic",
		Tag:      u.Fragment,
		Server:   u.Hostname(),
		Port:     portInt,
		UUID:     u.User.Username(),
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

	if res.Server == "" || res.Port == 0 || res.UUID == "" {
		return nil, errors.New("required field empty in " + raw)
	}

	return res, nil
}
