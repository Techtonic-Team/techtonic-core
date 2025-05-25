package main

import (
	"gitlab.com/techtonic-team/tdk/techtonic-core/pkg/logging/seismic"
)

func main() {
	logger := seismic.NewLogger("console")
	defer logger.Sync()

	logger.Info("Hello, World!", "greeting", "Hello", "target", "World")
	logger.Info("This is a test log message", "level", "info", "context", "example")
	logger.Info("Logging with key-value pairs", "key1", "value1", "key2", 42)
}
