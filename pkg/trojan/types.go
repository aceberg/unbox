package trojan

import (
	"github.com/aceberg/unbox/pkg/tls"
	"github.com/aceberg/unbox/pkg/transport"
)

// Trojan config
type Trojan struct {
	Type     string               `json:"type"`
	Tag      string               `json:"tag"`
	Server   string               `json:"server"`
	Port     int                  `json:"server_port"`
	Password string               `json:"password"`
	TLS      *tls.TLS             `json:"tls,omitempty"`
	Trans    *transport.Transport `json:"transport,omitempty"`
}
