package peano

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"runtime/debug"
	"strings"
	"time"
)

var (
	// DebugLevel level
	DebugLevel = zapcore.DebugLevel
	// InfoLevel level
	InfoLevel = zapcore.InfoLevel
	// WarnLevel level
	WarnLevel = zapcore.WarnLevel
	// ErrorLevel level
	ErrorLevel = zapcore.ErrorLevel
	// PanicLevel level
	PanicLevel = zapcore.PanicLevel
	// FatalLevel level
	FatalLevel = zapcore.FatalLevel
)

var sugar *zap.SugaredLogger

func getLogLevel(level string) zapcore.Level {
	switch strings.ToLower(level) {
	case "debug":
		return DebugLevel
	case "info":
		return InfoLevel
	case "warn":
		return WarnLevel
	case "error":
		return ErrorLevel
	case "panic":
		return PanicLevel
	case "fatal":
		return FatalLevel
	default:
		return DebugLevel
	}
}

func InitLogger(level, filePath string) {
	encoder := getEncoder()
	writeSyncer := getLogWriter(filePath)
	logLevel := getLogLevel(level)
	core := zapcore.NewCore(encoder, writeSyncer, logLevel)
	//添加将调用函数信息记录到日志中的功能。
	logger := zap.New(core, zap.AddCaller())
	sugar = logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	// 修改时间编码器
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// 在日志文件中使用大写字母记录日志级别
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// NewConsoleEncoder 打印更符合人们观察的方式
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter(filePath string) zapcore.WriteSyncer {
	file, _ := os.Create(filePath)
	return zapcore.AddSync(file)
}

func Debugf(format string, args ...interface{}) {
	sugar.With(zap.Time("time", time.Now())).Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	sugar.With(zap.Time("time", time.Now())).Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	sugar.With(zap.Time("time", time.Now())).Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	sugar.With(zap.Time("time", time.Now())).Errorf(format, args...)
}

func Panicf(format string, args ...interface{}) {
	sugar.With(zap.Time("time", time.Now())).Panicf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	debug.PrintStack()
	sugar.With(zap.Time("time", time.Now())).Fatalf(format, args...)
}

func Error(args ...interface{}) {
	sugar.With(zap.Time("time", time.Now())).Error(args...)
}

func Panic(args ...interface{}) {
	sugar.With(zap.Time("time", time.Now())).Panic(args...)
}

func Fatal(args ...interface{}) {
	sugar.With(zap.Time("time", time.Now())).Fatal(args...)
}
