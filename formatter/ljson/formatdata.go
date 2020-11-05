package ljson

import (
	"encoding/json"
	"time"

	"github.com/containerssh/log"
)

// FormatData formats a data object
func (formatter *LogFormatter) FormatData(level log.Level, data interface{}) []byte {
	l, err := level.String()
	if err != nil {
		panic(err)
	}
	line, err := json.Marshal(JsonLine{
		Time:    time.Now().Format(time.RFC3339),
		Level:   l,
		Message: "",
		Details: data,
	})
	if err != nil {
		panic(err)
	}
	return line
}
