package log

// NewTextLogFormatter Factory for the text format.
func NewTextLogFormatter() Formatter {
	return &textLogFormatter{}
}
