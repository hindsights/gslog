package main

import (
	"fmt"
	"io"

	"github.com/hindsights/gslog"
)

func main() {
	fmt.Println("test")
	gslog.Info("start")
	gslog.Info("start ok", gslog.Fields{"name": 123, "str": "string literal"})
	args := []interface{}{123, "abc", true}
	logger := gslog.GetLogger("app")
	logger.Debug("debug", 1, args)
	logger.Info("info", "abc")
	logger.Warn("warn", true)
	logger.Error("error", false)
	logger.Debugf("debugf %s", "name")
	logger.Infof("infof %s", "value")
	logger.Warnf("warnf %d", 20)
	logger.Errorf("errorf %v", 100)

	slogger := gslog.GetFieldLogger("test")
	slogger.Debug("debug output", gslog.Fields{"integer": 123})
	slogger.Info("info output", gslog.Fields{"string": "value"})
	slogger.Warn("warn output", gslog.Fields{"bool": true})
	slogger.Error("error output", gslog.Fields{"error": io.EOF})

	flogger := slogger.WithFields(gslog.Fields{"ip": "1.2.3.4", "port": "880"})
	flogger.Info("detect start", gslog.Fields{"time": 1, "date": "2020-12-30"})
}
