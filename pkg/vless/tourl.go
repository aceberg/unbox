package vless

import (
	"net"
	"net/url"
	"strconv"

	"github.com/aceberg/unbox/internal/check"
)

// ToURL converts VLESS struct to URL string
func (h VLESS) ToURL() string {
	u := &url.URL{
		Scheme:   "vless",
		User:     url.User(h.UUID),
		Host:     net.JoinHostPort(h.Server, strconv.Itoa(h.Port)),
		Fragment: h.Tag,
	}

	q := check.MergeURLValues(h.TLS.ToValues(), h.Trans.ToValues())

	if h.Flow != "" {
		q.Set("flow", h.Flow)
	}

	u.RawQuery = q.Encode()

	return u.String()
}
