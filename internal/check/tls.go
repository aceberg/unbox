package check

import (
	"crypto/tls"
	"time"

	"golang.org/x/net/proxy"
)

// HandshakeTLS verifies a server's TLS certificate through a SOCKS5 proxy.
func HandshakeTLS(proxyAddr, host string) error {
	dialer, err := proxy.SOCKS5("tcp", proxyAddr, nil, proxy.Direct)
	if err != nil {
		return err
	}

	conn, err := dialer.Dial("tcp", host+":443")
	if err != nil {
		return err
	}

	IfError(conn.SetDeadline(time.Now().Add(15 * time.Second)))

	tlsConn := tls.Client(conn, &tls.Config{
		ServerName: host,
	})

	err = tlsConn.Handshake()

	IfError(tlsConn.Close())

	return err
}
