package logger

import (
	"errors"
	"os"
	"path/filepath"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	envPodName = "POD_NAME"
	logExt     = ".log"
)

var (
	ErrInvalidLogPath = errors.New("invalid log path")
)

var (
	logger *zap.Logger
	sugar  *zap.SugaredLogger
)

func Init(path, base string) {
	if os.Getenv(envPodName) != "" {
		base = os.Getenv(envPodName)
	}
	if base == "" {
		panic(ErrInvalidLogPath)
	}
	base += logExt

	core := zapcore.NewCore(
		newEncoder(path),
		newSyncer(path, base),
		newLevel(),
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
		conf.EncodeLevel = zapcore.LowercaseColorLevelEncoder
		return zapcore.NewConsoleEncoder(conf)
	}
	return zapcore.NewJSONEncoder(conf)
}

func newSyncer(path, base string) zapcore.WriteSyncer {
	if path == "" {
		return zapcore.AddSync(os.Stdout)
	}
	bws := &zapcore.BufferedWriteSyncer{
		WS: zapcore.AddSync(&lumberjack.Logger{
			Filename:  filepath.Join(path, base),
			LocalTime: true,
		}),
		FlushInterval: 5 * time.Second,
	}
	return bws
}

func newLevel() zapcore.Level {
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
