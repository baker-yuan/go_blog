package filter

import (
	"github.com/baker-yuan/go-blog/application/blog/gateway/biz_ctx"
)

// HandlerFunc 定义了中间件处理函数的类型
type HandlerFunc func(bizCtx *biz_ctx.HttpContext, next func())

// Chain 构建中间件链
func Chain(handlers ...HandlerFunc) HandlerFunc {
	return func(bizCtx *biz_ctx.HttpContext, next func()) {
		// 构建中间件链
		chain := func(current HandlerFunc, next HandlerFunc) HandlerFunc {
			return func(bizCtx *biz_ctx.HttpContext, final func()) {
				current(bizCtx, func() {
					next(bizCtx, final)
				})
			}
		}

		// 最后一个中间件调用next
		final := func(bizCtx *biz_ctx.HttpContext, final func()) {
			next()
		}

		// 链接所有中间件
		for i := len(handlers) - 1; i >= 0; i-- {
			final = chain(handlers[i], final)
		}

		// 执行链
		final(bizCtx, next)
	}
}
