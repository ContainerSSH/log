package log

// NewLJsonLogFormatter Factory for the newline-delimited JSON formatter.
func NewLJsonLogFormatter() Formatter {
	return &ljsonLogFormatter{}
}
