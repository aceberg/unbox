package check

import (
	"strings"

	"github.com/google/uuid"
)

// ValidateUUID returns error on wrong UUID
func ValidateUUID(id string) error {

	_, err := uuid.Parse(strings.TrimSpace(id))
	return err
}
