package gslog

import "fmt"

var theBackend Backend
var defaultLogger Logger
var defaultSugaredLogger SugaredLogger

type Backend interface {
	GetLogger(name string) Logger
	GetSugaredLogger(name string) SugaredLogger
}

func GetLogger(name string) Logger {
	return theBackend.GetLogger(name)
}

func GetSugaredLogger(name string) SugaredLogger {
	return theBackend.GetSugaredLogger(name)
}

func SetBackend(backend Backend) {
	theBackend = backend
	defaultLogger = backend.GetLogger("log")
	defaultSugaredLogger = backend.GetSugaredLogger("log")
}

func Logf(level LogLevel, format string, args ...interface{}) {
	defaultSugaredLogger.LogfDirect(level, format, args...)
}

func Log(level LogLevel, args ...interface{}) {
	defaultSugaredLogger.LogDirect(level, args...)
}

func Debug(args ...interface{}) {
	defaultSugaredLogger.LogDirect(LogLevelDebug, args...)
}

func Info(args ...interface{}) {
	defaultSugaredLogger.LogDirect(LogLevelInfo, args...)
}

func Warn(args ...interface{}) {
	defaultSugaredLogger.LogDirect(LogLevelWarn, args...)
}

func Error(args ...interface{}) {
	defaultSugaredLogger.LogDirect(LogLevelError, args...)
}

func Fatal(args ...interface{}) {
	defaultSugaredLogger.LogDirect(LogLevelFatal, args...)
	panic(args)
}

func Debugf(format string, args ...interface{}) {
	defaultSugaredLogger.LogfDirect(LogLevelDebug, format, args...)
}

func Infof(format string, args ...interface{}) {
	defaultSugaredLogger.LogfDirect(LogLevelInfo, format, args...)
}

func Warnf(format string, args ...interface{}) {
	defaultSugaredLogger.LogfDirect(LogLevelWarn, format, args...)
}

func Errorf(format string, args ...interface{}) {
	defaultSugaredLogger.LogfDirect(LogLevelError, format, args...)
}

func Fatalf(format string, args ...interface{}) {
	defaultSugaredLogger.LogfDirect(LogLevelFatal, format, args...)
	panic(fmt.Sprintf(format, args...))
}
