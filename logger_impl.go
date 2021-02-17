package log

import (
	"fmt"
)

type logger struct {
	level  Level
	labels Labels
	writer Writer
}

func (pipeline *logger) Close() error {
	return pipeline.writer.Close()
}

func (pipeline *logger) Rotate() error {
	return pipeline.writer.Rotate()
}

func (pipeline *logger) WithLevel(level Level) Logger {
	return &logger{
		level:  level,
		labels: pipeline.labels,
		writer: pipeline.writer,
	}
}

func (pipeline *logger) WithLabel(labelName LabelName, labelValue LabelValue) Logger {
	newLabels := make(Labels, len(pipeline.labels))
	for k, v := range pipeline.labels {
		newLabels[k] = v
	}
	newLabels[labelName] = labelValue
	return &logger{
		level:  pipeline.level,
		labels: newLabels,
		writer: pipeline.writer,
	}
}

//region Format

func (pipeline *logger) write(level Level, err error) {
	if pipeline.level >= level {
		var msg Message
		var ok bool
		if msg, ok = err.(Message); !ok {
			msg = pipeline.wrapError(err)
		}

		for label, value := range pipeline.labels {
			msg = msg.Label(label, value)
		}

		if err := pipeline.writer.Write(level, msg); err != nil {
			panic(err)
		}
	}
}

func (pipeline *logger) wrapError(err error) Message {
	return WrapError(
		err,
		EUnknownError,
		"An unexpected error has happened.",
	)
}

//endregion

//region Messages

// Emergency writes a string message on the emergency level
func (pipeline *logger) Emergency(err error) {
	pipeline.write(LevelEmergency, err)
}

// Alert writes a string message on the alert level
func (pipeline *logger) Alert(err error) {
	pipeline.write(LevelAlert, err)
}

// Critical writes a string message on the critical level
func (pipeline *logger) Critical(err error) {
	pipeline.write(LevelCritical, err)
}

// Error writes a string message on the error level
func (pipeline *logger) Error(err error) {
	pipeline.write(LevelError, err)
}

// Warning writes a string message on the warning level
func (pipeline *logger) Warning(err error) {
	pipeline.write(LevelWarning, err)
}

// Notice writes a string message on the notice level
func (pipeline *logger) Notice(err error) {
	pipeline.write(LevelNotice, err)
}

// Info writes a string message on the info level
func (pipeline *logger) Info(err error) {
	pipeline.write(LevelInfo, err)
}

// Debug writes a string message on the debug level
func (pipeline *logger) Debug(err error) {
	pipeline.write(LevelDebug, err)
}

//endregion

//region Log

// Log provides a generic log method that logs on the info level.
func (pipeline *logger) Log(args ...interface{}) {
	pipeline.write(LevelInfo, fmt.Errorf("%v", args...))
}

// Logf provides a generic log method that logs on the info level with formatting.
func (pipeline *logger) Logf(format string, args ...interface{}) {
	pipeline.write(LevelInfo, fmt.Errorf(format, args...))
}

//endregion
