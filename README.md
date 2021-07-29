# gslog

A simple log interface library for golang, like slf4j.

## gslog.Logger

Simple logger interface.

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

	logrusLogger := logrus.New()
	logrusLogger.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
		// DisableQuote:   true,
		DisableSorting: true,
	})
	gslog.SetBackend(gslogrus.NewBackend(logrusLogger, gslog.LogLevelAll))
	gslog.Info("gs-logrus-hello")

	consoleWriter := zapcore.Lock(os.Stdout)
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)
	core := zapcore.NewCore(consoleEncoder, consoleWriter, logLevelChecker{level: zapcore.DebugLevel})
	tempLogger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	tempLogger = tempLogger.WithOptions(zap.AddCallerSkip(1))
	gslog.SetBackend(gszap.NewBackend(gslog.LogLevelAll, tempLogger))
	gslog.Info("gs-zap-hello")
	gslog.Warn("zap-start")
	logger.Info("output to zap")
}

```