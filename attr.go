package gslog

import (
	"fmt"
	"time"
)

const (
	ErrorKey      = "err"
	LoggerNameKey = "ctx"
)

const badKey = "<badkey>"

func Any(key string, val interface{}) Attr {
	return makeAttr(key, val)
}

func String(key string, val string) Attr {
	return Any(key, val)
}

func Bool(key string, val bool) Attr {
	return Any(key, val)
}

func Int(key string, val int) Attr {
	return Any(key, val)
}

func Int64(key string, val int64) Attr {
	return Any(key, val)
}

func Uint64(key string, val uint64) Attr {
	return Any(key, val)
}

func Int32(key string, val int32) Attr {
	return Any(key, val)
}

func Uint32(key string, val uint32) Attr {
	return Any(key, val)
}

func MakeAttr(key string, val interface{}) Attr {
	return makeAttr(key, val)
}

func makeAttr(key string, val interface{}) Attr {
	return Attr{Key: key, Value: val}
}

func extractAttr(args []interface{}) (Attr, []interface{}) {
	switch x := args[0].(type) {
	case string:
		if len(args) >= 2 {
			return makeAttr(x, args[1]), args[2:]
		}
		return makeAttr(x, args[1]), nil
	case Attr:
		return x, args[1:]
	default:
		return makeAttr(badKey, x), args[1:]
	}
}

func ToAttrs(args []interface{}) []Attr {
	attrs := make([]Attr, 0, len(args))
	var attr Attr
	for {
		if len(args) == 0 {
			break
		}
		attr, args = extractAttr(args)
		attrs = append(attrs, attr)
	}
	return attrs
}

func CopyAttrs(attrs []Attr) []Attr {
	if len(attrs) == 0 {
		return nil
	}
	return JoinAttrs(attrs)
}

func JoinAttrs(attrsList ...[]Attr) []Attr {
	if len(attrsList) == 0 {
		return nil
	}
	var ret []Attr
	for _, attrs := range attrsList {
		ret = append(ret, attrs...)
	}
	return ret
}

func FormatAttrs(attrs []Attr) []string {
	if len(attrs) == 0 {
		return nil
	}
	ret := make([]string, 0, len(attrs))
	for _, attr := range attrs {
		ret = append(ret, fmt.Sprintf("%s=%v", attr.Key, attr.Value))
	}
	return ret
}

type rawAttrLogger struct {
	logger SugaredLogger
	attrs  []Attr
}

func NewAttrLogger(logger SugaredLogger) Logger {
	return rawAttrLogger{logger: logger}
}

func (logger rawAttrLogger) NeedLog(level LogLevel) bool {
	return logger.logger.NeedLog(level)
}

func (logger rawAttrLogger) Log(level LogLevel, msg string, args ...interface{}) {
	logger.LogDirect(level, msg, args)
}

func (logger rawAttrLogger) LogDirect(level LogLevel, msg string, args ...interface{}) {
	if !logger.NeedLog(level) {
		return
	}
	var attrs []Attr
	attrs = append(attrs, logger.attrs...)
	attrs = append(attrs, ToAttrs(args)...)
	strArgs := FormatAttrs(attrs)
	logargs := make([]interface{}, 2+len(strArgs))
	logargs[0] = msg
	logargs[1] = "\t"
	for i, arg := range strArgs {
		logargs[i+2] = arg
	}
	logger.logger.Log(level, logargs...)
}

func (logger rawAttrLogger) Debug(msg string, args ...interface{}) {
	logger.LogDirect(LogLevelDebug, msg, args...)
}

func (logger rawAttrLogger) Info(msg string, args ...interface{}) {
	logger.LogDirect(LogLevelInfo, msg, args...)
}

func (logger rawAttrLogger) Warn(msg string, args ...interface{}) {
	logger.LogDirect(LogLevelWarn, msg, args...)
}

func (logger rawAttrLogger) Error(msg string, args ...interface{}) {
	logger.LogDirect(LogLevelError, msg, args...)
}

func (logger rawAttrLogger) Fatal(msg string, args ...interface{}) {
	logger.LogDirect(LogLevelFatal, msg, args...)
}

func (logger rawAttrLogger) Fields(fields Fields) Logger {
	attrs := make([]Attr, 0, len(fields))
	for k, v := range fields {
		attrs = append(attrs, makeAttr(k, v))
	}
	return logger.WithAttrs(attrs...)
}

func (logger rawAttrLogger) Field(key string, val interface{}) Logger {
	return logger.withAttr(key, val)
}

func (logger rawAttrLogger) withAttr(key string, val interface{}) Logger {
	return logger.WithAttrs(makeAttr(key, val))
}

func (logger rawAttrLogger) WithAttrs(attrs ...Attr) Logger {
	return rawAttrLogger{logger: logger.logger, attrs: JoinAttrs(logger.attrs, attrs)}
}

func (logger rawAttrLogger) With(args ...interface{}) Logger {
	return rawAttrLogger{logger: logger.logger, attrs: JoinAttrs(logger.attrs, ToAttrs(args))}
}

func (logger rawAttrLogger) Str(key string, val string) Logger {
	return logger.withAttr(key, val)
}

func (logger rawAttrLogger) Int(key string, val int) Logger {
	return logger.withAttr(key, val)
}

func (logger rawAttrLogger) Uint(key string, val uint) Logger {
	return logger.withAttr(key, val)
}

func (logger rawAttrLogger) Bool(key string, val bool) Logger {
	return logger.withAttr(key, val)
}

func (logger rawAttrLogger) Int64(key string, val int64) Logger {
	return logger.withAttr(key, val)
}

func (logger rawAttrLogger) Int32(key string, val int32) Logger {
	return logger.withAttr(key, val)
}

func (logger rawAttrLogger) Int16(key string, val int16) Logger {
	return logger.withAttr(key, val)
}

func (logger rawAttrLogger) Int8(key string, val int8) Logger {
	return logger.withAttr(key, val)
}

func (logger rawAttrLogger) Uint64(key string, val uint64) Logger {
	return logger.withAttr(key, val)
}

func (logger rawAttrLogger) Uint32(key string, val uint32) Logger {
	return logger.withAttr(key, val)
}

func (logger rawAttrLogger) Uint16(key string, val uint16) Logger {
	return logger.withAttr(key, val)
}

func (logger rawAttrLogger) Uint8(key string, val uint8) Logger {
	return logger.withAttr(key, val)
}

func (logger rawAttrLogger) Float32(key string, val float32) Logger {
	return logger.withAttr(key, val)
}

func (logger rawAttrLogger) Float64(key string, val float64) Logger {
	return logger.withAttr(key, val)
}

func (logger rawAttrLogger) Err(key string, val error) Logger {
	return logger.withAttr(key, val)
}

func (logger rawAttrLogger) Time(key string, val time.Time) Logger {
	return logger.withAttr(key, val)
}

func (logger rawAttrLogger) Duration(key string, val time.Duration) Logger {
	return logger.withAttr(key, val)
}
