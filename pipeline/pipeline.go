package pipeline

import (
	"fmt"
	"io"
	goLog "log"

	"github.com/containerssh/log"
	"github.com/containerssh/log/formatter"
)

func NewLoggerPipeline(level log.Level, formatter formatter.Formatter, writer io.Writer) *LoggerPipeline {
	return &LoggerPipeline{
		level: level,
		formatter: formatter,
		writer: writer,
	}
}

type LoggerPipeline struct {
	level     log.Level
	formatter formatter.Formatter
	writer    io.Writer
}

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

func (pipeline *LoggerPipeline) writeE(level log.Level, err error) {
	if pipeline.level >= level {
		line := pipeline.formatter.Format(level, err.Error())
		_, err := pipeline.writer.Write(line)
		if err != nil {
			//Fallback to Go logger
			goLog.Printf("failed to write log entry (%v)\n", err)
			goLog.Println(line)
		}
	}
}

func (pipeline *LoggerPipeline) writeF(level log.Level, format string, args ...interface{}) {
	if pipeline.level >= level {
		line := pipeline.formatter.FormatData(level, fmt.Sprintf(format, args...))
		_, err := pipeline.writer.Write(line)
		if err != nil {
			//Fallback to Go logger
			goLog.Printf("failed to write log entry (%v)\n", err)
			goLog.Println(line)
		}
	}
}

func (pipeline *LoggerPipeline) writeD(level log.Level, data interface{}) {
	if pipeline.level >= level {
		line := pipeline.formatter.FormatData(level, data)
		_, err := pipeline.writer.Write(line)
		if err != nil {
			//Fallback to Go logger
			goLog.Printf("failed to write log entry (%v)\n", err)
			goLog.Println(line)
		}
	}
}
// endregion

// region Emergency
func (pipeline *LoggerPipeline) Emergency(message string) {
	pipeline.write(log.LevelEmergency, message)
}

func (pipeline *LoggerPipeline) EmergencyE(err error) {
	pipeline.writeE(log.LevelEmergency, err)
}

func (pipeline *LoggerPipeline) EmergencyD(data interface{}) {
	pipeline.writeD(log.LevelEmergency, data)
}

func (pipeline *LoggerPipeline) EmergencyF(format string, args ...interface{}) {
	pipeline.writeF(log.LevelEmergency, format, args...)
}

// endregion

// region Alert
func (pipeline *LoggerPipeline) Alert(message string) {
	pipeline.write(log.LevelAlert, message)
}

func (pipeline *LoggerPipeline) AlertE(err error) {
	pipeline.writeE(log.LevelAlert, err)
}

func (pipeline *LoggerPipeline) AlertD(data interface{}) {
	pipeline.writeD(log.LevelAlert, data)
}

func (pipeline *LoggerPipeline) AlertF(format string, args ...interface{}) {
	pipeline.writeF(log.LevelAlert, format, args...)
}

// endregion

// region Critical
func (pipeline *LoggerPipeline) Critical(message string) {
	pipeline.write(log.LevelCritical, message)
}

func (pipeline *LoggerPipeline) CriticalE(err error) {
	pipeline.writeE(log.LevelCritical, err)
}

func (pipeline *LoggerPipeline) CriticalD(data interface{}) {
	pipeline.writeD(log.LevelCritical, data)
}

func (pipeline *LoggerPipeline) CriticalF(format string, args ...interface{}) {
	pipeline.writeF(log.LevelCritical, format, args...)
}

// endregion

// region Error
func (pipeline *LoggerPipeline) Error(message string) {
	pipeline.write(log.LevelError, message)
}

func (pipeline *LoggerPipeline) ErrorE(err error) {
	pipeline.writeE(log.LevelError, err)
}

func (pipeline *LoggerPipeline) ErrorD(data interface{}) {
	pipeline.writeD(log.LevelError, data)
}

func (pipeline *LoggerPipeline) ErrorF(format string, args ...interface{}) {
	pipeline.writeF(log.LevelError, format, args...)
}

// endregion

// region Warning
func (pipeline *LoggerPipeline) Warning(message string) {
	pipeline.write(log.LevelWarning, message)
}

func (pipeline *LoggerPipeline) WarningE(err error) {
	pipeline.writeE(log.LevelWarning, err)
}

func (pipeline *LoggerPipeline) WarningD(data interface{}) {
	pipeline.writeD(log.LevelWarning, data)
}

func (pipeline *LoggerPipeline) WarningF(format string, args ...interface{}) {
	pipeline.writeF(log.LevelWarning, format, args...)
}

// endregion

// region Notice
func (pipeline *LoggerPipeline) Notice(message string) {
	pipeline.write(log.LevelNotice, message)
}

func (pipeline *LoggerPipeline) NoticeE(err error) {
	pipeline.writeE(log.LevelNotice, err)
}

func (pipeline *LoggerPipeline) NoticeD(data interface{}) {
	pipeline.writeD(log.LevelNotice, data)
}

func (pipeline *LoggerPipeline) NoticeF(format string, args ...interface{}) {
	pipeline.writeF(log.LevelNotice, format, args...)
}

// endregion

// region Info
func (pipeline *LoggerPipeline) Info(message string) {
	pipeline.write(log.LevelInfo, message)
}

func (pipeline *LoggerPipeline) InfoE(err error) {
	pipeline.writeE(log.LevelInfo, err)
}

func (pipeline *LoggerPipeline) InfoD(data interface{}) {
	pipeline.writeD(log.LevelInfo, data)
}

func (pipeline *LoggerPipeline) InfoF(format string, args ...interface{}) {
	pipeline.writeF(log.LevelInfo, format, args...)
}

// endregion

// region Debug
func (pipeline *LoggerPipeline) Debug(message string) {
	pipeline.write(log.LevelDebug, message)
}

func (pipeline *LoggerPipeline) DebugE(err error) {
	pipeline.writeE(log.LevelDebug, err)
}

func (pipeline *LoggerPipeline) DebugD(data interface{}) {
	pipeline.writeD(log.LevelDebug, data)
}

func (pipeline *LoggerPipeline) DebugF(format string, args ...interface{}) {
	pipeline.writeF(log.LevelDebug, format, args...)
}
// endregion

// region Log
func (pipeline *LoggerPipeline) Log(args ...interface{}) {
	pipeline.writeF(log.LevelInfo, "%v", args)
}
//endregion