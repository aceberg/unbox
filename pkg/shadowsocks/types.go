package shadowsocks

// Shadowsocks config
type Shadowsocks struct {
	Type     string `json:"type"`
	Tag      string `json:"tag"`
	Server   string `json:"server"`
	Port     int    `json:"server_port"`
	Method   string `json:"method"`
	Password string `json:"password"`
}
