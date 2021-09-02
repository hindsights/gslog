package gslog

var theBackend Backend
var defaultLogger Logger

type Backend interface {
	GetLogger(name string) Logger
	GetFieldLogger(name string) FieldLogger
}

func GetFieldLogger(name string) FieldLogger {
	return theBackend.GetFieldLogger(name)
}

func GetLogger(name string) Logger {
	return theBackend.GetLogger(name)
}

func SetBackend(backend Backend) {
	theBackend = backend
	defaultLogger = backend.GetLogger("log")
}

func Logf(level LogLevel, format string, args ...interface{}) {
	defaultLogger.Logf(level, format, args...)
}

func Log(level LogLevel, args ...interface{}) {
	defaultLogger.Log(level, args...)
}

func Debug(args ...interface{}) {
	defaultLogger.Log(LogLevelDebug, args...)
}

func Info(args ...interface{}) {
	defaultLogger.Log(LogLevelInfo, args...)
}

func Warn(args ...interface{}) {
	defaultLogger.Log(LogLevelWarn, args...)
}

func Error(args ...interface{}) {
	defaultLogger.Log(LogLevelError, args...)
}

func Fatal(args ...interface{}) {
	defaultLogger.Log(LogLevelFatal, args...)
}

func Debugf(format string, args ...interface{}) {
	defaultLogger.Logf(LogLevelDebug, format, args...)
}

func Infof(format string, args ...interface{}) {
	defaultLogger.Logf(LogLevelInfo, format, args...)
}

func Warnf(format string, args ...interface{}) {
	defaultLogger.Logf(LogLevelWarn, format, args...)
}

func Errorf(format string, args ...interface{}) {
	defaultLogger.Logf(LogLevelError, format, args...)
}

func Fatalf(format string, args ...interface{}) {
	defaultLogger.Logf(LogLevelFatal, format, args...)
}
