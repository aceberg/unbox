package hysteria2

import (
	"net"
	"net/url"
	"strconv"

	"github.com/aceberg/unbox/pkg/tls"
)

// ToURL converts Hysteria2 struct to URL string
func ToURL(h Hysteria2) string {
	u := &url.URL{
		Scheme:   "hysteria2",
		User:     url.User(h.Password),
		Host:     net.JoinHostPort(h.Server, strconv.Itoa(h.Port)),
		Fragment: h.Tag,
	}

	q := tls.ToURL(h.TLS)

	u.RawQuery = q.Encode()

	return u.String()
}
