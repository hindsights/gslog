package gslog

import (
	"time"
)

type Attr struct {
	Key   string
	Value interface{}
}

type SugaredLogger interface {
	NeedLog(level LogLevel) bool
	LogDirect(level LogLevel, args ...interface{})
	Log(level LogLevel, args ...interface{})
	LogfDirect(level LogLevel, format string, args ...interface{})
	Logf(level LogLevel, format string, args ...interface{})

	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})

	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
}

type Logger interface {
	NeedLog(level LogLevel) bool
	LogDirect(level LogLevel, msg string, args ...interface{})
	Log(level LogLevel, msg string, args ...interface{})

	Debug(msg string, args ...interface{})
	Info(msg string, args ...interface{})
	Warn(msg string, args ...interface{})
	Error(msg string, args ...interface{})
	Fatal(msg string, args ...interface{})

	Fields(fields Fields) Logger
	Field(key string, val interface{}) Logger
	WithAttrs(attrs ...Attr) Logger
	With(args ...interface{}) Logger

	Str(key string, val string) Logger
	Int(key string, val int) Logger
	Uint(key string, val uint) Logger
	Bool(key string, val bool) Logger

	Int64(key string, val int64) Logger
	Int32(key string, val int32) Logger
	Int16(key string, val int16) Logger
	Int8(key string, val int8) Logger

	Uint64(key string, val uint64) Logger
	Uint32(key string, val uint32) Logger
	Uint16(key string, val uint16) Logger
	Uint8(key string, val uint8) Logger

	Float32(key string, val float32) Logger
	Float64(key string, val float64) Logger

	Err(key string, val error) Logger
	Time(key string, val time.Time) Logger
	Duration(key string, val time.Duration) Logger
}
