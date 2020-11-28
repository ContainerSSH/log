package log_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/containerssh/log"
)

func TestNewline(t *testing.T) {
	formatter := log.NewLJsonLogFormatter()
	message := formatter.Format(log.LevelDebug, "", "test")
	data := map[string]interface{}{}
	if err := json.Unmarshal(message, &data); err != nil {
		assert.Fail(t, "failed to unmarshal message", err)
		return
	}
	assert.Equal(t, "debug", data["level"])
	assert.Equal(t, "test", data["message"])
	assert.Equal(t, []byte("\n"), message[len(message)-1:])
}
