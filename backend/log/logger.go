package log

import "go.uber.org/zap"

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

var internalLogger Logger = nil
var createdDefault bool = false
var zapper *zapLogger = nil

func NewLogger(wantDefault bool) Logger {
	if !isLoggerNil() {
		internalLogger.Error("Logger is already created")
		return internalLogger
	}
	if wantDefault {
		createdDefault = true
		internalLogger = initDefaultLogger()
		return internalLogger
	}
	createdDefault = false
	internalLogger, zapper = initZapLogger()
	return internalLogger
}

func isLoggerNil() bool {
	return internalLogger == nil
}

func Instance() Logger {
	if isLoggerNil() {
		_ = NewLogger(false)
	}
	return internalLogger
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
