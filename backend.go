package gslog

import "fmt"

var theBackend Backend
var defaultLogger Logger
var defaultSugaredLogger SugaredLogger

type Backend interface {
	GetLogger(name string) Logger
	GetSugaredLogger(name string) SugaredLogger
}

func NewLogger() Logger {
	return theBackend.GetLogger("")
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

func Log(level LogLevel, msg string, args ...interface{}) {
	defaultLogger.LogDirect(level, msg, args...)
}

func Debug(msg string, args ...interface{}) {
	defaultLogger.LogDirect(LogLevelDebug, msg, args...)
}

func Info(msg string, args ...interface{}) {
	defaultLogger.LogDirect(LogLevelInfo, msg, args...)
}

func Warn(msg string, args ...interface{}) {
	defaultLogger.LogDirect(LogLevelWarn, msg, args...)
}

func Error(msg string, args ...interface{}) {
	defaultLogger.LogDirect(LogLevelError, msg, args...)
}

func Fatal(msg string, args ...interface{}) {
	defaultLogger.LogDirect(LogLevelFatal, msg, args...)
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
