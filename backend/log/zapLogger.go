package log

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type zapLogger struct {
	sugarLogger *zap.SugaredLogger
}

func (z *zapLogger) Debug(args ...interface{}) {
	z.sugarLogger.Debug(args)
}

func (z *zapLogger) Info(args ...interface{}) {
	z.sugarLogger.Info(args)
}

func (z *zapLogger) Warn(args ...interface{}) {
	z.sugarLogger.Warn(args)
}

func (z *zapLogger) Error(args ...interface{}) {
	z.sugarLogger.Error(args)
}

func (z *zapLogger) Panic(args ...interface{}) {
	z.sugarLogger.Panic(args)
}

func (z *zapLogger) Fatal(args ...interface{}) {
	z.sugarLogger.Fatal(args)
}

func (z *zapLogger) With(args ...interface{}) *Logger {
	var log Logger
	log = &zapLogger{sugarLogger: z.sugarLogger.With(args)}
	return &log
}

func (z *zapLogger) Close() {
	_ = z.sugarLogger.Sync()
}

func initZapLogger() (Logger, *zapLogger) {
	fileWriter := getLogFileWriter()
	encoderFile := getFileEncoder()
	encoderConsole := getConsoleEncoder()

	core := zapcore.NewTee(
		zapcore.NewCore(encoderFile, fileWriter, zapcore.DebugLevel),
		zapcore.NewCore(encoderConsole, zapcore.AddSync(os.Stdout), zapcore.DebugLevel),
	)

	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel)).Sugar()
	logger.Debug("Logger Init(ed)")

	var zapper *zapLogger
	zapper = &zapLogger{sugarLogger: logger}
	return zapper, zapper
}

func getConsoleEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getFileEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogFileWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger {
		Filename: "./console.log",
		MaxSize: 10,
		MaxBackups: 5,
		MaxAge: 30,
		Compress: false,
	}
	return zapcore.AddSync(lumberJackLogger)
}
