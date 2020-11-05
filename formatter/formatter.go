package formatter

import "github.com/containerssh/log"

// A formatter is an interface that can format a log message or a data interface into a data stream.
type Formatter interface {
	// Format a string message
	Format(level log.Level, message string) []byte
	// Format a data object
	FormatData(level log.Level, data interface{}) []byte
}
