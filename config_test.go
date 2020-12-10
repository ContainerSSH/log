package log_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"

	"github.com/containerssh/log"
)

func TestJSONDecode(t *testing.T) {
	cfg := "{\"level\":\"debug\",\"format\":\"text\"}"
	config := log.Config{}
	err := json.Unmarshal([]byte(cfg), &config)
	assert.NoError(t, err)
	assert.Equal(t, log.LevelDebug, config.Level)
	assert.Equal(t, log.FormatText, config.Format)
}

func TestYAMLDecode(t *testing.T) {
	cfg := "---\nlevel: debug\nformat: text\n"
	config := log.Config{}
	err := yaml.Unmarshal([]byte(cfg), &config)
	assert.NoError(t, err)
	assert.Equal(t, log.LevelDebug, config.Level)
	assert.Equal(t, log.FormatText, config.Format)
}