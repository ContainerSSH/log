package log

import (
	"io"
)

// Writer is a specialized writer to write a line of log messages.
type Writer interface {
	// Write writes a log message to the output.
	Write(level Level, message Message) error
	// Rotate attempts to rotate the logs. Has no effect on non-file based loggers.
	Rotate() error

	io.Closer
}
