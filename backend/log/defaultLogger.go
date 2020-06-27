package log

import "log"

type defaultLogger struct {}

func initDefaultLogger() Logger {
	var log Logger
	log = &defaultLogger{}
	return log
}

func (d *defaultLogger) Debug(args ...interface{}) {
	log.Println(args)
}

func (d *defaultLogger) Info(args ...interface{}) {
	log.Println(args)
}

func (d *defaultLogger) Warn(args ...interface{}) {
	log.Println(args)
}

func (d *defaultLogger) Error(args ...interface{}) {
	log.Println(args)
}

func (d *defaultLogger) Panic(args ...interface{}) {
	log.Println(args)
}

func (d *defaultLogger) Fatal(args ...interface{}) {
	log.Println(args)
}

func (d *defaultLogger) With(args ...interface{}) *Logger {
	var log Logger
	log = d
	return &log
}

func (d *defaultLogger) Close() {}
