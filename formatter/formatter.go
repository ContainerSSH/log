package formatter

import "github.com/containerssh/log"

type Formatter interface {
	Format(level log.Level, message string) []byte
	FormatData(level log.Level, data interface{}) []byte
}
