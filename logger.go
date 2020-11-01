package gslog

type Fields map[string]interface{}

type BaseLogger interface {
	NeedLog(level LogLevel) bool
	Logf(level LogLevel, format string, args ...interface{})
	Log(level LogLevel, args ...interface{})

	Trace(args ...interface{})
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})

	Tracef(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
}

type Logger interface {
	BaseLogger
	WithFields(fields Fields) Logger
}

// func WithFields(logger Logger, fields Fields) Logger {
// 	return logger.WithFields(fields)
// }
