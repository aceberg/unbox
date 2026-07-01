package hysteria2

import "github.com/aceberg/unbox/internal/tls"

// Hysteria2 config
type Hysteria2 struct {
	Type     string   `json:"type"`
	Tag      string   `json:"tag"`
	Server   string   `json:"server"`
	Port     int      `json:"server_port"`
	Password string   `json:"password,omitempty"`
	TLS      *tls.TLS `json:"tls"`
}
