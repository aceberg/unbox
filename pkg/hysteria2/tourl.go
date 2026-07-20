package hysteria2

import (
	"net"
	"net/url"
	"strconv"
)

// ToURL converts Hysteria2 struct to URL string
func (h Hysteria2) ToURL() string {
	u := &url.URL{
		Scheme:   "hysteria2",
		User:     url.User(h.Password),
		Host:     net.JoinHostPort(h.Server, strconv.Itoa(h.Port)),
		Fragment: h.Tag,
	}

	q := h.TLS.ToValues()

	u.RawQuery = q.Encode()

	return u.String()
}
