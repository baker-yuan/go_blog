package filter

import (
	"github.com/baker-yuan/go-blog/application/blog/gateway/biz_ctx"
	"github.com/baker-yuan/go-blog/application/blog/gateway/config"
)

//// HandlerFunc 定义了中间件处理函数的类型
//type HandlerFunc func(bizCtx biz_ctx.BizContext, next func())
//
//// Chain 构建中间件链
//func Chain(handlers ...HandlerFunc) HandlerFunc {
//	return func(bizCtx biz_ctx.BizContext, next func()) {
//		// 构建中间件链
//		chain := func(current HandlerFunc, next HandlerFunc) HandlerFunc {
//			return func(bizCtx biz_ctx.BizContext, final func()) {
//				current(bizCtx, func() {
//					next(bizCtx, final)
//				})
//			}
//		}
//
//		// 最后一个中间件调用next
//		final := func(bizCtx biz_ctx.BizContext, final func()) {
//			next()
//		}
//
//		// 链接所有中间件
//		for i := len(handlers) - 1; i >= 0; i-- {
//			final = chain(handlers[i], final)
//		}
//
//		// 执行链
//		final(bizCtx, next)
//	}
//}

// HandlerFunc 定义了业务处理函数的类型
type HandlerFunc func(bizCtx biz_ctx.BizContext, cfg *config.Config)

// Middleware 定义了中间件的类型
type Middleware func(HandlerFunc) HandlerFunc

// Chain 构建中间件链
func Chain(middlewares ...Middleware) Middleware {
	return func(final HandlerFunc) HandlerFunc {
		// 如果没有中间件，则直接返回最终处理函数
		if len(middlewares) == 0 {
			return final
		}
		// 将中间件链中的最后一个中间件包装在最终处理函数周围
		wrapped := middlewares[len(middlewares)-1](final)
		// 逆序遍历中间件切片，将每个中间件包装在前一个中间件的外层
		for i := len(middlewares) - 2; i >= 0; i-- {
			wrapped = middlewares[i](wrapped)
		}
		// 返回包装后的处理函数
		return wrapped
	}
}
