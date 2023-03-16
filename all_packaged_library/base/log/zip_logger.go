package log

import (
	"context"
	"fmt"

	"gitee.com/baker-yuan/go-blog/all_packaged_library/base/config"
	"go.uber.org/zap"
)

type zapLog struct {
	log *zap.Logger
}

func NewZapLog(appName string, conf *config.LogConf) (Logger, error) {
	var (
		cfg zap.Config
	)
	cfg = zap.Config{
		Level:         zap.NewAtomicLevelAt(zap.DebugLevel),
		Development:   true,
		Encoding:      "console",
		EncoderConfig: zap.NewDevelopmentEncoderConfig(),
		OutputPaths: []string{
			conf.Path,
			"stdout",
		},
		ErrorOutputPaths: []string{
			conf.Path,
			"stdout",
		},
	}
	build, err := cfg.Build()
	if err != nil {
		return nil, err
	}
	return zapLog{
		log: build,
	}, nil
}

func (z zapLog) Debug(ctx context.Context, format string, args ...interface{}) {
	var (
		msg string
	)
	if len(args) != 0 {
		msg = fmt.Sprintf(format, args...)
	} else {
		msg = format
	}
	z.log.Debug(msg)
}

func (z zapLog) Info(ctx context.Context, format string, args ...interface{}) {
	var (
		msg string
	)
	msg = fmt.Sprintf(format, args)
	z.log.Info(msg)
}

func (z zapLog) Warn(ctx context.Context, format string, args ...interface{}) {
	var (
		msg string
	)
	msg = fmt.Sprintf(format, args)
	z.log.Warn(msg)
}

func (z zapLog) Error(ctx context.Context, format string, args ...interface{}) {
	var (
		msg string
	)
	msg = fmt.Sprintf(format, args)
	z.log.Error(msg)
}
