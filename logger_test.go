package log_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/containerssh/log"
)

func TestLogLevelFiltering(t *testing.T) {
	for logLevelInt := 0; logLevelInt < 8; logLevelInt++ {
		t.Run(fmt.Sprintf("filter=%s", log.Level(logLevelInt).MustName()), func(t *testing.T) {
			for writeLogLevelInt := 0; writeLogLevelInt < 8; writeLogLevelInt++ {
				logLevel := log.Level(logLevelInt)
				writeLogLevel := log.Level(writeLogLevelInt)
				t.Run(
					fmt.Sprintf("write=%s", log.Level(writeLogLevelInt).MustName()),
					func(t *testing.T) {
						testLevel(t, logLevel, writeLogLevel)
					},
				)
			}
		})
	}
}

func testLevel(t *testing.T, logLevel log.Level, writeLogLevel log.Level) {
	var buf bytes.Buffer
	p := log.MustNewLogger(log.Config{
		Level:       logLevel,
		Format:      log.FormatLJSON,
		Destination: log.DestinationStdout,
		Stdout:      &buf,
	})
	message := log.UserMessage("E_TEST", "test", "test")
	switch writeLogLevel {
	case log.LevelDebug:
		p.Debug(message)
	case log.LevelInfo:
		p.Info(message)
	case log.LevelNotice:
		p.Notice(message)
	case log.LevelWarning:
		p.Warning(message)
	case log.LevelError:
		p.Error(message)
	case log.LevelCritical:
		p.Critical(message)
	case log.LevelAlert:
		p.Alert(message)
	case log.LevelEmergency:
		p.Emergency(message)
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

		expectedLevel := writeLogLevel.String()
		assert.Equal(t, string(expectedLevel), data["level"])
		assert.Equal(t, "test", data["message"])
	}
}
