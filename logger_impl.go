package log

import (
	"fmt"
	"io"
	goLog "log"
)

// NewLoggerPipeline creates a new logger pipeline with the configured minimum log level,
// a formatter to transform, and a writer to write the log to.
func NewLoggerPipeline(
	level Level,
	module string,
	formatter Formatter,
	writer io.Writer,
) Logger {
	return &loggerPipeline{
		level:     level,
		formatter: formatter,
		writer:    writer,
		module:    module,
	}
}

type loggerPipeline struct {
	level     Level
	formatter Formatter
	writer    io.Writer
	module    string
}

func (pipeline *loggerPipeline) SetLevel(level Level) {
	pipeline.level = level
}

// region Format

func (pipeline *loggerPipeline) write(level Level, message string) {
	if pipeline.level >= level {
		line := pipeline.formatter.Format(level, pipeline.module, message)
		_, err := pipeline.writer.Write(line)
		if err != nil {
			//Fallback to Go logger
			goLog.Printf("failed to write log entry (%v)\n", err)
			goLog.Println(line)
		}
	}
}

func (pipeline *loggerPipeline) writee(level Level, err error) {
	if pipeline.level >= level {
		line := pipeline.formatter.Format(level, pipeline.module, err.Error())
		_, err := pipeline.writer.Write(line)
		if err != nil {
			//Fallback to Go logger
			goLog.Printf("failed to write log entry (%v)\n", err)
			goLog.Print(line)
		}
	}
}

func (pipeline *loggerPipeline) writef(level Level, format string, args ...interface{}) {
	if pipeline.level >= level {
		line := pipeline.formatter.FormatData(
			level,
			pipeline.module,
			fmt.Sprintf(format, args...),
		)
		_, err := pipeline.writer.Write(line)
		if err != nil {
			//Fallback to Go logger
			goLog.Printf("failed to write log entry (%v)\n", err)
			goLog.Print(line)
		}
	}
}

