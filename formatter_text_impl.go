package log

import (
	"fmt"
	"time"
)

type textLogFormatter struct {
}

func (t *textLogFormatter) getTimestamp() string {
	return time.Now().Format(time.RFC3339)
}

func (t *textLogFormatter) getLevel(level Level) LevelString {
	levelString, err := level.String()
	if err != nil {
		levelString = LevelCriticalString
	}
	return levelString
}

func (t *textLogFormatter) Format(level Level, module string, message string) []byte {
	return []byte(fmt.Sprintf("%s\t%s\t%s\t%s\n", t.getTimestamp(), t.getLevel(level), module, message))
}

func (t *textLogFormatter) FormatData(level Level, module string, data interface{}) []byte {
	return []byte(fmt.Sprintf("%s\t%s\t%s\t%v\n", t.getTimestamp(), t.getLevel(level), module, data))
}
