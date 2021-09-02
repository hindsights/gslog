package gslog

import (
	"testing"
)

func TestLog(t *testing.T) {
	logger := GetFieldLogger("test")
	logger.Debug("hello2", Fields{"x": 123})
	logger.Debug("world", Fields{"y": 123})
	logger.Info("earth")
	logger.Warn("sun")
	logger.Error("moon")
	logger.Info("log with fields", Fields{"key1": "value1", "key2": 234})
	Warn("warn", "str", 123)
	Debug("debug", "str", 234)
	Error("error", "str", 345)
	Info("info", "str", 456)
}
