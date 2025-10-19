package logger

import (
	"go.uber.org/zap"
)

// Logger is a facade for zap.Logger
type Logger = zap.Logger

// Config is a facade for zap.Config
type Config = zap.Config

// Field is a facade for zap.Field
type Field = zap.Field

// NewProduction creates a production logger
func NewProduction() (*Logger, error) {
	return zap.NewProduction()
}

// NewDevelopment creates a development logger
func NewDevelopment() (*Logger, error) {
	return zap.NewDevelopment()
}

// New creates a logger with the given config
func New(cfg Config) (*Logger, error) {
	return cfg.Build()
}

// Field constructors
var (
	String   = zap.String
	Int      = zap.Int
	Int64    = zap.Int64
	Float64  = zap.Float64
	Bool     = zap.Bool
	Error    = zap.Error
	Any      = zap.Any
	Duration = zap.Duration
	Time     = zap.Time
)

// NewNop creates a no-op logger
func NewNop() *Logger {
	return zap.NewNop()
}
