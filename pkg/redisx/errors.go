package redisx

import "errors"

// ErrKeyNotFound is returned when a key doesn't exist
var ErrKeyNotFound = errors.New("key not found")
