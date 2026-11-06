package vmess

import (
	"github.com/aceberg/unbox/pkg/tls"
	"github.com/aceberg/unbox/pkg/transport"
)

// VMESS config
type VMESS struct {
	Type     string               `json:"type"`
	Tag      string               `json:"tag"`
	Server   string               `json:"server"`
	Port     int                  `json:"server_port"`
	UUID     string               `json:"uuid"`
	Security string               `json:"security,omitempty"`
	TLS      *tls.TLS             `json:"tls,omitempty"`
	Trans    *transport.Transport `json:"transport,omitempty"`
	PackEnc  string               `json:"packet_encoding,omitempty"`
}
