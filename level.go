package gslog

import "strconv"

var logLevelStrings []string

func init() {
	logLevelStrings = []string{
		"ALL",
		"DEBUG",
		"INFO",
		"WARN",
		"ERROR",
		"FATAL",
	}
}

type LogLevel int

const (
	LogLevelAll LogLevel = iota
	LogLevelDebug
	LogLevelInfo
	LogLevelWarn
	LogLevelError
	LogLevelFatal
	LogLevelDisable
)

func (level LogLevel) String() string {
	if level > LogLevelAll && level < LogLevelDisable {
		return logLevelStrings[level]
	}
	return strconv.FormatInt(int64(level), 10)
}
