package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	desugar *zap.Logger
	sugar   *zap.SugaredLogger
)

type OptionFunc func(opts *Options)

type Options struct {
	Mode     string
	FileName string

	MaxSize    int
	MaxAge     int
	MaxBackups int

	LocalTime bool
	Compress  bool
}

func WithMode(mode string) OptionFunc {
	return func(opts *Options) {
		opts.Mode = mode
	}
}

func Init(path string, optsF ...OptionFunc) {
	opts := &Options{
		Mode:       "debug",
		FileName:   path,
		MaxSize:    50,
		MaxAge:     30,
		MaxBackups: 10,
		LocalTime:  true,
		Compress:   false,
	}
	for _, optF := range optsF {
		optF(opts)
	}

	core := zapcore.NewCore(
		encoder(opts.Mode),
		syncer(opts),
		level(opts.Mode),
	)
	desugar = zap.New(core, zap.AddCaller())
	sugar = desugar.Sugar()
}

func syncer(opts *Options) zapcore.WriteSyncer {
	var syncer zapcore.WriteSyncer
	switch opts.Mode {
	case "debug":
		syncer = zapcore.AddSync(os.Stdout)

	case "release":
		rotation := &lumberjack.Logger{
			Filename:   opts.FileName,
			MaxSize:    opts.MaxSize,
			MaxAge:     opts.MaxAge,
			MaxBackups: opts.MaxBackups,
			LocalTime:  opts.LocalTime,
			Compress:   opts.Compress,
		}
		syncer = zapcore.AddSync(rotation)
	}
	return syncer
}

func encoder(mode string) zapcore.Encoder {
	var encoder zapcore.Encoder
	switch mode {
	case "debug":
		cfg := zap.NewDevelopmentEncoderConfig()
		cfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
		cfg.EncodeCaller = zapcore.FullCallerEncoder
		encoder = zapcore.NewConsoleEncoder(cfg)

	case "release":
		cfg := zap.NewProductionEncoderConfig()
		cfg.EncodeTime = zapcore.ISO8601TimeEncoder
		encoder = zapcore.NewJSONEncoder(cfg)
	}
	return encoder
}

func level(mode string) zapcore.Level {
	switch mode {
	case "release":
		return zap.InfoLevel
	}
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
	_ = desugar.Sync()
}
