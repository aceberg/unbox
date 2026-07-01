package vless

import (
	"github.com/aceberg/unbox/pkg/tls"
	"github.com/aceberg/unbox/pkg/transport"
)

// VLESS config
type VLESS struct {
	Type    string               `json:"type"`
	Tag     string               `json:"tag"`
	Server  string               `json:"server"`
	Port    int                  `json:"server_port"`
	UUID    string               `json:"uuid"`
	Flow    string               `json:"flow,omitempty"`
	TLS     *tls.TLS             `json:"tls,omitempty"`
	Trans   *transport.Transport `json:"transport,omitempty"`
	PackEnc string               `json:"packet_encoding,omitempty"`
}
