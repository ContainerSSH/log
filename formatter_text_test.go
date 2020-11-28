package log_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/containerssh/log"
)

func TestTextLogger(t *testing.T) {
	formatter := log.NewTextLogFormatter()

	message := formatter.Format(log.LevelDebug, "test", "Hello world!")
	messageParts := strings.Split(string(message), "\t")
	assert.Equal(t, "debug", messageParts[1])
	assert.Equal(t, "test", messageParts[2])
	assert.Equal(t, "Hello world!\n", messageParts[3])
}
