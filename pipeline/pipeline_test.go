package pipeline_test

import (
	"bufio"
	"bytes"
	"encoding/json"
	"testing"

	"github.com/containerssh/log"
	"github.com/containerssh/log/formatter/ljson"
	"github.com/containerssh/log/pipeline"

	"github.com/stretchr/testify/assert"
)

func TestLogLevelFiltering(t *testing.T) {
	for logLevelInt := 0; logLevelInt < 8; logLevelInt++ {
		for writeLogLevelInt := 0; writeLogLevelInt < 8; writeLogLevelInt++ {
			logLevel := log.Level(logLevelInt)
			writeLogLevel := log.Level(writeLogLevelInt)
			testLevel(t, logLevel, writeLogLevel)
		}
	}
}

func testLevel(t *testing.T, logLevel log.Level, writeLogLevel log.Level) {
	var buf bytes.Buffer
	writer := bufio.NewWriter(&buf)
	p := pipeline.NewLoggerPipeline(logLevel, ljson.NewLJsonLogFormatter(), writer)
	switch writeLogLevel {
	case log.LevelDebug:
		p.Debug("test")
	case log.LevelInfo:
		p.Info("test")
	case log.LevelNotice:
		p.Notice("test")
	case log.LevelWarning:
		p.Warning("test")
	case log.LevelError:
		p.Error("test")
	case log.LevelCritical:
		p.Critical("test")
	case log.LevelAlert:
		p.Alert("test")
	case log.LevelEmergency:
		p.Emergency("test")
	}

	if err := writer.Flush(); err != nil {
		assert.Fail(t, "failed to flush writer", err)
	}
	if logLevel < writeLogLevel {
		assert.Equal(t, 0, buf.Len())
	} else {
		assert.NotEqual(t, 0, buf.Len())

		rawData := buf.Bytes()
		data := map[string]interface{}{}
		if err := json.Unmarshal(rawData, &data); err != nil {
			assert.Fail(t, "failed to unmarshal JSON from writer", err)
		}

		expectedLevel, _ := writeLogLevel.String()
		assert.Equal(t, string(expectedLevel), data["level"])
		assert.Equal(t, "test", data["message"])
	}
}
