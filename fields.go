package gslog

import (
	"fmt"
)

func transformFields(fields Fields) []interface{} {
	if len(fields) == 0 {
		return nil
	}
	strs := make([]interface{}, 1, 1+len(fields))
	strs[0] = "\t"
	for k, v := range fields {
		strs = append(strs, fmt.Sprintf("%s=%v", k, v))
	}
	return strs
}

type fieldsLogger struct {
	baseLogger BaseLogger
	fields     Fields
}

func NewFieldsLogger(baseLogger BaseLogger, fields Fields) Logger {
	return &fieldsLogger{baseLogger: baseLogger, fields: fields}
}

func (logger *fieldsLogger) NeedLog(level LogLevel) bool {
	return logger.baseLogger.NeedLog(level)
}

func (logger *fieldsLogger) Logf(level LogLevel, format string, args ...interface{}) {
	if !logger.NeedLog(level) {
		return
	}
	msg := fmt.Sprintf(format, args...)
	logger.Log(level, msg)
}

func (logger *fieldsLogger) Log(level LogLevel, args ...interface{}) {
	if !logger.NeedLog(level) {
		return
	}
	fields := transformFields(logger.fields)
	newArgs := append(args, fields...)
	logger.baseLogger.Log(level, newArgs...)
}

func (logger *fieldsLogger) Trace(args ...interface{}) {
	logger.Log(LogLevelTrace, args...)
}

func (logger *fieldsLogger) Debug(args ...interface{}) {
	logger.Log(LogLevelDebug, args...)
}

func (logger *fieldsLogger) Info(args ...interface{}) {
	logger.Log(LogLevelInfo, args...)
}

func (logger *fieldsLogger) Warn(args ...interface{}) {
	logger.Log(LogLevelWarn, args...)
}

func (logger *fieldsLogger) Error(args ...interface{}) {
	logger.Log(LogLevelError, args...)
}

func (logger *fieldsLogger) Fatal(args ...interface{}) {
	logger.Log(LogLevelFatal, args...)
}

func (logger *fieldsLogger) Tracef(format string, args ...interface{}) {
	logger.Logf(LogLevelTrace, format, args...)
}

func (logger *fieldsLogger) Debugf(format string, args ...interface{}) {
	logger.Logf(LogLevelDebug, format, args...)
}

func (logger *fieldsLogger) Infof(format string, args ...interface{}) {
	logger.Logf(LogLevelInfo, format, args...)
}

func (logger *fieldsLogger) Warnf(format string, args ...interface{}) {
	logger.Logf(LogLevelWarn, format, args...)
}

func (logger *fieldsLogger) Errorf(format string, args ...interface{}) {
	logger.Logf(LogLevelError, format, args...)
}

func (logger *fieldsLogger) Fatalf(format string, args ...interface{}) {
	logger.Logf(LogLevelFatal, format, args...)
}

func (logger *fieldsLogger) WithFields(fields Fields) Logger {
	return NewFieldsLogger(logger, fields)
}
