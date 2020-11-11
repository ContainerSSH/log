package pipeline

import (
	"fmt"
	"io"
	goLog "log"

	"github.com/containerssh/log"
	"github.com/containerssh/log/formatter"
)

// NewLoggerPipeline creates a new logger pipeline with the configured minimum log level,
// a formatter to transform, and a writer to write the log to.
func NewLoggerPipeline(level log.Level, formatter formatter.Formatter, writer io.Writer) *LoggerPipeline {
	return &LoggerPipeline{
		level:     level,
		formatter: formatter,
		writer:    writer,
	}
}

// LoggerPipeline is a regular pipeline that transforms and writes logs to a regular io.Writer
type LoggerPipeline struct {
	level     log.Level
	formatter formatter.Formatter
	writer    io.Writer
}

// SetLevel sets the minimum log level
func (pipeline *LoggerPipeline) SetLevel(level log.Level) {
	pipeline.level = level
}

// region Format

func (pipeline *LoggerPipeline) write(level log.Level, message string) {
	if pipeline.level >= level {
		line := pipeline.formatter.Format(level, message)
		_, err := pipeline.writer.Write(line)
		if err != nil {
			//Fallback to Go logger
			goLog.Printf("failed to write log entry (%v)\n", err)
			goLog.Println(line)
		}
	}
}

func (pipeline *LoggerPipeline) writee(level log.Level, err error) {
	if pipeline.level >= level {
		line := pipeline.formatter.Format(level, err.Error())
		_, err := pipeline.writer.Write(line)
		if err != nil {
			//Fallback to Go logger
			goLog.Printf("failed to write log entry (%v)\n", err)
			goLog.Print(line)
		}
	}
}

func (pipeline *LoggerPipeline) writef(level log.Level, format string, args ...interface{}) {
	if pipeline.level >= level {
		line := pipeline.formatter.FormatData(level, fmt.Sprintf(format, args...))
		_, err := pipeline.writer.Write(line)
		if err != nil {
			//Fallback to Go logger
			goLog.Printf("failed to write log entry (%v)\n", err)
			goLog.Print(line)
		}
	}
}

func (pipeline *LoggerPipeline) writed(level log.Level, data interface{}) {
	if pipeline.level >= level {
		line := pipeline.formatter.FormatData(level, data)
		_, err := pipeline.writer.Write(line)
		if err != nil {
			//Fallback to Go logger
			goLog.Printf("failed to write log entry (%v)\n", err)
			goLog.Print(line)
		}
	}
}

// endregion

// region Emergency

// Emergency writes a string message on the emergency level
func (pipeline *LoggerPipeline) Emergency(message string) {
	pipeline.write(log.LevelEmergency, message)
}

// Emergencye writes an error on the emergency level
func (pipeline *LoggerPipeline) Emergencye(err error) {
	pipeline.writee(log.LevelEmergency, err)
}

// Emergencyd writes a generic data interface on the emergency level
func (pipeline *LoggerPipeline) Emergencyd(data interface{}) {
	pipeline.writed(log.LevelEmergency, data)
}

// Emergencyf writes messages in an sprintf-style format on the emergency level
func (pipeline *LoggerPipeline) Emergencyf(format string, args ...interface{}) {
	pipeline.writef(log.LevelEmergency, format, args...)
}

// endregion

// region Alert

// Alert writes a string message on the alert level
func (pipeline *LoggerPipeline) Alert(message string) {
	pipeline.write(log.LevelAlert, message)
}

// Alerte writes an error on the alert level
func (pipeline *LoggerPipeline) Alerte(err error) {
	pipeline.writee(log.LevelAlert, err)
}

// Alertd writes a generic data interface on the alert level
func (pipeline *LoggerPipeline) Alertd(data interface{}) {
	pipeline.writed(log.LevelAlert, data)
}

// Alertf writes messages in an sprintf-style format on the alert level
func (pipeline *LoggerPipeline) Alertf(format string, args ...interface{}) {
	pipeline.writef(log.LevelAlert, format, args...)
}

// endregion

// region Critical

// Critical writes a string message on the critical level
func (pipeline *LoggerPipeline) Critical(message string) {
	pipeline.write(log.LevelCritical, message)
}

// Criticale writes an error on the critical level
func (pipeline *LoggerPipeline) Criticale(err error) {
	pipeline.writee(log.LevelCritical, err)
}

