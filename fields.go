package gslog

import (
	"fmt"
)

type Fields map[string]interface{}

func CopyFields(fields Fields) Fields {
	if fields == nil || len(fields) == 0 {
		return nil
	}
	return JoinFields(fields)
}

func JoinFields(fields ...Fields) Fields {
	if len(fields) == 0 {
		return nil
	}
	ret := make(Fields)
	for _, fs := range fields {
		for k, v := range fs {
			ret[k] = v
		}
	}
	return ret
}

func FormatFields(fields ...Fields) []string {
	if len(fields) == 0 {
		return nil
	}
	ret := make([]string, 0, GetFieldCount(fields...))
	for _, fs := range fields {
		for k, v := range fs {
			ret = append(ret, fmt.Sprintf("%s=%v", k, v))
		}
	}
	return ret
}

func GetFieldCount(fields ...Fields) int {
	count := 0
	for _, fs := range fields {
		count += len(fs)
	}
	return count
}

type rawSLogger struct {
	logger Logger
	fields Fields
}

func NewFieldLogger(logger Logger) FieldLogger {
	return rawSLogger{logger: logger}
}

func (logger rawSLogger) NeedLog(level LogLevel) bool {
	return logger.logger.NeedLog(level)
}

func (logger rawSLogger) WithFields(fields Fields) FieldLogger {
	return rawSLogger{logger: logger.logger, fields: JoinFields(logger.fields, fields)}
}

func (logger rawSLogger) Log(level LogLevel, msg string, fields ...Fields) {
	if !logger.NeedLog(level) {
		return
	}
	fieldArgs := FormatFields(append([]Fields{logger.fields}, fields...)...)
	args := make([]interface{}, 1+len(fieldArgs))
	args[0] = msg
	for i, arg := range fieldArgs {
		args[i+1] = arg
	}
	logger.logger.Log(level, args...)
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
