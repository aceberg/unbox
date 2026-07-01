package tls

import "net/url"

// UTLS for TLS struct
type UTLS struct {
	Enabled bool   `json:"enabled"`
	Finger  string `json:"fingerprint,omitempty"`
}

// Reality for TLS struct
type Reality struct {
	Enabled bool   `json:"enabled"`
	Key     string `json:"public_key,omitempty"`
	ID      string `json:"short_id,omitempty"`
}

// TLS config struct
type TLS struct {
	Enabled  bool     `json:"enabled"`
	SNI      string   `json:"server_name,omitempty"`
	Insecure bool     `json:"insecure,omitempty"`
	Utls     *UTLS    `json:"utls,omitempty"`
	Real     *Reality `json:"reality,omitempty"`
}

// Get converts url.Values to TLS struct
func Get(q url.Values) (TLS, bool) {
	var res TLS

	if q.Get("sni") != "" {
		res.Enabled = true
		res.SNI = q.Get("sni")
		if q.Get("insecure") == "1" {
			res.Insecure = true
		}
	}

	if q.Get("security") == "reality" {
		res.Enabled = true
		res.Real = &Reality{
			Enabled: true,
			Key:     q.Get("pbk"),
			ID:      q.Get("sid"),
		}
	}

	if q.Get("fp") != "" {
		res.Enabled = true
		res.Utls = &UTLS{
			Enabled: true,
			Finger:  q.Get("fp"),
		}
	}

	return res, res.Enabled
}
