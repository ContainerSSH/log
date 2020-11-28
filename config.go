package log

import (
	"fmt"
)

// Config describes the logging settings.
type Config struct {
	// Level describes the minimum level to log at
	Level Level `json:"level" yaml:"level" default:"notice"`
	// Format describes the log message format
	Format Format `json:"format" yaml:"format" default:"ljson"`
}

// region Level

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

// String Convert the int level to the string representation
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

// endregion

// region LevelString

// LevelString is a type for supported log level strings
type LevelString string

// List of valid string values for log levels
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

// ToLevel convert the string level to the int representation
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

// endregion

// region Format

// Format is the logging format to use.
//swagger:enum
type Format string

const (
	// FormatLJSON is a newline-delimited JSON log format.
	FormatLJSON Format = "ljson"
	// FormatText prints the logs as plain text.
	FormatText Format = "text"
)

// Validate returns an error if the format is invalid.
func (format Format) Validate() error {
	switch format {
	case FormatLJSON:
	case FormatText:
	default:
		return fmt.Errorf("invalid log format: %s", format)
	}
	return nil
}

// endregion
