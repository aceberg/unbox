package hysteria2

// TLS for Hysteria2 config struct
type TLS struct {
	Enabled  bool   `json:"enabled"`
	SNI      string `json:"server_name"`
	Insecure bool   `json:"insecure"`
}

// Hysteria2 config
type Hysteria2 struct {
	Type     string `json:"type"`
	Tag      string `json:"tag"`
	Server   string `json:"server"`
	Port     int    `json:"server_port"`
	Password string `json:"password"`
	TLS      TLS    `json:"tls"`
}
