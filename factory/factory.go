package factory

import "github.com/containerssh/log"

// A factory to create a logger on demand
type LoggerFactory interface {
	Make(level log.Level) log.Logger
}
