package factory

import "github.com/containerssh/log"

// LoggerFactory is a factory to create a logger on demand
type LoggerFactory interface {
	Make(level log.Level) log.Logger
}
