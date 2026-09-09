package check

import (
	"errors"
	"strconv"
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
