package shadowsocks

import (
	"encoding/base64"
	"net"
	"net/url"
	"strconv"
)

// ToURL converts Shadowsocks struct to URL string
func ToURL(h Shadowsocks) string {

	cred := h.Method + ":" + h.Password
	encoded := base64.RawURLEncoding.EncodeToString([]byte(cred))

	u := &url.URL{
		Scheme:   "ss",
		User:     url.User(encoded),
		Host:     net.JoinHostPort(h.Server, strconv.Itoa(h.Port)),
		Fragment: h.Tag,
	}

	return u.String()
}
