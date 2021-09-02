# gslog

A structured log interface library for golang, like slf4j.

## gslog.Logger

Structured logger interface.

## gslog.Backend

Interface for logging service provider.

## Example

```go
package main

import (
	"fmt"
	"os"

	"github.com/hindsights/gslog"
	"github.com/hindsights/gslog-logrus/gslogrus"
	"github.com/hindsights/gslog-zap/gszap"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type logLevelChecker struct {
	level zapcore.Level
}

func (checker logLevelChecker) Enabled(l zapcore.Level) bool {
	return l >= checker.level
}

func main() {
	fmt.Println("test")
	gslog.Info("start")
	logger := gslog.GetLogger("app")
	flogger := gslog.GetFieldLogger("app")
	logger.Debug("debug", gslog.Fields{"key1": 1, "key2": "val2"})
	logger.Info("info", gslog.Fields{"int": 1, "str": "val2"})
	logger.Warn("warn")
	logger.Error("error", gslog.Fields{"key1": 1, "key2": "val2"})
	flogger.Error("field output")
	flogger.Info("field output", gslog.Fields{"val": 567})
	flogger.WithFields(gslog.Fields{"key1": 1, "key2": "val2"}).Error("field output")
	flogger.WithFields(gslog.Fields{"key1": 1, "key2": "val2"}).Info("field output", gslog.Fields{"val": 567})
	gslog.Debugf("debugf %s", "name")
	gslog.Infof("infof %s", "value")
	gslog.Warnf("warnf %d", 20)
	gslog.Errorf("errorf %v", 100)

	logrusLogger := logrus.New()
	logrusLogger.SetFormatter(&logrus.TextFormatter{
		DisableColors:  true,
		FullTimestamp:  true,
		DisableSorting: true,
	})
	gslog.SetBackend(gslogrus.NewBackend(logrusLogger))
	gslog.Info("gs-logrus-hello")
	logger = gslog.GetLogger("logrus")
	flogger = gslog.GetFieldLogger("logrus")
	logger.Info("output to zap", 123)
	flogger.Info("output to zap", gslog.Fields{"value": 123})

	consoleWriter := zapcore.Lock(os.Stdout)
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)
	core := zapcore.NewCore(consoleEncoder, consoleWriter, logLevelChecker{level: zapcore.DebugLevel})
	tempLogger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	tempLogger = tempLogger.WithOptions(zap.AddCallerSkip(1))
	gslog.SetBackend(gszap.NewBackend(tempLogger))
	gslog.Info("gs-zap-hello")
	gslog.Warn("zap-start")
	logger = gslog.GetLogger("zap")
	flogger = gslog.GetFieldLogger("zap")
	logger.Info("output to zap", 123)
	flogger.Info("output to zap", gslog.Fields{"value": 123})
}
```