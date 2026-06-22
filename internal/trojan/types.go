package trojan

import "github.com/aceberg/unbox/internal/transport"

// TLS for Trojan config struct
type TLS struct {
	Enabled bool   `json:"enabled"`
	SNI     string `json:"server_name"`
}

// Trojan config
type Trojan struct {
	Type     string               `json:"type"`
	Tag      string               `json:"tag"`
	Server   string               `json:"server"`
	Port     int                  `json:"server_port"`
	Password string               `json:"password"`
	TLS      TLS                  `json:"tls"`
	Trans    *transport.Transport `json:"transport,omitempty"`
}