// Criticald writes a generic data interface on the critical level
func (pipeline *LoggerPipeline) Criticald(data interface{}) {
	pipeline.writed(log.LevelCritical, data)
}

// Criticalf writes messages in an sprintf-style format on the critical level
func (pipeline *LoggerPipeline) Criticalf(format string, args ...interface{}) {
	pipeline.writef(log.LevelCritical, format, args...)
}

// endregion

// region Error

// Error writes a string message on the error level
func (pipeline *LoggerPipeline) Error(message string) {
	pipeline.write(log.LevelError, message)
}

// Errore writes an error on the error level
func (pipeline *LoggerPipeline) Errore(err error) {
	pipeline.writee(log.LevelError, err)
}

// Errord writes a generic data interface on the error level
func (pipeline *LoggerPipeline) Errord(data interface{}) {
	pipeline.writed(log.LevelError, data)
}

// Errorf writes messages in an sprintf-style format on the error level
func (pipeline *LoggerPipeline) Errorf(format string, args ...interface{}) {
	pipeline.writef(log.LevelError, format, args...)
}

// endregion

// region Warning

// Warning writes a string message on the warning level
func (pipeline *LoggerPipeline) Warning(message string) {
	pipeline.write(log.LevelWarning, message)
}

// Warninge writes an error on the warning level
func (pipeline *LoggerPipeline) Warninge(err error) {
	pipeline.writee(log.LevelWarning, err)
}

// Warningd writes a generic data interface on the warning level
func (pipeline *LoggerPipeline) Warningd(data interface{}) {
	pipeline.writed(log.LevelWarning, data)
}

// Warningf writes messages in an sprintf-style format on the warning level
func (pipeline *LoggerPipeline) Warningf(format string, args ...interface{}) {
	pipeline.writef(log.LevelWarning, format, args...)
}

// endregion

// region Notice

// Notice writes a string message on the notice level
func (pipeline *LoggerPipeline) Notice(message string) {
	pipeline.write(log.LevelNotice, message)
}

// Noticee writes an error on the notice level
func (pipeline *LoggerPipeline) Noticee(err error) {
	pipeline.writee(log.LevelNotice, err)
}

// Noticed writes a generic data interface on the notice level
func (pipeline *LoggerPipeline) Noticed(data interface{}) {
	pipeline.writed(log.LevelNotice, data)
}

// Noticef writes messages in an sprintf-style format on the notice level
func (pipeline *LoggerPipeline) Noticef(format string, args ...interface{}) {
	pipeline.writef(log.LevelNotice, format, args...)
}

// endregion

// region Info

// Info writes a string message on the info level
func (pipeline *LoggerPipeline) Info(message string) {
	pipeline.write(log.LevelInfo, message)
}

// Infoe writes an error on the info level
func (pipeline *LoggerPipeline) Infoe(err error) {
	pipeline.writee(log.LevelInfo, err)
}

// Infod writes a generic data interface on the info level
func (pipeline *LoggerPipeline) Infod(data interface{}) {
	pipeline.writed(log.LevelInfo, data)
}

// Infof writes messages in an sprintf-style format on the info level
func (pipeline *LoggerPipeline) Infof(format string, args ...interface{}) {
	pipeline.writef(log.LevelInfo, format, args...)
}

// endregion

// region Debug

// Debug writes a string message on the debug level
func (pipeline *LoggerPipeline) Debug(message string) {
	pipeline.write(log.LevelDebug, message)
}

// Debuge writes an error on the debug level
func (pipeline *LoggerPipeline) Debuge(err error) {
	pipeline.writee(log.LevelDebug, err)
}

// Debugd writes a generic data interface on the debug level
func (pipeline *LoggerPipeline) Debugd(data interface{}) {
	pipeline.writed(log.LevelDebug, data)
}

// Debugf writes messages in an sprintf-style format on the debug level
func (pipeline *LoggerPipeline) Debugf(format string, args ...interface{}) {
	pipeline.writef(log.LevelDebug, format, args...)
}

// endregion

// region Log

// Log provides a generic log method that logs on the info level
func (pipeline *LoggerPipeline) Log(args ...interface{}) {
	if len(args) == 1 {
		if arg, ok := args[0].(string); ok {
			pipeline.write(log.LevelInfo, arg)
			return
		}
	}
	pipeline.writef(log.LevelInfo, "%v", args)
}

//endregion
