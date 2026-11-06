package vmess

import (
	"net"
	"net/url"
	"strconv"

	"github.com/aceberg/unbox/internal/check"
)

// ToURL converts VMESS struct to URL string
func (h VMESS) ToURL() string {
	u := &url.URL{
		Scheme:   "vmess",
		User:     url.User(h.UUID),
		Host:     net.JoinHostPort(h.Server, strconv.Itoa(h.Port)),
		Fragment: h.Tag,
	}

	q := check.MergeURLValues(h.TLS.ToValues(), h.Trans.ToValues())

	u.RawQuery = q.Encode()

	return u.String()
}
