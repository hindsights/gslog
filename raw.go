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
	return rawSLogger{backend: backend, name: name}
}

func NewRawBackend(logLevel LogLevel) Backend {
	return &rawBackend{logLevel: logLevel}
}

type rawSLogger struct {
	backend *rawBackend
	name    string
	fields  Fields
}

func (logger rawSLogger) NeedLog(level LogLevel) bool {
	return level >= logger.backend.logLevel
}

func (logger rawSLogger) WithFields(fields Fields) FieldLogger {
	return rawSLogger{backend: logger.backend, name: logger.name, fields: JoinFields(logger.fields, fields)}
}

func (logger rawSLogger) Log(level LogLevel, msg string, fields ...Fields) {
	if !logger.NeedLog(level) {
		return
	}
	fieldCount := 0
	if len(fields) > 0 {
		fieldCount += len(fields[0])
	}
	fieldArgs := FormatFields(append([]Fields{logger.fields}, fields...)...)
	vars := make([]interface{}, 3+len(fieldArgs))
	vars[0] = fmt.Sprintf("[%5s]", level.String())
	vars[1] = fmt.Sprintf("[%8s]", logger.name)
	vars[2] = msg
	for i, arg := range fieldArgs {
		vars[i+3] = arg
	}
	log.Println(vars...)
}

func (logger rawSLogger) Debug(msg string, fields ...Fields) {
	logger.Log(LogLevelDebug, msg, fields...)
}

func (logger rawSLogger) Info(msg string, fields ...Fields) {
	logger.Log(LogLevelInfo, msg, fields...)
}

func (logger rawSLogger) Warn(msg string, fields ...Fields) {
	logger.Log(LogLevelWarn, msg, fields...)
}

func (logger rawSLogger) Error(msg string, fields ...Fields) {
	logger.Log(LogLevelError, msg, fields...)
}

func (logger rawSLogger) Fatal(msg string, fields ...Fields) {
	logger.Log(LogLevelFatal, msg, fields...)
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
