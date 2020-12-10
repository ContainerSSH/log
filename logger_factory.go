package log

import (
	"io"
)

// NewLoggerPipelineFactory create a pipeline logger factory
//goland:noinspection GoUnusedExportedFunction
func NewLoggerPipelineFactory(writer io.Writer) LoggerFactory {
	return &loggerPipelineFactory{
		writer: writer,
	}
}

type loggerPipelineFactory struct {
	writer    io.Writer
}

func (f *loggerPipelineFactory) Make(config Config, module string) (Logger, error) {
	if err := config.Level.Validate(); err != nil {
		return nil, err
	}

	if err := config.Format.Validate(); err != nil {
		return nil, err
	}

	var formatter Formatter
	switch config.Format {
	case FormatLJSON:
		formatter = NewLJsonLogFormatter()
	case FormatText:
		formatter = NewTextLogFormatter()
	}

	return NewLoggerPipeline(config.Level, module, formatter, f.writer), nil
}
