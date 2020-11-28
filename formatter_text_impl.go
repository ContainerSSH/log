package log

import (
	"fmt"
)

type textLogFormatter struct {
}

func (t *textLogFormatter) getLevel(level Level) LevelString {
	levelString, err := level.String()
	if err != nil {
		levelString = LevelCriticalString
	}
	return levelString
}

func (t *textLogFormatter) Format(level Level, module string, message string) []byte {
	return []byte(fmt.Sprintf("[%s]\t[%s]\t%s\n", t.getLevel(level), module, message))
}

func (t *textLogFormatter) FormatData(level Level, module string, data interface{}) []byte {
	return []byte(fmt.Sprintf("[%s]\t[%s]\t%v\n", t.getLevel(level), module, data))
}
