package log

import (
	"io"
)

// New creates a standard logger pipeline.
//
// - config is the configuration structure for the logger
// - module is a descriptor for the module that's logging. Can be empty.
// - writer is the target writer the logs should be sent to.
//
// goland:noinspection GoUnusedExportedFunction
func New(config Config, module string, writer io.Writer) (Logger, error) {
	return NewFactory(writer).Make(config, module)
}

// NewFactory creates a standard logger pipeline factory.
//
// - writer is the target writer.
//
// goland:noinspection GoUnusedExportedFunction
func NewFactory(writer io.Writer) LoggerFactory {
	return NewLoggerPipelineFactory(writer)
}

// LoggerFactory is a factory to create a logger on demand
type LoggerFactory interface {
	// Make creates a new logger with the specified configuration and module.
	//
	// - config is the configuration structure.
	// - module is the module that's logging. Can be empty.
	//
	// Return:
	//
	// - Logger is the logger created.
	// - error is returned if the configuration was invalid.
	Make(config Config, module string) (Logger, error)
}
