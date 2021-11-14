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
	logger := gslog.GetSimpleLogger("app")
	logger.Debug("debug", 1, args)
	logger.Info("info", "abc")
	logger.Warn("warn", true)
	logger.Error("error", false)
	logger.Debugf("debugf %s", "name")
	logger.Infof("infof %s", "value")
	logger.Warnf("warnf %d", 20)
	logger.Errorf("errorf %v", 100)

	slogger := gslog.GetLogger("test")
	slogger.Int("integer", 123).Debug("debug output")
	slogger.Str("string", "value").Info("info output")
	slogger.Bool("bool", true).Warn("warn output")
	slogger.Err("error", io.EOF).Error("error output")

	flogger := slogger.Fields(gslog.Fields{"ip": "1.2.3.4", "port": "880"})
	flogger.Int("time", 1).Str("date", "2020-12-30").Info("detect start")
}
