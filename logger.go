package gslog

type Fields map[string]interface{}

type Logger interface {
	// WithFields(fields Fields) Logger

	Tracef(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})

	Trace(args ...interface{})
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
}

// func WithFields(logger Logger, fields Fields) Logger {
// 	return logger.WithFields(fields)
// }
