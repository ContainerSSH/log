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
	logger := log.NewLoggerPipeline(log.LevelInfo, "", log.NewLJsonLogFormatter(), writer)
	goLogWriter := log.NewGoLogWriter(logger)
	goLogger := goLog.New(goLogWriter, "", 0)
	goLogger.Printf("test")
	assert.True(t, len(writer.String()) > 0)
}
