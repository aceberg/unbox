package tls

import "net/url"

// ToValues converts TLS struct to url.Values
func (tl *TLS) ToValues() url.Values {
	q := url.Values{}

	if tl == nil {
		return q
	}

	if tl.SNI != "" {
		q.Set("sni", tl.SNI)
	}
	if tl.Insecure {
		q.Set("insecure", "1")
	}
	if tl.Utls != nil && tl.Utls.Finger != "" {
		q.Set("fp", tl.Utls.Finger)
	}
	if tl.Real != nil {
		q.Set("security", "reality")
		q.Set("pbk", tl.Real.Key)
		q.Set("sid", tl.Real.ID)
	}

	return q
}
