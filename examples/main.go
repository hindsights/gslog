package main

import (
	"fmt"

	"github.com/hindsights/gslog"
)

func main() {
	fmt.Println("test")
	gslog.Info("start")
	logger := gslog.GetLogger("app")
	logger.Debug("debug", 1)
	logger.Info("info", "abc")
	logger.Warn("warn", true)
	logger.Error("error", false)
	logger.WithFields(gslog.Fields{"key1": 1, "key2": "val2"}).Error("field output")
	logger.WithFields(gslog.Fields{"key1": 1, "key2": "val2"}).Errorf("field output %d", 567)
	gslog.Debugf("debugf %s", "name")
	gslog.Infof("infof %s", "value")
	gslog.Warnf("warnf %d", 20)
	gslog.Errorf("errorf %v", 100)
}
