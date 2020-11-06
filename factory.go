package log

// LoggerFactory is a factory to create a logger on demand
type LoggerFactory interface {
	Make(level Level) Logger
}
