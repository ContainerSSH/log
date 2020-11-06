package ljson_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/containerssh/log"
	"github.com/containerssh/log/formatter/ljson"
)

func TestNewline(t *testing.T) {
	formatter := ljson.NewLJsonLogFormatter()
	message := formatter.Format(log.LevelDebug, "test")
	data := map[string]interface{}{}
	if err := json.Unmarshal(message, &data); err != nil {
		assert.Fail(t, "failed to unmarshal message", err)
		return
	}
	assert.Equal(t, "debug", data["level"])
	assert.Equal(t, "test", data["message"])
	assert.Equal(t, []byte("\n"), message[len(message)-1:])
}
