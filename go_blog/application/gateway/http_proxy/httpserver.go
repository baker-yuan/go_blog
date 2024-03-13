package http_proxy

import (
	"net/http"

	"github.com/baker-yuan/go-blog/application/blog/gateway/biz_ctx"
	"github.com/baker-yuan/go-blog/application/blog/gateway/filter"
	auth_pb "github.com/baker-yuan/go-blog/protocol/auth"
	"github.com/valyala/fasthttp"
	"trpc.group/trpc-go/trpc-go/client"
	"trpc.group/trpc-go/trpc-go/log"
)

var (
	HttpSrvHandler *http.Server
)

func HttpServerRun() {
	// 读取接口配置
	proxy := auth_pb.NewAuthApiClientProxy(
		client.WithTarget("ip://127.0.0.1:8000"),
	)
	//resourceRsp, err := proxy.GetEffectiveResource(context.Background(), &auth_pb.GetEffectiveResourceReq{})
	//if err != nil {
	//	return
	//}
	//// 初始化路由信息
	//router.LoadResourceList(resourceRsp.Data)

	proxy = proxy

	// 监听客户端请求
	server := &fasthttp.Server{
		Handler: func(ctx *fasthttp.RequestCtx) {
			bizCtx := biz_ctx.NewContext(ctx)
			// 构建中间件链并执行
			filter.Chain(filter.Recovery, filter.RequestLog)(bizCtx, func() {
				// 最后执行请求转发
				filter.ForwardHandler(bizCtx)
			})
		},
	}
	// 监听并服务
	if err := server.ListenAndServe(":8090"); err != nil {
		log.Fatalf("server listenAndServe fail err: %+v", err)
	}
}

func HttpServerStop() {

}
