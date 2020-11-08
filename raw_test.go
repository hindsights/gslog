package gslog

import (
	"testing"
)

func TestLog(t *testing.T) {
	logger := GetLogger("test")
	logger.Trace("hello", "x", 123)
	logger.Debug("world", "y", 123)
	logger.Info("earth")
	logger.Warn("sun")
	logger.Error("moon")
	logger.WithFields(Fields{"key1": "value1", "key2": "values"}).Info("log with fields")
	Warn("warn")
	Debug("debug")
	Error("error")
	Info("info")
}
