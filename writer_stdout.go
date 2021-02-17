package log

import (
	"io"
	"sync"
)

// newStdoutWriter creates a log writer that writes to the stdout (io.Writer) in the specified format.
func newStdoutWriter(stdout io.Writer, format Format) (Writer, error) {
	return &stdoutWriter{
		fileHandleWriter: newFileHandleWriter(stdout, format, &sync.Mutex{}),
	}, nil
}

// stdoutWriter inherits the write method from fileHandleWriter and writes to the stdout.
type stdoutWriter struct {
	*fileHandleWriter
}
