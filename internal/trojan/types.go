package trojan

// TLS for Trojan config struct
type TLS struct {
	Enabled bool   `json:"enabled"`
	SNI     string `json:"server_name"`
}

// Headers for Transport struct
type Headers struct {
	Host string `json:"Host"`
}

// Transport for Trojan config struct
type Transport struct {
	Type string   `json:"type"`
	Path string   `json:"path,omitempty"`
	Head *Headers `json:"headers,omitempty"`
}

// Trojan config
type Trojan struct {
	Type     string    `json:"type"`
	Tag      string    `json:"tag"`
	Server   string    `json:"server"`
	Port     int       `json:"server_port"`
	Password string    `json:"password"`
	TLS      TLS       `json:"tls"`
	Trans    Transport `json:"transport"`
}
