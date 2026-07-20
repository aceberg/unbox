package trojan

import (
	"net"
	"net/url"
	"strconv"

	"github.com/aceberg/unbox/internal/check"
)

// ToURL converts Trojan struct to URL string
func (h Trojan) ToURL() string {
	u := &url.URL{
		Scheme:   "trojan",
		User:     url.User(h.Password),
		Host:     net.JoinHostPort(h.Server, strconv.Itoa(h.Port)),
		Fragment: h.Tag,
	}

	q := check.MergeURLValues(h.TLS.ToValues(), h.Trans.ToValues())

	u.RawQuery = q.Encode()

	return u.String()
}
