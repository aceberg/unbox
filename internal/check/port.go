package check

import (
	"errors"
	"net"
	"strconv"
	"time"
)

// StringToPort converts a string to a valid TCP/UDP port number
func StringToPort(s string) (int, error) {

	p, err := strconv.Atoi(s)
	if err == nil {
		if p > 0 && p < 65536 {
			return p, nil
		}

		err = errors.New("invalid port " + strconv.Itoa(p))
	}

	return 0, err
}

// PortReachable checks if server:port is reachable
func PortReachable(server string, port string) bool {
	address := net.JoinHostPort(server, port)

	conn, err := net.DialTimeout("tcp", address, 3*time.Second)
	if err != nil {
		return false
	}

	err = conn.Close()
	IfError(err)

	return true
}
