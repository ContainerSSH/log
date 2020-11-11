package log_test

import (
	"bytes"
	goLog "log"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/containerssh/log"
	"github.com/containerssh/log/formatter/ljson"
	"github.com/containerssh/log/pipeline"
)

func TestGoLog(t *testing.T) {
	writer := &bytes.Buffer{}
	logger := pipeline.NewLoggerPipeline(log.LevelInfo, ljson.NewLJsonLogFormatter(), writer)
	goLogWriter := log.NewGoLogWriter(logger)
	goLogger := goLog.New(goLogWriter, "", 0)
	goLogger.Printf("test")
	assert.True(t, len(writer.String()) > 0)
}
