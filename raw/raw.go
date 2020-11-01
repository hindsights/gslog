package raw

import (
	"fmt"
	"log"

	"github.com/hindsights/gslog"
)

func init() {
	gslog.SetBackend(NewRawBackend(gslog.LogLevelDebug))
}

type rawBackend struct {
	logLevel gslog.LogLevel
}

func (backend *rawBackend) GetLogger(name string) gslog.Logger {
	return rawLogger{backend: backend, name: name}
}

func NewRawBackend(logLevel gslog.LogLevel) gslog.Backend {
	return &rawBackend{logLevel: logLevel}
}

type rawLogger struct {
	backend *rawBackend
	name    string
}

// func (logger rawLogger) Name() string {
// 	return logger.name
// }

func (logger rawLogger) NeedLog(level gslog.LogLevel) bool {
	return level >= logger.backend.logLevel
}

func (logger rawLogger) Logf(level gslog.LogLevel, format string, args ...interface{}) {
	if !logger.NeedLog(level) {
		return
	}
	logger.Log(level, fmt.Sprintf(format, args...))
}

func (logger rawLogger) Log(level gslog.LogLevel, args ...interface{}) {
	if !logger.NeedLog(level) {
		return
	}
	vars := append([]interface{}{fmt.Sprintf("[%5s] [%8s]", level.String(), logger.name)}, args...)
	log.Println(vars...)
}

func (logger rawLogger) Trace(args ...interface{}) {
	logger.Log(gslog.LogLevelTrace, args...)
}

func (logger rawLogger) Debug(args ...interface{}) {
	logger.Log(gslog.LogLevelDebug, args...)
}

func (logger rawLogger) Info(args ...interface{}) {
	logger.Log(gslog.LogLevelInfo, args...)
}

func (logger rawLogger) Warn(args ...interface{}) {
	logger.Log(gslog.LogLevelWarn, args...)
}

func (logger rawLogger) Error(args ...interface{}) {
	logger.Log(gslog.LogLevelError, args...)
}

func (logger rawLogger) Fatal(args ...interface{}) {
	logger.Log(gslog.LogLevelFatal, args...)
}

func (logger rawLogger) Tracef(format string, args ...interface{}) {
	logger.Logf(gslog.LogLevelTrace, format, args...)
}

func (logger rawLogger) Debugf(format string, args ...interface{}) {
	logger.Logf(gslog.LogLevelDebug, format, args...)
}

func (logger rawLogger) Infof(format string, args ...interface{}) {
	logger.Logf(gslog.LogLevelInfo, format, args...)
}

func (logger rawLogger) Warnf(format string, args ...interface{}) {
	logger.Logf(gslog.LogLevelWarn, format, args...)
}

func (logger rawLogger) Errorf(format string, args ...interface{}) {
	logger.Logf(gslog.LogLevelError, format, args...)
}

func (logger rawLogger) Fatalf(format string, args ...interface{}) {
	logger.Logf(gslog.LogLevelFatal, format, args...)
}

func (logger rawLogger) WithFields(fields gslog.Fields) gslog.Logger {
	return gslog.NewFieldsLogger(logger, fields)
}
