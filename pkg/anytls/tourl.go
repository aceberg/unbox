package anytls

import (
	"net"
	"net/url"
	"strconv"
)

// ToURL converts AnyTLS struct to URL string
func (h AnyTLS) ToURL() string {
	u := &url.URL{
		Scheme:   "anytls",
		User:     url.User(h.Password),
		Host:     net.JoinHostPort(h.Server, strconv.Itoa(h.Port)),
		Fragment: h.Tag,
	}

	q := h.TLS.ToValues()

	u.RawQuery = q.Encode()

	return u.String()
}
