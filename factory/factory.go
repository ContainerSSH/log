package factory

import "github.com/containerssh/log"

type LoggerFactory interface{
	Make(level log.Level) log.Logger
}
