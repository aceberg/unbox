package transport

import (
	"log"
	"net/url"
	"slices"
)

// Headers for Transport struct
type Headers struct {
	Host string `json:"Host"`
}

// Transport for VLESS or Trojan config struct
type Transport struct {
	Type     string   `json:"type"`
	Path     string   `json:"path,omitempty"`
	Head     *Headers `json:"headers,omitempty"`
	ServName string   `json:"service_name,omitempty"`
	IdleTout string   `json:"idle_timeout,omitempty"`
	PingTout string   `json:"ping_timeout,omitempty"`
}

// Get converts url.Values to Transport struct
func Get(q url.Values) (Transport, bool) {
	var res Transport

	tp := q.Get("type")

	if tp == "tcp" || tp == "" || tp == "raw" {
		return res, false
	}

	allowedTransport := []string{"ws", "quic", "grpc", "httpupgrade"}
	if !slices.Contains(allowedTransport, tp) {
		log.Println("ERROR: unsupported transport type", tp)
		return res, false
	}

	if tp == "quic" {
		res.Type = tp
	}

	if tp == "grpc" {
		res.Type = tp
		res.ServName = q.Get("serviceName")
		res.IdleTout = "15s"
		res.PingTout = "15s"
	}

	if tp == "ws" || tp == "httpupgrade" {
		var head *Headers
		if q.Get("host") != "" {
			head = &Headers{
				Host: q.Get("host"),
			}
		}
		res.Type = tp
		res.Head = head
		res.Path = q.Get("path")
	}

	return res, true
}
