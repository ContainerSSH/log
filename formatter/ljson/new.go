package ljson

import "github.com/containerssh/log/formatter"

// NewLJsonLogFormatter Factory for the newline-delimited JSON formatter
func NewLJsonLogFormatter() formatter.Formatter {
	return &LogFormatter{}
}
