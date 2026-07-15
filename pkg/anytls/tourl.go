package anytls

import (
	"net"
	"net/url"
	"strconv"

	"github.com/aceberg/unbox/pkg/tls"
)

// ToURL converts AnyTLS struct to URL string
func ToURL(h AnyTLS) string {
	u := &url.URL{
		Scheme:   "anytls",
		User:     url.User(h.Password),
		Host:     net.JoinHostPort(h.Server, strconv.Itoa(h.Port)),
		Fragment: h.Tag,
	}

	q := tls.ToValues(h.TLS)

	u.RawQuery = q.Encode()

	return u.String()
}
