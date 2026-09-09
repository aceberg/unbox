package check

import (
	"strings"

	"github.com/google/uuid"
)

func ValidateUUID(id string) error {

	_, err := uuid.Parse(strings.TrimSpace(id))
	return err
}
