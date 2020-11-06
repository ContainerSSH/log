package pipeline

import (
	"io"

	"github.com/containerssh/log"
	"github.com/containerssh/log/factory"
	"github.com/containerssh/log/formatter"
)

// LoggerPipelineFactory Create a pipeline logger
type LoggerPipelineFactory struct {
	formatter formatter.Formatter
	writer    io.Writer
}

// NewLoggerPipelineFactory Create a pipeline logger factory
//goland:noinspection GoUnusedExportedFunction
func NewLoggerPipelineFactory(formatter formatter.Formatter, writer io.Writer) factory.LoggerFactory {
	return &LoggerPipelineFactory{
		formatter: formatter,
		writer:    writer,
	}
}

// Make Create the pipeline
func (f *LoggerPipelineFactory) Make(level log.Level) log.Logger {
	return NewLoggerPipeline(
		level,
		f.formatter,
		f.writer,
	)
}
