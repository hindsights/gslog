package raw

import (
	"testing"

	"github.com/hindsights/gslog"
)

func TestLog(t *testing.T) {
	logger := gslog.GetLogger("test")
	logger.Trace("hello")
	logger.Debug("world")
	logger.Info("earth")
	logger.Warn("sun")
	logger.Error("moon")
}
