package trojan

import (
	"net"
	"net/url"
	"strconv"

	"github.com/aceberg/unbox/internal/check"
	"github.com/aceberg/unbox/pkg/tls"
	"github.com/aceberg/unbox/pkg/transport"
)

// ToURL converts Trojan struct to URL string
func ToURL(h Trojan) string {
	u := &url.URL{
		Scheme:   "trojan",
		User:     url.User(h.Password),
		Host:     net.JoinHostPort(h.Server, strconv.Itoa(h.Port)),
		Fragment: h.Tag,
	}

	q := check.MergeURLValues(tls.ToValues(h.TLS), transport.ToValues(h.Trans))

	u.RawQuery = q.Encode()

	return u.String()
}
