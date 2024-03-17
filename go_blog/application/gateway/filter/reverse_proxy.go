package filter

import (
	"fmt"

	"github.com/baker-yuan/go-blog/application/blog/gateway/biz_ctx"
	"github.com/baker-yuan/go-blog/application/blog/gateway/config"
	"github.com/baker-yuan/go-blog/application/blog/gateway/router"
	"github.com/baker-yuan/go-blog/application/blog/gateway/service"
	"github.com/valyala/fasthttp"
)

// ForwardHandler 请求转发处理
func ForwardHandler(bizCtx biz_ctx.BizContext, cfg *config.Config) {
	httpCtx, _ := bizCtx.(*biz_ctx.HttpContext)
	ctx := httpCtx.FastCtx()

	resource := router.MatchResource(bizCtx)
	if resource == nil {
		return
	}

	instance, err := service.GetOneInstance(cfg.Global.Namespace, resource.Service)
	if err != nil {
		return
	}

	// 目标服务器的URL
	targetURL := fmt.Sprintf("http://%s%s", instance, resource.Url)

	// 创建一个新的请求对象
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	// 复制原始请求的方法、请求URI和正文
	req.SetRequestURI(targetURL)
	req.Header.SetMethodBytes(ctx.Method())
	req.SetBody(ctx.PostBody())

	// 复制原始请求的头部（根据需要可以修改或过滤这些头部）
	//ctx.Request.Header.CopyTo(&req.Header)

	req.Header.Set("Content-Type", "application/json")

	// 创建一个新的响应对象
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	// 使用fasthttp的客户端发送请求
	client := &fasthttp.Client{}
	if err := client.Do(req, resp); err != nil {
		//ctx.Error("failed to forward request: "+err.Error(), fasthttp.StatusInternalServerError)
		fmt.Printf("%+v", err)
		//log.ErrorContextf(ctx, "")
		return
	}

	// 将目标服务器的响应复制回原始客户端的响应
	resp.Header.CopyTo(&ctx.Response.Header)
	ctx.SetBody(resp.Body())
	ctx.SetStatusCode(resp.StatusCode())
}
