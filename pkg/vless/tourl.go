package vless

import (
	"net"
	"net/url"
	"strconv"

	"github.com/aceberg/unbox/internal/check"
	"github.com/aceberg/unbox/pkg/tls"
	"github.com/aceberg/unbox/pkg/transport"
)

// ToURL converts VLESS struct to URL string
func ToURL(h VLESS) string {
	u := &url.URL{
		Scheme:   "vless",
		User:     url.User(h.UUID),
		Host:     net.JoinHostPort(h.Server, strconv.Itoa(h.Port)),
		Fragment: h.Tag,
	}

	q := check.MergeURLValues(tls.ToValues(h.TLS), transport.ToValues(h.Trans))

	if h.Flow != "" {
		q.Set("flow", h.Flow)
	}

	u.RawQuery = q.Encode()

	return u.String()
}
