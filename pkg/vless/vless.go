package vless

import (
	"errors"
	"net/url"
	"strings"

	"github.com/aceberg/unbox/internal/check"
	"github.com/aceberg/unbox/pkg/tls"
	"github.com/aceberg/unbox/pkg/transport"
)

// Parse converts VLESS URL string to struct
func Parse(raw string) (*VLESS, error) {
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

	q := u.Query()

	res := &VLESS{
		Type:    "vless",
		Tag:     u.Fragment,
		Server:  u.Hostname(),
		Port:    portInt,
		UUID:    id,
		Flow:    q.Get("flow"),
		PackEnc: "xudp",
	}

	if res.Server == "" || res.Port == 0 || res.UUID == "" {
		return nil, errors.New("required field empty in " + raw)
	}

	if res.Flow != "" && res.Flow != "xtls-rprx-vision" {
		return nil, errors.New("unsupported flow: " + res.Flow)
	}

	t, ok := tls.Get(q)
	if ok {
		res.TLS = &t
	}

	tr, ok := transport.Get(q)
	if ok {
		res.Trans = &tr

		if res.Trans.Type == "quic" && res.TLS == nil {
			return nil, errors.New("TLS required for transport type quic: " + raw)
		}
	}

	return res, nil
}
