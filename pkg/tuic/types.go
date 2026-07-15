package tuic

import "github.com/aceberg/unbox/pkg/tls"

// TUIC config
type TUIC struct {
	Type     string   `json:"type"`
	Tag      string   `json:"tag"`
	Server   string   `json:"server"`
	Port     int      `json:"server_port"`
	UUID     string   `json:"uuid"`
	Password string   `json:"password,omitempty"`
	TLS      *tls.TLS `json:"tls,omitempty"`
}
