package log

import (
	"testing"
)

// GetTestLogger creates a logger for testing purposes.
//goland:noinspection GoUnusedExportedFunction
func GetTestLogger(t *testing.T) Logger {
	logger, err := NewLogger(
		Config{
			Level:  LevelDebug,
			Format: FormatText,
			Output: OutputTest,
			T:      t,
		},
	)
	if err != nil {
		panic(err)
	}
	return logger
}
