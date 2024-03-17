package http_proxy

import (
	"context"

	"github.com/baker-yuan/go-blog/application/blog/gateway/biz_ctx"
	"github.com/baker-yuan/go-blog/application/blog/gateway/config"
	"github.com/baker-yuan/go-blog/application/blog/gateway/filter"
	"github.com/valyala/fasthttp"
	"trpc.group/trpc-go/trpc-go/log"
)

var (
	httpServer *fasthttp.Server
)

func HttpServerRun(ctx context.Context, cfg *config.Config) {
	// 监听客户端请求
	httpServer = &fasthttp.Server{
		Handler: func(ctx *fasthttp.RequestCtx) {
			// OPTIONS请求直接返回
			if string(ctx.Method()) == fasthttp.MethodOptions {
				ctx.SetStatusCode(fasthttp.StatusNoContent)
				return
			}
			bizCtx := biz_ctx.NewContext(ctx, cfg)
			// 构建中间件链并执行
			handler := filter.Chain(
				filter.Recovery,
				filter.RequestLog,
			)(filter.ForwardHandler)
			handler(bizCtx, cfg)
		},
	}
	// 监听并服务
	if err := httpServer.ListenAndServe(cfg.Http.Addr); err != nil {
		log.Fatalf("server listenAndServe fail err: %+v", err)
	}
}

func HttpServerStop() {
	_ = httpServer.Shutdown()
}
