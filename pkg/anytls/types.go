package anytls

import "github.com/aceberg/unbox/pkg/tls"

// AnyTLS config
type AnyTLS struct {
	Type     string   `json:"type"`
	Tag      string   `json:"tag"`
	Server   string   `json:"server"`
	Port     int      `json:"server_port"`
	Password string   `json:"password,omitempty"`
	TLS      *tls.TLS `json:"tls,omitempty"`
}
