package gslog

var theBackend Backend
var defaultLogger SimpleLogger

type Backend interface {
	GetLogger(name string) Logger
	GetSimpleLogger(name string) SimpleLogger
}

func GetLogger(name string) Logger {
	return theBackend.GetLogger(name)
}

func GetSimpleLogger(name string) SimpleLogger {
	return theBackend.GetSimpleLogger(name)
}

func SetBackend(backend Backend) {
	theBackend = backend
	defaultLogger = backend.GetSimpleLogger("log")
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
