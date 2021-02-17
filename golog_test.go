package log_test

import (
	"bytes"
	goLog "log"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/containerssh/log"
)

func TestGoLog(t *testing.T) {
	writer := &bytes.Buffer{}
	logger := log.MustNewLogger(
		log.Config{
			Level:  log.LevelDebug,
			Format: log.FormatText,
			Output: log.OutputStdout,
			Stdout: writer,
		},
	)
	goLogWriter := log.NewGoLogWriter(logger)
	goLogger := goLog.New(goLogWriter, "", 0)
	goLogger.Printf("test")
	assert.True(t, len(writer.Bytes()) > 0)
}
