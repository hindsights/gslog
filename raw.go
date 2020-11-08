package gslog

import (
	"fmt"
	"log"
)

func init() {
	SetBackend(NewRawBackend(LogLevelDebug))
}

type rawBackend struct {
	logLevel LogLevel
}

func (backend *rawBackend) GetLogger(name string) Logger {
	return rawLogger{backend: backend, name: name}
}

func NewRawBackend(logLevel LogLevel) Backend {
	return &rawBackend{logLevel: logLevel}
}

type rawLogger struct {
	backend *rawBackend
	name    string
}

// func (logger rawLogger) Name() string {
// 	return logger.name
// }

func (logger rawLogger) NeedLog(level LogLevel) bool {
	return level >= logger.backend.logLevel
}

func (logger rawLogger) Logf(level LogLevel, format string, args ...interface{}) {
	if !logger.NeedLog(level) {
		return
	}
	logger.Log(level, fmt.Sprintf(format, args...))
}

func (logger rawLogger) Log(level LogLevel, args ...interface{}) {
	if !logger.NeedLog(level) {
		return
	}
	vars := append([]interface{}{fmt.Sprintf("[%5s] [%8s]", level.String(), logger.name)}, args...)
	log.Println(vars...)
}

func (logger rawLogger) Trace(args ...interface{}) {
	logger.Log(LogLevelTrace, args...)
}

func (logger rawLogger) Debug(args ...interface{}) {
	logger.Log(LogLevelDebug, args...)
}

func (logger rawLogger) Info(args ...interface{}) {
	logger.Log(LogLevelInfo, args...)
}

func (logger rawLogger) Warn(args ...interface{}) {
	logger.Log(LogLevelWarn, args...)
}

func (logger rawLogger) Error(args ...interface{}) {
	logger.Log(LogLevelError, args...)
}

func (logger rawLogger) Fatal(args ...interface{}) {
	logger.Log(LogLevelFatal, args...)
}

func (logger rawLogger) Tracef(format string, args ...interface{}) {
	logger.Logf(LogLevelTrace, format, args...)
}

func (logger rawLogger) Debugf(format string, args ...interface{}) {
	logger.Logf(LogLevelDebug, format, args...)
}

func (logger rawLogger) Infof(format string, args ...interface{}) {
	logger.Logf(LogLevelInfo, format, args...)
}

func (logger rawLogger) Warnf(format string, args ...interface{}) {
	logger.Logf(LogLevelWarn, format, args...)
}

func (logger rawLogger) Errorf(format string, args ...interface{}) {
	logger.Logf(LogLevelError, format, args...)
}

func (logger rawLogger) Fatalf(format string, args ...interface{}) {
	logger.Logf(LogLevelFatal, format, args...)
}

func (logger rawLogger) WithFields(fields Fields) Logger {
	return NewFieldsLogger(logger, fields)
}
