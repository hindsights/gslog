package gslog

import (
	"fmt"
	"log"
	"os"
)

func init() {
	SetBackend(NewRawBackend(LogLevelDebug))
}

type rawBackend struct {
	logger   *log.Logger
	logLevel LogLevel
}

func (backend *rawBackend) GetLogger(name string) Logger {
	return NewAttrLogger(backend.GetSugaredLogger(name))
}

func (backend *rawBackend) GetSugaredLogger(name string) SugaredLogger {
	return rawLogger{backend: backend, name: name}
}

func NewRawBackend(logLevel LogLevel) Backend {
	logger := log.New(os.Stderr, "", log.Ldate|log.Lmicroseconds)
	return &rawBackend{logLevel: logLevel, logger: logger}
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
	logger.LogfDirect(level, format, args...)
}

func (logger rawLogger) LogfDirect(level LogLevel, format string, args ...interface{}) {
	if !logger.NeedLog(level) {
		return
	}
	newFormat, vars := logger.prepareArgs(level, format, args...)
	logger.backend.logger.Printf(newFormat+"\n", vars...)
}

func (logger rawLogger) Log(level LogLevel, args ...interface{}) {
	logger.LogDirect(level, args...)
}

func (logger rawLogger) LogDirect(level LogLevel, args ...interface{}) {
	if !logger.NeedLog(level) {
		return
	}
	_, vars := logger.prepareArgs(level, "", args...)
	logger.backend.logger.Println(vars...)
}

func (logger rawLogger) Debug(args ...interface{}) {
	logger.LogDirect(LogLevelDebug, args...)
}

func (logger rawLogger) Info(args ...interface{}) {
	logger.LogDirect(LogLevelInfo, args...)
}

func (logger rawLogger) Warn(args ...interface{}) {
	logger.LogDirect(LogLevelWarn, args...)
}

func (logger rawLogger) Error(args ...interface{}) {
	logger.LogDirect(LogLevelError, args...)
}

func (logger rawLogger) Fatal(args ...interface{}) {
	logger.LogDirect(LogLevelFatal, args...)
	panic(args)
}

func (logger rawLogger) Debugf(format string, args ...interface{}) {
	logger.LogfDirect(LogLevelDebug, format, args...)
}

func (logger rawLogger) Infof(format string, args ...interface{}) {
	logger.LogfDirect(LogLevelInfo, format, args...)
}

func (logger rawLogger) Warnf(format string, args ...interface{}) {
	logger.LogfDirect(LogLevelWarn, format, args...)
}

func (logger rawLogger) Errorf(format string, args ...interface{}) {
	logger.LogfDirect(LogLevelError, format, args...)
}

func (logger rawLogger) Fatalf(format string, args ...interface{}) {
	logger.LogfDirect(LogLevelFatal, format, args...)
	panic(fmt.Sprintf(format, args...))
}
