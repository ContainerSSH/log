package log

// The logger interface provides logging facilities on various levesl
type Logger interface {
	// Change the minimum log level
	SetLevel(level Level)
	// Log a string message on debug level
	Debug(message string)
	// Log an error on debug level
	Debuge(err error)
	// Log a generic data on debug level
	Debugd(data interface{})
	// Log an Sprintf-style message with args on debug level
	Debugf(format string, args ...interface{})
	Info(message string)
	Infoe(err error)
	Infod(data interface{})
	Infof(format string, args ...interface{})
	Notice(message string)
	Noticee(err error)
	Noticed(data interface{})
	Noticef(format string, args ...interface{})
	Warning(message string)
	Warninge(err error)
	Warningd(data interface{})
	Warningf(format string, args ...interface{})
	Error(message string)
	Errore(err error)
	Errord(data interface{})
	Errorf(format string, args ...interface{})
	Critical(message string)
	Criticale(err error)
	Criticald(data interface{})
	Criticalf(format string, args ...interface{})
	Alert(message string)
	Alerte(err error)
	Alertd(data interface{})
	Alertf(format string, args ...interface{})
	Emergency(message string)
	Emergencye(err error)
	Emergencyd(data interface{})
	Emergencyf(format string, args ...interface{})
	Log(...interface{})
}
