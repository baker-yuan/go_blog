package filter

import (
	"fmt"
	"time"

	"github.com/baker-yuan/go-blog/go_blog/application/blog/gateway/proxy/biz_ctx"
	"github.com/baker-yuan/go-blog/go_blog/application/blog/gateway/proxy/biz_ctx/http"
	"github.com/baker-yuan/go-blog/go_blog/application/blog/gateway/proxy/config"
	"github.com/baker-yuan/go-blog/go_blog/application/blog/gateway/proxy/ret"
	"github.com/baker-yuan/go-blog/go_blog/application/blog/gateway/proxy/service"
	auth_pb "github.com/baker-yuan/go-blog/go_blog/protocol/auth"
	"trpc.group/trpc-go/trpc-go/client"
	"trpc.group/trpc-go/trpc-go/codec"
)

func ForwardHandler(bizCtx biz_ctx.IBizContext, cfg *config.Config) {
	httpCtx, _ := bizCtx.(*http.HttpContext)
	resource := bizCtx.Value(biz_ctx.ResourceCtxKey).(*auth_pb.Resource)
	instance, err := service.GetOneInstance(cfg.Global.Namespace, resource.Service)
	if err != nil {
		ret.SetFilResponse(httpCtx, err.Error())
		return
	}
	if resource.Protocol == auth_pb.MethodProtocol_Trpc {
		trpcForwardHandler(httpCtx, instance, resource)
	} else {
		httpForwardHandler(httpCtx, instance, resource)
	}
}

func trpcForwardHandler(httpCtx *http.HttpContext, instance string, resource *auth_pb.Resource) {
	opts := []client.Option{
		client.WithNetwork("tcp"),
		client.WithTarget(fmt.Sprintf("ip://%s", instance)),
		client.WithTimeout(time.Duration(8000) * time.Millisecond),
		client.WithProtocol("trpc"),
	}
	// 下游的反序列化方式和上游保持一致
	opts = append(opts, client.WithCurrentSerializationType(codec.SerializationTypeNoop))
	ctx, msg := codec.WithCloneMessage(httpCtx.Context())
	// 服务端寻找对应的方法
	// handler, ok := s.handlers[msg.ServerRPCName()]
	// clientRPCName => serverRPCName
	msg.WithClientRPCName(resource.Method)

	// 指定序列化方式为json
	msg.WithSerializationType(codec.SerializationTypeJSON)

	// 请求头透传
	msg.WithClientMetaData(map[string][]byte{
		"X-User-ID": []byte(httpCtx.Proxy().Header().GetHeader("X-User-ID")),
	})

	// codec.go#RequestProtocol
	rawBody, _ := httpCtx.Request().Body().RawBody()
	reqBody := &codec.Body{Data: rawBody}
	rspBody := &codec.Body{}
	if err := client.DefaultClient.Invoke(ctx, reqBody, rspBody, opts...); err != nil {
		ret.SetFilResponse(httpCtx, err.Error())
		return
	}

	// 设置响应
	ret.SetSucResponse(httpCtx, resource.IsNeedStdRes, rspBody.Data)

	httpCtx.Response().SetHeader("Access-Control-Allow-Origin", "*")
	httpCtx.Response().SetHeader("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
	httpCtx.Response().SetHeader("Access-Control-Allow-Headers", "Content-Type, Authorization")
	httpCtx.Response().SetHeader("Access-Control-Allow-Credentials", "true")

	httpCtx.Response().SetHeader("Content-Type", "application/json")
}

// httpForwardHandler 请求转发处理
func httpForwardHandler(httpCtx *http.HttpContext, instance string, resource *auth_pb.Resource) {
	httpCtx.Proxy().URI().SetPath(resource.Method)
	httpCtx.Proxy().URI().SetHost(instance)
	httpCtx.Proxy().Header().SetHeader("Content-Type", "application/json")
	//httpCtx.Proxy().Header().SetHeader("X-Galileo-Trace-Id", "application/json")
	//httpCtx.Proxy().Header().SetHeader("X-Galileo-Trace-Id", "application/json")

	httpCtx.SendTo("http", service.NewInstance(instance), 3*time.Minute)

	httpCtx.Response().SetHeader("Access-Control-Allow-Origin", "*")
	httpCtx.Response().SetHeader("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
	httpCtx.Response().SetHeader("Access-Control-Allow-Headers", "Content-Type, Authorization")
	httpCtx.Response().SetHeader("Access-Control-Allow-Credentials", "true")

	//httpCtx, _ := bizCtx.(*http.HttpContext)
	//
	//resource := bizCtx.Value(biz_ctx.ResourceCtxKey).(*auth_pb.Resource)
	//
	//instance, err := service.GetOneInstance(cfg.Global.Namespace, resource.Service)
	//if err != nil {
	//	return
	//}
	//
	//// 目标服务器的URL
	//targetURL := fmt.Sprintf("http://%s%s", instance, resource.Url)
	//
	//// 创建一个新的请求对象
	//req := fasthttp.AcquireRequest()
	//defer fasthttp.ReleaseRequest(req)
	//
	//// 复制原始请求的方法、请求URI和正文
	//req.SetRequestURI(targetURL)
	//req.Header.SetMethodBytes([]byte(httpCtx.Request().Method()))
	//body, _ := httpCtx.Request().Body().RawBody()
	//req.SetBody(body)
	//
	//// 复制原始请求的头部（根据需要可以修改或过滤这些头部）
	////ctx.Request.Header.CopyTo(&req.Header)
	//
	//req.Header.Set("Content-Type", "application/json")
	//
	//// 创建一个新的响应对象
	//resp := fasthttp.AcquireResponse()
	//defer fasthttp.ReleaseResponse(resp)
	//
	//// 使用fasthttp的客户端发送请求
	//client := &fasthttp.Client{}
	//if err := client.Do(req, resp); err != nil {
	//	//ctx.Error("failed to forward request: "+err.Error(), fasthttp.StatusInternalServerError)
	//	fmt.Printf("%+v", err)
	//	//log.ErrorContextf(ctx, "")
	//	return
	//}
	//
	//// 将目标服务器的响应复制回原始客户端的响应
	//resp.Header.CopyTo(&ctx.Response.Header)
	//ctx.SetBody(resp.Body())
	//ctx.SetStatusCode(resp.StatusCode())
}
