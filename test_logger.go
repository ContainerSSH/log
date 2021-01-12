package log

import (
	"strings"
	"testing"
)

// GetTestLogger creates a logger for testing purposes.
//goland:noinspection GoUnusedExportedFunction
func GetTestLogger(t *testing.T) Logger {
	writer := &testLogWriter{
		t: t,
	}
	logger, err := New(
		Config{
			Level:  LevelDebug,
			Format: FormatText,
		},
		t.Name(),
		writer,
	)
	if err != nil {
		panic(err)
	}
	return logger
}

type testLogWriter struct {
	t *testing.T
}

func (t *testLogWriter) Write(p []byte) (n int, err error) {
	t.t.Log(strings.TrimSpace(string(p)))
	return len(p), nil
}