func (pipeline *loggerPipeline) writed(level Level, data interface{}) {
	if pipeline.level >= level {
		line := pipeline.formatter.FormatData(level, pipeline.module, data)
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
func (pipeline *loggerPipeline) Emergency(message string) {
	pipeline.write(LevelEmergency, message)
}

// Emergencye writes an error on the emergency level
func (pipeline *loggerPipeline) Emergencye(err error) {
	pipeline.writee(LevelEmergency, err)
}

// Emergencyd writes a generic data interface on the emergency level
func (pipeline *loggerPipeline) Emergencyd(data interface{}) {
	pipeline.writed(LevelEmergency, data)
}

// Emergencyf writes messages in an sprintf-style format on the emergency level
func (pipeline *loggerPipeline) Emergencyf(format string, args ...interface{}) {
	pipeline.writef(LevelEmergency, format, args...)
}

// endregion

// region Alert

// Alert writes a string message on the alert level
func (pipeline *loggerPipeline) Alert(message string) {
	pipeline.write(LevelAlert, message)
}

// Alerte writes an error on the alert level
func (pipeline *loggerPipeline) Alerte(err error) {
	pipeline.writee(LevelAlert, err)
}

// Alertd writes a generic data interface on the alert level
func (pipeline *loggerPipeline) Alertd(data interface{}) {
	pipeline.writed(LevelAlert, data)
}

// Alertf writes messages in an sprintf-style format on the alert level
func (pipeline *loggerPipeline) Alertf(format string, args ...interface{}) {
	pipeline.writef(LevelAlert, format, args...)
}

// endregion

// region Critical

// Critical writes a string message on the critical level
func (pipeline *loggerPipeline) Critical(message string) {
	pipeline.write(LevelCritical, message)
}

// Criticale writes an error on the critical level
func (pipeline *loggerPipeline) Criticale(err error) {
	pipeline.writee(LevelCritical, err)
}

// Criticald writes a generic data interface on the critical level
func (pipeline *loggerPipeline) Criticald(data interface{}) {
	pipeline.writed(LevelCritical, data)
}

// Criticalf writes messages in an sprintf-style format on the critical level
func (pipeline *loggerPipeline) Criticalf(format string, args ...interface{}) {
	pipeline.writef(LevelCritical, format, args...)
}

// endregion

// region Error

// Error writes a string message on the error level
func (pipeline *loggerPipeline) Error(message string) {
	pipeline.write(LevelError, message)
}

// Errore writes an error on the error level
func (pipeline *loggerPipeline) Errore(err error) {
	pipeline.writee(LevelError, err)
}

// Errord writes a generic data interface on the error level
func (pipeline *loggerPipeline) Errord(data interface{}) {
	pipeline.writed(LevelError, data)
}

// Errorf writes messages in an sprintf-style format on the error level
func (pipeline *loggerPipeline) Errorf(format string, args ...interface{}) {
	pipeline.writef(LevelError, format, args...)
}

// endregion

// region Warning

// Warning writes a string message on the warning level
func (pipeline *loggerPipeline) Warning(message string) {
	pipeline.write(LevelWarning, message)
}

// Warninge writes an error on the warning level
func (pipeline *loggerPipeline) Warninge(err error) {
	pipeline.writee(LevelWarning, err)
}

// Warningd writes a generic data interface on the warning level
func (pipeline *loggerPipeline) Warningd(data interface{}) {
	pipeline.writed(LevelWarning, data)
}

// Warningf writes messages in an sprintf-style format on the warning level
func (pipeline *loggerPipeline) Warningf(format string, args ...interface{}) {
	pipeline.writef(LevelWarning, format, args...)
}

// endregion

// region Notice

// Notice writes a string message on the notice level
func (pipeline *loggerPipeline) Notice(message string) {
	pipeline.write(LevelNotice, message)
}

// Noticee writes an error on the notice level
func (pipeline *loggerPipeline) Noticee(err error) {
	pipeline.writee(LevelNotice, err)
}

// Noticed writes a generic data interface on the notice level
func (pipeline *loggerPipeline) Noticed(data interface{}) {
	pipeline.writed(LevelNotice, data)
}

// Noticef writes messages in an sprintf-style format on the notice level
func (pipeline *loggerPipeline) Noticef(format string, args ...interface{}) {
	pipeline.writef(LevelNotice, format, args...)
}

// endregion

// region Info

// Info writes a string message on the info level
func (pipeline *loggerPipeline) Info(message string) {
	pipeline.write(LevelInfo, message)
}

// Infoe writes an error on the info level
func (pipeline *loggerPipeline) Infoe(err error) {
	pipeline.writee(LevelInfo, err)
}

// Infod writes a generic data interface on the info level
func (pipeline *loggerPipeline) Infod(data interface{}) {
	pipeline.writed(LevelInfo, data)
}

// Infof writes messages in an sprintf-style format on the info level
func (pipeline *loggerPipeline) Infof(format string, args ...interface{}) {
	pipeline.writef(LevelInfo, format, args...)
}

// endregion

// region Debug

// Debug writes a string message on the debug level
func (pipeline *loggerPipeline) Debug(message string) {
	pipeline.write(LevelDebug, message)
}

// Debuge writes an error on the debug level
func (pipeline *loggerPipeline) Debuge(err error) {
	pipeline.writee(LevelDebug, err)
}

// Debugd writes a generic data interface on the debug level
func (pipeline *loggerPipeline) Debugd(data interface{}) {
	pipeline.writed(LevelDebug, data)
}

// Debugf writes messages in an sprintf-style format on the debug level
func (pipeline *loggerPipeline) Debugf(format string, args ...interface{}) {
	pipeline.writef(LevelDebug, format, args...)
}

// endregion

// region Log

// Log provides a generic log method that logs on the info level
func (pipeline *loggerPipeline) Log(args ...interface{}) {
	if len(args) == 1 {
		if arg, ok := args[0].(string); ok {
			pipeline.write(LevelInfo, arg)
			return
		}
	}
	pipeline.writef(LevelInfo, "%v", args)
}

//endregion
