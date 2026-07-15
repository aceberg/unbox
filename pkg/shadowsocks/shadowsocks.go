package shadowsocks

import (
	"encoding/base64"
	"errors"
	"net/url"
	"strconv"
	"strings"
)

// Parse converts Shadowsocks URL string to struct
func Parse(raw string) (*Shadowsocks, error) {
	u, err := url.Parse(raw)
	if err != nil {
		return nil, err
	}

	portInt, err := strconv.Atoi(u.Port())
	if err != nil {
		return nil, err
	}

	// Base64 URL-safe decoding (with padding fix)
	encoded := u.User.Username()
	if m := len(encoded) % 4; m != 0 {
		encoded += strings.Repeat("=", 4-m)
	}

	decoded, err := base64.RawURLEncoding.DecodeString(encoded)
	if err != nil {
		decoded, err = base64.StdEncoding.DecodeString(encoded)
		if err != nil {
			return nil, err
		}
	}
	mp := strings.SplitN(string(decoded), ":", 2)

	res := &Shadowsocks{
		Type:     "shadowsocks",
		Tag:      u.Fragment,
		Server:   u.Hostname(),
		Port:     portInt,
		Method:   mp[0],
		Password: mp[1],
	}

	if res.Server == "" || res.Port == 0 || res.Password == "" {
		return nil, errors.New("required field empty in " + raw)
	}

	return res, nil
}
