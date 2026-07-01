package trojan

import (
	"net"
	"net/url"
	"strconv"

	"github.com/aceberg/unbox/internal/check"
	"github.com/aceberg/unbox/internal/tls"
	"github.com/aceberg/unbox/internal/transport"
)

// ToURL converts Trojan struct to URL string
func ToURL(h Trojan) string {
	u := &url.URL{
		Scheme:   "trojan",
		User:     url.User(h.Password),
		Host:     net.JoinHostPort(h.Server, strconv.Itoa(h.Port)),
		Fragment: h.Tag,
	}

	q := check.SumURL(tls.ToURL(h.TLS), transport.ToURL(h.Trans))

	u.RawQuery = q.Encode()

	return u.String()
}
