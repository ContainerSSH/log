package ljson

import (
	"encoding/json"
	"time"

	"github.com/containerssh/log"
)

// Format formats a string message
func (formatter *LogFormatter) Format(level log.Level, message string) []byte {
	l, err := level.String()
	if err != nil {
		panic(err)
	}
	line, err := json.Marshal(JsonLine{
		Time:    time.Now().Format(time.RFC3339),
		Level:   l,
		Message: message,
	})
	if err != nil {
		panic(err)
	}
	return line
}
