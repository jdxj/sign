package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/jdxj/sign/internal/pkg/config"
)

var (
	logger *zap.Logger
	sugar  *zap.SugaredLogger
)

func Init(conf config.Logger) {
	core := zapcore.NewCore(
		newEncoder(conf.Path),
		newSyncer(conf.Path),
		newLevel(conf.Path),
	)
	logger = zap.New(
		core,
		zap.AddCaller(),
		zap.AddCallerSkip(1),
		zap.AddStacktrace(zap.ErrorLevel),
	)
	sugar = logger.Sugar()
}

func newEncoder(path string) zapcore.Encoder {
	conf := zap.NewProductionEncoderConfig()
	conf.EncodeTime = zapcore.ISO8601TimeEncoder

	if path == "" {
		return zapcore.NewConsoleEncoder(conf)
	}
	return zapcore.NewJSONEncoder(conf)
}

func newSyncer(path string) zapcore.WriteSyncer {
	if path == "" {
		return zapcore.AddSync(os.Stdout)
	}
	return zapcore.AddSync(&lumberjack.Logger{
		Filename:  path,
		LocalTime: true,
	})
}

func newLevel(path string) zapcore.Level {
	return zap.DebugLevel
}

func Debugf(template string, args ...interface{}) {
	sugar.Debugf(template, args...)
}

func Infof(template string, args ...interface{}) {
	sugar.Infof(template, args...)
}

func Warnf(template string, args ...interface{}) {
	sugar.Warnf(template, args...)
}

func Errorf(template string, args ...interface{}) {
	sugar.Errorf(template, args...)
}

func Sync() {
	_ = logger.Sync()
}
