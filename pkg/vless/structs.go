package vless

// UTLS for TLS struct
type UTLS struct {
	Enabled bool   `json:"enabled"`
	Finger  string `json:"fingerprint"`
}

// Reality for TLS struct
type Reality struct {
	Enabled bool   `json:"enabled"`
	Key     string `json:"public_key"`
	ID      string `json:"short_id"`
}

// TLS for VLESS config struct
type TLS struct {
	Enabled bool     `json:"enabled"`
	SNI     string   `json:"server_name"`
	Utls    *UTLS    `json:"utls,omitempty"`
	Real    *Reality `json:"reality,omitempty"`
}

// Headers for Transport struct
type Headers struct {
	Host string `json:"Host"`
}

// Transport for VLESS config struct
type Transport struct {
	Type     string   `json:"type"`
	Path     string   `json:"path,omitempty"`
	Head     *Headers `json:"headers,omitempty"`
	ServName string   `json:"service_name,omitempty"`
}

// VLESS config
type VLESS struct {
	Type    string     `json:"type"`
	Tag     string     `json:"tag"`
	Server  string     `json:"server"`
	Port    int        `json:"server_port"`
	UUID    string     `json:"uuid"`
	Flow    string     `json:"flow,omitempty"`
	TLS     TLS        `json:"tls"`
	Trans   *Transport `json:"transport,omitempty"`
	PackEnc string     `json:"packet_encoding"`
}
