package shadowsocks

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net/url"
	"slices"
	"strings"

	"github.com/aceberg/unbox/internal/check"
)

// Parse converts Shadowsocks URL string to struct
func Parse(raw string) (*Shadowsocks, error) {
	u, err := url.Parse(raw)
	if err != nil {
		return nil, err
	}

	portInt, err := check.StringToPort(u.Port())
	if err != nil {
		return nil, err
	}

	if u.User == nil {
		return nil, errors.New("missing credentials in " + raw)
	}

	method := u.User.Username()
	password, _ := u.User.Password()

	q := u.Query()

	if q.Get("method") == "none" || q.Get("encryption") == "none" {
		method = "none"
		password = u.User.Username()
	} else if password == "" {
		method, password = base64Decode(u.User.Username())
	}

	allowedMethod := []string{"2022-blake3-aes-128-gcm", "2022-blake3-aes-256-gcm", "2022-blake3-chacha20-poly1305", "none", "aes-128-gcm", "aes-192-gcm", "aes-256-gcm", "chacha20-ietf-poly1305", "xchacha20-ietf-poly1305", "aes-128-ctr", "aes-192-ctr", "aes-256-ctr", "aes-128-cfb", "aes-192-cfb", "aes-256-cfb", "rc4-md5", "chacha20-ietf", "xchacha20"}

	if !slices.Contains(allowedMethod, method) {
		return nil, errors.New("unsupported method " + method + " in " + raw)
	}

	// EIH is only supported by AES-2022 methods.
	if method == "2022-blake3-chacha20-poly1305" && strings.Contains(password, ":") {
		return nil, errors.New("EIH is not supported with " + method + " in " + raw)
	}

	err = validate2022Key(method, password)
	if err != nil {
		return nil, err
	}

	res := &Shadowsocks{
		Type:     "shadowsocks",
		Tag:      u.Fragment,
		Server:   u.Hostname(),
		Port:     portInt,
		Method:   method,
		Password: password,
	}

	if res.Server == "" || res.Port == 0 || res.Password == "" || res.Method == "" {
		return nil, errors.New("required field empty in " + raw)
	}

	return res, nil
}

func base64Decode(encoded string) (string, string) {

	// Base64 URL-safe decoding (with padding fix)

	if m := len(encoded) % 4; m != 0 {
		encoded += strings.Repeat("=", 4-m)
	}

	decoded, err := base64.RawURLEncoding.DecodeString(encoded)
	if err != nil {
		decoded, err = base64.StdEncoding.DecodeString(encoded)
	}

	if err == nil {
		mp := strings.SplitN(string(decoded), ":", 2)

		if len(mp) == 2 {
			return mp[0], mp[1]
		}
	}

	return "", ""
}

func validate2022Key(method, password string) error {
	want := 0

	switch method {
	case "2022-blake3-aes-128-gcm":
		want = 16
	case "2022-blake3-aes-256-gcm", "2022-blake3-chacha20-poly1305":
		want = 32
	default:
		return nil // not a 2022 method
	}

	// EIH passwords have two base64 keys separated by ':'.
	for _, key := range strings.Split(password, ":") {
		b, err := base64.StdEncoding.DecodeString(key)
		if err != nil {
			return fmt.Errorf("invalid base64 key")
		}
		if len(b) != want {
			return fmt.Errorf("bad key length, required %d, got %d", want, len(b))
		}
	}

	return nil
}
