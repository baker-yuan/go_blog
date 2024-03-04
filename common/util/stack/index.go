package stack

import (
	"context"
	"runtime"

	"trpc.group/trpc-go/trpc-go/log"
)

// StackUtils 堆栈处理
type StackUtils struct {
}

// Stack 获取堆栈信息
func (c StackUtils) Stack() string {
	return stack()
}

// CatchPanic 捕获panic
func (c StackUtils) CatchPanic(ctx context.Context) {
	if err := recover(); err != nil {
		log.ErrorContextf(ctx, "catch panic: %+v, stack: %s", err, stack())
	}
}

func stack() string {
	buf := make([]byte, 1024)
	n := runtime.Stack(buf, false)
	return string(buf[:n])
}
