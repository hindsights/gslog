package gslog

import (
	"fmt"
	"time"
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

type rawFieldLogger struct {
	logger SugaredLogger
	fields Fields
}

func NewFieldLogger(logger SugaredLogger) Logger {
	return rawFieldLogger{logger: logger}
}

func (logger rawFieldLogger) NeedLog(level LogLevel) bool {
	return logger.logger.NeedLog(level)
}

func (logger rawFieldLogger) Log(level LogLevel, msg string) {
	if !logger.NeedLog(level) {
		return
	}
	fieldArgs := FormatFields(logger.fields)
	args := make([]interface{}, 2+len(fieldArgs))
	args[0] = msg
	args[1] = "\t"
	for i, arg := range fieldArgs {
		args[i+2] = arg
	}
	logger.logger.Log(level, args...)
}

func (logger rawFieldLogger) Debug(msg string) {
	logger.Log(LogLevelDebug, msg)
}

func (logger rawFieldLogger) Info(msg string) {
	logger.Log(LogLevelInfo, msg)
}

func (logger rawFieldLogger) Warn(msg string) {
	logger.Log(LogLevelWarn, msg)
}

func (logger rawFieldLogger) Error(msg string) {
	logger.Log(LogLevelError, msg)
}

func (logger rawFieldLogger) Fatal(msg string) {
	logger.Log(LogLevelFatal, msg)
}

func (logger rawFieldLogger) Fields(fields Fields) Logger {
	return rawFieldLogger{logger: logger.logger, fields: JoinFields(logger.fields, fields)}
}

func (logger rawFieldLogger) Field(key string, val interface{}) Logger {
	return rawFieldLogger{logger: logger.logger, fields: JoinFields(logger.fields, Fields{key: val})}
}

func (logger rawFieldLogger) Str(key string, val string) Logger {
	return logger.Field(key, val)
}

func (logger rawFieldLogger) Int(key string, val int) Logger {
	return logger.Field(key, val)
}

func (logger rawFieldLogger) Uint(key string, val uint) Logger {
	return logger.Field(key, val)
}

func (logger rawFieldLogger) Bool(key string, val bool) Logger {
	return logger.Field(key, val)
}

func (logger rawFieldLogger) Int64(key string, val int64) Logger {
	return logger.Field(key, val)
}

func (logger rawFieldLogger) Int32(key string, val int32) Logger {
	return logger.Field(key, val)
}

func (logger rawFieldLogger) Int16(key string, val int16) Logger {
	return logger.Field(key, val)
}

func (logger rawFieldLogger) Int8(key string, val int8) Logger {
	return logger.Field(key, val)
}

func (logger rawFieldLogger) Uint64(key string, val uint64) Logger {
	return logger.Field(key, val)
}

func (logger rawFieldLogger) Uint32(key string, val uint32) Logger {
	return logger.Field(key, val)
}

func (logger rawFieldLogger) Uint16(key string, val uint16) Logger {
	return logger.Field(key, val)
}

func (logger rawFieldLogger) Uint8(key string, val uint8) Logger {
	return logger.Field(key, val)
}

func (logger rawFieldLogger) Float32(key string, val float32) Logger {
	return logger.Field(key, val)
}

func (logger rawFieldLogger) Float64(key string, val float64) Logger {
	return logger.Field(key, val)
}

func (logger rawFieldLogger) Err(key string, val error) Logger {
	return logger.Field(key, val)
}

func (logger rawFieldLogger) Time(key string, val time.Time) Logger {
	return logger.Field(key, val)
}

func (logger rawFieldLogger) Duration(key string, val time.Duration) Logger {
	return logger.Field(key, val)
}
