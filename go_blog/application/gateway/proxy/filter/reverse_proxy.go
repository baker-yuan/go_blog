package filter

import (
	"time"

	"github.com/baker-yuan/go-blog/application/blog/gateway/proxy/biz_ctx"
	"github.com/baker-yuan/go-blog/application/blog/gateway/proxy/biz_ctx/http"
	"github.com/baker-yuan/go-blog/application/blog/gateway/proxy/config"
	"github.com/baker-yuan/go-blog/application/blog/gateway/proxy/service"
	auth_pb "github.com/baker-yuan/go-blog/protocol/auth"
)

// ForwardHandler 请求转发处理
func ForwardHandler(bizCtx biz_ctx.IBizContext, cfg *config.Config) {

	httpCtx, _ := bizCtx.(*http.HttpContext)

	resource := bizCtx.Value(biz_ctx.ResourceCtxKey).(*auth_pb.Resource)
	instance, err := service.GetOneInstance(cfg.Global.Namespace, resource.Service)
	if err != nil {
		return
	}

	httpCtx.Proxy().URI().SetPath(resource.Url)
	httpCtx.Proxy().URI().SetHost(instance)
	httpCtx.Proxy().Header().SetHeader("Content-Type", "application/json")

	httpCtx.SendTo("http", service.InstanceImpl{}, 3*time.Minute)

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
