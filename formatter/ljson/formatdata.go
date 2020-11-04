package ljson

import (
	"encoding/json"
	"github.com/containerssh/log"
	"time"
)

func (formatter *LogFormatter) FormatData(level log.Level, data interface{}) []byte {
	l, err := level.String()
	if err != nil {
		panic(err)
	}
	line, err := json.Marshal(JsonLine{
		Time:    time.Now().Format(time.RFC3339),
		Level:   l,
		Details: data,
	})
	if err != nil {
		panic(err)
	}
	return line
}
