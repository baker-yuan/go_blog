package log

import (
	"context"
)

// Level 日志等级
type Level int

// Enums 日志等级常量
const (
	LevelDebug = 1 // debug
	LevelInfo  = 2 // info
	LevelWarn  = 3 // warn
	LevelError = 4 // error
)

var LevelStrings = map[Level]string{
	LevelDebug: "debug",
	LevelInfo:  "info",
	LevelWarn:  "warn",
	LevelError: "error",
}

var LevelNames = map[string]Level{
	"debug": LevelDebug,
	"info":  LevelInfo,
	"warn":  LevelWarn,
	"error": LevelError,
}

// Logger 日志抽象
type Logger interface {
	Debug(ctx context.Context, format string, args ...interface{})
	Info(ctx context.Context, format string, args ...interface{})
	Warn(ctx context.Context, format string, args ...interface{})
	Error(ctx context.Context, format string, args ...interface{})
}
