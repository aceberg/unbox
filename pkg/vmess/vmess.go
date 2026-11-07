package vmess

import (
	"errors"
	"net/url"
	"strings"

	"github.com/aceberg/unbox/internal/check"
	"github.com/aceberg/unbox/pkg/tls"
	"github.com/aceberg/unbox/pkg/transport"
)

// Parse converts VMESS URL string to struct
func Parse(raw string) (*VMESS, error) {

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

	res := &VMESS{
		Type:       "vmess",
		Tag:        u.Fragment,
		Server:     u.Hostname(),
		Port:       portInt,
		UUID:       id,
		Security:   "auto",
		AuthLength: true,
		PackEnc:    "xudp",
	}

	if res.Server == "" || res.Port == 0 || res.UUID == "" {
		return nil, errors.New("required field empty in " + raw)
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
