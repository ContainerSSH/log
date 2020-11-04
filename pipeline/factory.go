package pipeline

import (
	"io"

	"github.com/containerssh/log"
	"github.com/containerssh/log/factory"
	"github.com/containerssh/log/formatter"
)

type LoggerPipelineFactory struct {
	formatter formatter.Formatter
	writer io.Writer
}

func NewLoggerPipelineFactory(formatter formatter.Formatter) factory.LoggerFactory {
	return &LoggerPipelineFactory{
		formatter: formatter,
	}
}

func (f *LoggerPipelineFactory) Make(level log.Level) log.Logger {
	return NewLoggerPipeline(
		level,
		f.formatter,
		f.writer,
	)
}

