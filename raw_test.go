package gslog

import (
	"fmt"
	"testing"
)

func TestLog(t *testing.T) {
	logger := GetLogger("test")
	logger.Int("x", 123).Debug("hello2")
	logger.Int("y", 123).Debug("world")
	logger.Info("earth")
	logger.Warn("sun")
	logger.Error("moon")
	logger.Str("key1", "val1").Int("key2", 234).Info("log with fields")
	logger.Warn("bad things", "key1", 123, "key2", "val2")
	logger.With("key0", "val0", "boolkey", true, "intkey", 345).Warn("bad things", "key1", 123, "key2", "val2")
	Warn("warn", "str", 123)
	Debug("debug", "str", 234)
	Error("error", "str", 345)
	Info("info", "str", 456)
	args := []interface{}{"abc", 1, "def", "xxx"}
	attrs := ToAttrs(args)
	fmt.Println("attrs", attrs)
}
