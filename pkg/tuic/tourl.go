package tuic

import (
	"net"
	"net/url"
	"strconv"

	"github.com/aceberg/unbox/pkg/tls"
)

// ToURL converts TUIC struct to URL string
func ToURL(h TUIC) string {
	u := &url.URL{
		Scheme:   "tuic",
		User:     url.User(h.UUID),
		Host:     net.JoinHostPort(h.Server, strconv.Itoa(h.Port)),
		Fragment: h.Tag,
	}

	q := tls.ToValues(h.TLS)

	u.RawQuery = q.Encode()

	return u.String()
}
