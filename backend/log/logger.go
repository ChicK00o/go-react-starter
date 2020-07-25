package log

import (
	"github.com/ChicK00o/container"
	"go.uber.org/zap"
)

type Logger interface {
	With(args ...interface{}) *Logger
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Panic(args ...interface{})
	Fatal(args ...interface{})
	Close()
}

var createdDefault = false
var zapper *zapLogger = nil

func init() {
	container.Singleton(func() Logger {
		return newLogger(false)
	})
}

func newLogger(wantDefault bool) Logger {
	if wantDefault {
		createdDefault = true
		return initDefaultLogger()
	}

	createdDefault = false

	var logger Logger
	logger, zapper = initZapLogger()
	return logger
}

func IsZapLogger() bool {
	return createdDefault == false
}

func GetZapLogger() *zap.SugaredLogger {
	if IsZapLogger() {
		return zapper.sugarLogger
	}
	return nil
}
