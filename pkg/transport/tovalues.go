package transport

import "net/url"

// ToValues converts Transport struct to url.Values
func (tr *Transport) ToValues() url.Values {
	q := url.Values{}

	if tr == nil {
		return q
	}

	switch tr.Type {
	case "quic":
		q.Set("type", "quic")

	case "grpc":
		q.Set("type", "grpc")
		if tr.ServName != "" {
			q.Set("serviceName", tr.ServName)
		}

	case "ws", "httpupgrade":
		q.Set("type", tr.Type)
		if tr.Path != "" {
			q.Set("path", tr.Path)
		}
		if tr.Head != nil && tr.Head.Host != "" {
			q.Set("host", tr.Head.Host)
		}

	}

	return q
}
