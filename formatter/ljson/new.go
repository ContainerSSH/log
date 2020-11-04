package ljson

import "github.com/containerssh/log/formatter"

func NewLJsonLogFormatter() formatter.Formatter {
	return &LogFormatter{}
}
