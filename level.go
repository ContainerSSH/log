package log

import (
	"fmt"
)

// Level syslog-style log level identifiers
type Level int8

// Supported values for Level
const (
	LevelDebug     Level = 7
	LevelInfo      Level = 6
	LevelNotice    Level = 5
	LevelWarning   Level = 4
	LevelError     Level = 3
	LevelCritical  Level = 2
	LevelAlert     Level = 1
	LevelEmergency Level = 0
)

func (level Level) String() (LevelString, error) {
	switch level {
	case LevelDebug:
		return LevelDebugString, nil
	case LevelInfo:
		return LevelInfoString, nil
	case LevelNotice:
		return LevelNoticeString, nil
	case LevelWarning:
		return LevelWarningString, nil
	case LevelError:
		return LevelErrorString, nil
	case LevelCritical:
		return LevelCriticalString, nil
	case LevelAlert:
		return LevelAlertString, nil
	case LevelEmergency:
		return LevelEmergencyString, nil
	}
	return "", fmt.Errorf("invalid log level (%d)", level)
}

// Validate if the log level has a valid value
func (level Level) Validate() error {
	if level < LevelEmergency || level > LevelDebug {
		return fmt.Errorf("invalid log level (%d)", level)
	}
	return nil
}
