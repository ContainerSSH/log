package log

// LabelName is a name for a Message label. Can only contain A-Z, a-z, 0-9, -, _.
type LabelName string

// LabelValue is a string, int, bool, or float.
type LabelValue interface{}

// Labels is a map linking
type Labels map[LabelName]LabelValue

// Logger The logger interface provides logging facilities on various levels.
type Logger interface {
	// WithLevel returns a copy of the logger for a specified log level. Panics if the log level provided is invalid.
	WithLevel(level Level) Logger
	// WithLabel returns a logger with an added label (e.g. username, IP, etc.) Panics if the label name is empty.
	WithLabel(labelName LabelName, labelValue LabelValue) Logger

	// Debug logs a message at the debug level.
	Debug(message ...interface{})
	// Debugf logs a message at the debug level with a formatting string.
	// Deprecated: use Debug with a Message instead.
	Debugf(format string, args ...interface{})

	// Info logs a message at the info level.
	Info(message ...interface{})
	// Infof logs a message at the info level with a formatting string.
	// Deprecated: use Info with a Message instead.
	Infof(format string, args ...interface{})

	// Notice logs a message at the notice level.
	Notice(message ...interface{})
	// Noticef logs a message at the notice level with a formatting string.
	// Deprecated: use Notice with a Message instead.
	Noticef(format string, args ...interface{})

	// Warning logs a message at the warning level.
	Warning(message ...interface{})
	// Warningf logs a message at the warning level with a formatting string.
	// Deprecated: use Warning with a Message instead.
	Warningf(format string, args ...interface{})

	// Error logs a message at the error level.
	Error(message ...interface{})
	// Errorf logs a message at the error level with a formatting string.
	// Deprecated: use Error with a Message instead.
	Errorf(format string, args ...interface{})

	// Critical logs a message at the critical level.
	Critical(message ...interface{})
	// Criticalf logs a message at the critical level with a formatting string.
	// Deprecated: use Critical with a Message instead.
	Criticalf(format string, args ...interface{})

	// Alert logs a message at the alert level.
	Alert(message ...interface{})
	// Alertf logs a message at the alert level with a formatting string
	// Deprecated: use Alert with a Message instead.
	Alertf(format string, args ...interface{})

	// Emergency logs a message at the emergency level.
	Emergency(message ...interface{})
	// Emergencyf logs a message at the emergency level with a formatting string.
	// Deprecated: use Emergency with a Message instead.
	Emergencyf(format string, message ...interface{})

	// Log logs a number of objects or strings to the log.
	Log(v ...interface{})
	// Logf formats a message and logs it.
	Logf(format string, v ...interface{})

	// Rotate triggers the logging backend to close all connections and reopen them to allow for rotating log files.
	Rotate() error
	// Close closes the logging backend.
	Close() error
}

// LoggerFactory is a factory to create a logger on demand
type LoggerFactory interface {
	// Make creates a new logger with the specified configuration and module.
	//
	// - config is the configuration structure.
	//
	// Return:
	//
	// - Logger is the logger created.
	// - error is returned if the configuration was invalid.
	Make(config Config) (Logger, error)

	// MustMake is identical to Make but panics if an error happens
	MustMake(config Config) Logger
}
