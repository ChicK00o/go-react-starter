package log

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

func NewLogger(wantDefault bool) Logger {
	if !isLoggerNil() {
		internalLogger.Error("Logger is already created")
		return internalLogger
	}
	if wantDefault {
		internalLogger = initDefaultLogger()
		return internalLogger
	}
	internalLogger = initZapLogger()
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
