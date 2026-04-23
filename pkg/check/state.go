package check

import (
	"net"
	"strconv"
	"time"
)

// State - returns state of a service
func State(addr string, port int) bool {

	timeout := 3 * time.Second
	target := addr + ":" + strconv.Itoa(port)

	conn, err := net.DialTimeout("tcp", target, timeout)
	if err != nil {
		return false
	}
	err = conn.Close()
	IfError(err)

	return true
}
