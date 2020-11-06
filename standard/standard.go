package standard

import (
	"os"

	"github.com/containerssh/log"
	"github.com/containerssh/log/formatter/ljson"
	"github.com/containerssh/log/pipeline"
)

// New creates a standard logger pipeline
//goland:noinspection GoUnusedExportedFunction
func New() log.Logger {
	return pipeline.NewLoggerPipeline(log.LevelInfo, ljson.NewLJsonLogFormatter(), os.Stdout)
}

// NewFactory creates a standard logger pipeline factory.
//goland:noinspection GoUnusedExportedFunction
func NewFactory() log.LoggerFactory {
	return pipeline.NewLoggerPipelineFactory(ljson.NewLJsonLogFormatter(), os.Stdout)
}
