package vless

import (
	"errors"
	"net/url"
	"strconv"

	"github.com/aceberg/unbox/internal/tls"
	"github.com/aceberg/unbox/internal/transport"
)

// ParseVLESS converts VLESS URL string to struct
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
		Type:    "vless",
		Tag:     u.Fragment,
		Server:  u.Hostname(),
		Port:    portInt,
		UUID:    u.User.Username(),
		Flow:    q.Get("flow"),
		PackEnc: "xudp",
	}

	if res.Server == "" || res.Port == 0 || res.UUID == "" {
		return nil, errors.New("required field empty in " + raw)
	}

	if res.Flow != "" && res.Flow != "xtls-rprx-vision" {
		return nil, errors.New("unsupported flow: " + res.Flow)
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
