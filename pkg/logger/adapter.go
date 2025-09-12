package logger

import (
	"go.uber.org/zap"
)

// Logger interface defines the logging contract
type Logger interface {
	Debug(message string, keysAndValues ...interface{})
	Info(message string, keysAndValues ...interface{})
	Warn(message string, keysAndValues ...interface{})
	Error(message string, keysAndValues ...interface{})
	Fatal(message string, keysAndValues ...interface{})
}

// ZapAdapter adapts zap.Logger to our Logger interface
type ZapAdapter struct {
	logger *zap.Logger
	sugar  *zap.SugaredLogger
}

// NewZapAdapter creates a new ZapAdapter
func NewZapAdapter(logger *zap.Logger) *ZapAdapter {
	return &ZapAdapter{
		logger: logger,
		sugar:  logger.Sugar(),
	}
}

// Debug logs a debug message with key-value pairs
func (z *ZapAdapter) Debug(message string, keysAndValues ...interface{}) {
	z.sugar.Debugw(message, keysAndValues...)
}

// Info logs an info message with key-value pairs
func (z *ZapAdapter) Info(message string, keysAndValues ...interface{}) {
	z.sugar.Infow(message, keysAndValues...)
}

// Warn logs a warning message with key-value pairs
func (z *ZapAdapter) Warn(message string, keysAndValues ...interface{}) {
	z.sugar.Warnw(message, keysAndValues...)
}

// Error logs an error message with key-value pairs
func (z *ZapAdapter) Error(message string, keysAndValues ...interface{}) {
	z.sugar.Errorw(message, keysAndValues...)
}

// Fatal logs a fatal message with key-value pairs and calls os.Exit(1)
func (z *ZapAdapter) Fatal(message string, keysAndValues ...interface{}) {
	z.sugar.Fatalw(message, keysAndValues...)
}