package redis

import "errors"

// Domain errors for cache operations
var (
	// ErrNotFound is returned when a cache key is not found
	ErrNotFound = errors.New("cache: key not found")

	// ErrInvalidValue is returned when a cached value cannot be unmarshaled
	ErrInvalidValue = errors.New("cache: invalid value")

	// ErrConnectionFailed is returned when connection to cache fails
	ErrConnectionFailed = errors.New("cache: connection failed")
)
