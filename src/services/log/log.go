package log

import (
	"go.uber.org/zap"
)

var (
	// Log is the global logger instance
	Log *zap.SugaredLogger
)

// Register a logger
func Register(logger *zap.SugaredLogger) {
	Log = logger
}

// Named scopes a logger to a method or class name for easier tracking
func Named(name string) *zap.SugaredLogger {
	// Nil check for pre-initialization logging
	if Log == nil {
		l, _ := zap.NewDevelopment()
		return l.Sugar().Named(name)
	}

	return Log.Named(name)
}
