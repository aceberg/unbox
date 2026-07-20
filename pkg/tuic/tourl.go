package tuic

import (
	"net"
	"net/url"
	"strconv"
)

// ToURL converts TUIC struct to URL string
func (h TUIC) ToURL() string {
	u := &url.URL{
		Scheme:   "tuic",
		User:     url.User(h.UUID),
		Host:     net.JoinHostPort(h.Server, strconv.Itoa(h.Port)),
		Fragment: h.Tag,
	}

	q := h.TLS.ToValues()

	u.RawQuery = q.Encode()

	return u.String()
}
