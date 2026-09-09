package tuic

import (
	"errors"
	"net/url"
	"strings"

	"github.com/aceberg/unbox/internal/check"
	"github.com/aceberg/unbox/pkg/tls"
)

// Parse converts TUIC URL string to struct
func Parse(raw string) (*TUIC, error) {
	u, err := url.Parse(raw)
	if err != nil {
		return nil, err
	}

	portInt, err := check.StringToPort(u.Port())
	if err != nil {
		return nil, err
	}

	id := strings.TrimSpace(u.User.Username())
	err = check.ValidateUUID(id)
	if err != nil {
		return nil, err
	}

	res := &TUIC{
		Type:     "tuic",
		Tag:      u.Fragment,
		Server:   u.Hostname(),
		Port:     portInt,
		UUID:     id,
		Password: u.User.Username(),
	}

	q := u.Query()

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
