package log

import (
	"context"

	"github.com/baker-yuan/go-blog/all_packaged_library/base/config"
)

var (
	logger Logger
)

func InitLog() {
	logConf := config.GetLogConf()
	baseConf := config.GetBaseConf()
	var (
		err error
	)
	log, err := NewZapLog(baseConf.AppName, logConf)
	if err != nil {
		panic(err)
	}
	logger = log
}

func Debug(ctx context.Context, format string, args ...interface{}) {
	if len(args) == 0 {
		logger.Debug(ctx, format)
	} else {
		logger.Debug(ctx, format, args...)
	}
}

func Info(ctx context.Context, format string, args ...interface{}) {
	if len(args) == 0 {
		logger.Info(ctx, format)
	} else {
		logger.Info(ctx, format, args...)
	}
}

func Warn(ctx context.Context, format string, args ...interface{}) {
	if len(args) == 0 {
		logger.Warn(ctx, format)
	} else {
		logger.Warn(ctx, format, args...)
	}
}

func Error(ctx context.Context, format string, args ...interface{}) {
	if len(args) == 0 {
		logger.Error(ctx, format)
	} else {
		logger.Error(ctx, format, args...)
	}
}
