package log

import "fmt"

// swagger:enum LevelString
type LevelString string
const (
	LevelDebugString     LevelString = "debug"
	LevelInfoString      LevelString = "info"
	LevelNoticeString    LevelString = "notice"
	LevelWarningString   LevelString = "warning"
	LevelErrorString     LevelString = "error"
	LevelCriticalString  LevelString = "crit"
	LevelAlertString     LevelString = "alert"
	LevelEmergencyString LevelString = "emerg"
)

func (level LevelString) ToLevel() (Level, error) {
	switch level {
	case LevelDebugString:
		return LevelDebug, nil
	case LevelInfoString:
		return LevelInfo, nil
	case LevelNoticeString:
		return LevelNotice, nil
	case LevelWarningString:
		return LevelWarning, nil
	case LevelErrorString:
		return LevelError, nil
	case LevelCriticalString:
		return LevelCritical, nil
	case LevelAlertString:
		return LevelAlert, nil
	case LevelEmergencyString:
		return LevelEmergency, nil
	}
	return -1, fmt.Errorf("invalid log level (%s)", level)
}
