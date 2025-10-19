package types

import "github.com/google/uuid"

// UUID is a facade for google/uuid
type UUID = uuid.UUID

// NewUUID generates a new random UUID
func NewUUID() UUID {
	return uuid.New()
}

// ParseUUID parses a UUID from string
func ParseUUID(s string) (UUID, error) {
	return uuid.Parse(s)
}

// MustParseUUID parses a UUID from string or panics
func MustParseUUID(s string) UUID {
	return uuid.MustParse(s)
}
