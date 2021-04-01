package log

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

func (pipeline *logger) write(level Level, message ...interface{}) {
	if pipeline.level >= level {
		if len(message) == 0 {
			return
		}
		var msg Message
		if len(message) == 1 {
			switch message[0].(type) {
			case string:
				msg = NewMessage(EUnknownError, message[0].(string))
			default:
				if m, ok := message[0].(Message); ok {
					msg = m
				} else if m, ok := message[0].(error); ok {
					msg = pipeline.wrapError(m)
				} else {
					msg = NewMessage(EUnknownError, "%v", message[0])
				}
			}
		} else {
			msg = NewMessage(EUnknownError, "%v", message)
		}

		for label, value := range pipeline.labels {
			msg = msg.Label(label, value)
		}

		if err := pipeline.writer.Write(level, msg); err != nil {
			panic(err)
		}
	}
}

func (pipeline *logger) writef(level Level, format string, args ...interface{}) {
	if pipeline.level >= level {
		var msg Message

		msg = NewMessage(EUnknownError, format, args...)

		for label, value := range pipeline.labels {
			msg = msg.Label(label, value)
		}

		if err := pipeline.writer.Write(level, msg); err != nil {
			panic(err)
		}
	}
}

func (pipeline *logger) wrapError(err error) Message {
	return Wrap(
		err,
		EUnknownError,
		"An unexpected error has happened.",
	)
}

//endregion

//region Messages

func (pipeline *logger) Emergency(message ...interface{}) {
	pipeline.write(LevelEmergency, message...)
}

func (pipeline *logger) Alert(message ...interface{}) {
	pipeline.write(LevelAlert, message...)
}

func (pipeline *logger) Critical(message ...interface{}) {
	pipeline.write(LevelCritical, message...)
}

func (pipeline *logger) Error(message ...interface{}) {
	pipeline.write(LevelError, message...)
}

func (pipeline *logger) Warning(message ...interface{}) {
	pipeline.write(LevelWarning, message...)
}

func (pipeline *logger) Notice(message ...interface{}) {
	pipeline.write(LevelNotice, message...)
}

func (pipeline *logger) Info(message ...interface{}) {
	pipeline.write(LevelInfo, message...)
}

func (pipeline *logger) Debug(message ...interface{}) {
	pipeline.write(LevelDebug, message...)
}

//endregion

//region Log

// Log provides a generic log method that logs on the info level.
func (pipeline *logger) Log(args ...interface{}) {
	pipeline.writef(LevelInfo, "%v", args...)
}

// Logf provides a generic log method that logs on the info level with formatting.
func (pipeline *logger) Logf(format string, args ...interface{}) {
	pipeline.writef(LevelInfo, format, args...)
}

//endregion
