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

func (backend *rawBackend) GetFieldLogger(name string) FieldLogger {
	return NewFieldLogger(backend.GetLogger(name))
}

func NewRawBackend(logLevel LogLevel) Backend {
	return &rawBackend{logLevel: logLevel}
}

type rawLogger struct {
	backend *rawBackend
	name    string
	fields  Fields
}

func (logger rawLogger) NeedLog(level LogLevel) bool {
	return level >= logger.backend.logLevel
}

func (logger rawLogger) prepareArgs(level LogLevel, format string, args ...interface{}) (string, []interface{}) {
	fieldsArgs := FormatFields(logger.fields)
	vars := make([]interface{}, 2, 2+len(args)+len(fieldsArgs))
	vars[0] = fmt.Sprintf("[%5s]", level.String())
	vars[1] = fmt.Sprintf("[%8s]", logger.name)
	vars = append(vars, args...)
	newFormat := "%s %s " + format
	for _, arg := range fieldsArgs {
		vars = append(vars, arg)
		newFormat += " %s"
	}
	return newFormat, vars
}

func (logger rawLogger) Logf(level LogLevel, format string, args ...interface{}) {
	if !logger.NeedLog(level) {
		return
	}
	newFormat, vars := logger.prepareArgs(level, format, args...)
	log.Printf(newFormat+"\n", vars...)
}

func (logger rawLogger) Log(level LogLevel, args ...interface{}) {
	if !logger.NeedLog(level) {
		return
	}
	_, vars := logger.prepareArgs(level, "", args...)
	log.Println(vars...)
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
