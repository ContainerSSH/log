package log

// Formatter an interface that can format a log message or a data interface into a data stream.
type Formatter interface {
	// Format a string message. The module may be empty.
	Format(level Level, module string, message string) []byte
	// Format a data object. The module may be empty.
	FormatData(level Level, module string, data interface{}) []byte
}
