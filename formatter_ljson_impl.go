package log

import (
	"encoding/json"
	"fmt"
	"time"
)

// jsonLine is the line object in the LJSON format
type jsonLine struct {
	Time    string      `json:"timestamp"`
	Level   LevelString `json:"level"`
	Module  string      `json:"module,omitempty"`
	Message string      `json:"message,omitempty"`
	Details interface{} `json:"details,omitempty"`
}

type ljsonLogFormatter struct {
}

func (formatter *ljsonLogFormatter) Format(level Level, module string, message string) []byte {
	l, err := level.String()
	if err != nil {
		panic(err)
	}
	line, err := json.Marshal(jsonLine{
		Time:    time.Now().Format(time.RFC3339),
		Level:   l,
		Module:  module,
		Message: message,
		Details: nil,
	})
	if err != nil {
		panic(err)
	}
	line = append(line, '\n')
	return line
}

func (formatter *ljsonLogFormatter) FormatData(level Level, module string, data interface{}) []byte {
	l, err := level.String()
	if err != nil {
		panic(err)
	}
	line, err := json.Marshal(jsonLine{
		Time:    time.Now().Format(time.RFC3339),
		Level:   l,
		Module:  module,
		Message: "",
		Details: data,
	})
	if err != nil {
		return formatter.Format(level, module, fmt.Sprintf("%v", data))
	}
	line = append(line, '\n')
	return line
}
