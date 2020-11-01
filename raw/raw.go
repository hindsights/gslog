package raw

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hindsights/gslog"
)

type LogLevel int

const (
	LogLevelAll LogLevel = iota
	LogLevelTrace
	LogLevelDebug
	LogLevelInfo
	LogLevelWarn
	LogLevelError
	LogLevelDisable
)

var logLevelStrings []string

func init() {
	logLevelStrings = []string{
		"ALL",
		"TRACE",
		"DEBUG",
		"INFO",
		"WARN",
		"ERROR",
	}

	gslog.SetBackend(dummyBackend{})
}

type dummyBackend struct {
}

func (backend dummyBackend) GetLogger(name string) gslog.Logger {
	return dummyLogger{name: name}
}

func getLogLevel(level LogLevel) string {
	if level > LogLevelAll && level < LogLevelDisable {
		return logLevelStrings[level]
	}
	return strconv.FormatInt(int64(level), 10)
}

func dummyLog(name string, level LogLevel, args ...interface{}) {
	vars := append([]interface{}{"[" + name + "]", getLogLevel(level)}, args...)
	log.Println(vars...)
}

func dummyLogF(name string, level LogLevel, format string, args ...interface{}) {
	dummyLog(name, level, fmt.Sprintf(format, args...))
}

type dummyLogger struct {
	name string
}

func (logger dummyLogger) Tracef(format string, args ...interface{}) {
	dummyLogF(logger.name, LogLevelTrace, format, args...)
}

func (logger dummyLogger) Debugf(format string, args ...interface{}) {

}

func (logger dummyLogger) Infof(format string, args ...interface{}) {

}

func (logger dummyLogger) Warnf(format string, args ...interface{}) {

}

func (logger dummyLogger) Errorf(format string, args ...interface{}) {

}

func (logger dummyLogger) Trace(args ...interface{}) {
	dummyLog(logger.name, LogLevelTrace, args...)
}

func (logger dummyLogger) Debug(args ...interface{}) {

}

func (logger dummyLogger) Info(args ...interface{}) {

}

func (logger dummyLogger) Print(args ...interface{}) {

}

func (logger dummyLogger) Warn(args ...interface{}) {

}

func (logger dummyLogger) Error(args ...interface{}) {

}
