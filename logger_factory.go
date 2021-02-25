package log

import (
	"io"
	"os"
)

// NewLogger creates a standard logger pipeline.
//
// - config is the configuration structure for the logger
//
// goland:noinspection GoUnusedExportedFunction
func NewLogger(config Config) (Logger, error) {
	return NewLoggerFactory().Make(config)
}

// MustNewLogger is identical to NewLogger, except that it panics instead of returning an error
func MustNewLogger(
	config Config,
) Logger {
	logger, err := NewLogger(config)
	if err != nil {
		panic(err)
	}
	return logger
}

// NewLoggerFactory create a pipeline logger factory
//goland:noinspection GoUnusedExportedFunction
func NewLoggerFactory() LoggerFactory {
	return &loggerFactory{}
}

type loggerFactory struct {
}

func (f *loggerFactory) MustMake(config Config) Logger {
	logger, err := f.Make(config)
	if err != nil {
		panic(err)
	}
	return logger
}

func (f *loggerFactory) Make(config Config) (Logger, error) {
	if err := config.Level.Validate(); err != nil {
		return nil, err
	}

	if err := config.Format.Validate(); err != nil {
		return nil, err
	}

	if err := config.Output.Validate(); err != nil {
		return nil, err
	}

	var writer Writer
	var err error = nil
	switch config.Output {
	case OutputFile:
		writer, err = newFileWriter(config.File, config.Format)
	case OutputStdout:
		var stdout io.Writer = os.Stdout
		if config.Stdout != nil {
			stdout = config.Stdout
		}
		writer, err = newStdoutWriter(stdout, config.Format)
	case OutputSyslog:
		writer, err = newSyslogWriter(config.Syslog, config.Format)
	case OutputTest:
		writer = newGoTest(config.T)
	}
	if err != nil {
		return nil, err
	}

	return &logger{
		level:  config.Level,
		labels: map[LabelName]LabelValue{},
		writer: writer,
	}, nil
}
