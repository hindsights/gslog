package raw

import (
	"testing"

	"github.com/hindsights/gslog"
)

func TestLog(t *testing.T) {
	logger := gslog.GetLogger("test")
	logger.Trace("hello", "x", 123)
	logger.Debug("world", "y", 123)
	logger.Info("earth")
	logger.Warn("sun")
	logger.Error("moon")
	logger.WithFields(gslog.Fields{"key1": "value1", "key2": "values"}).Info("log with fields")
}
