package biz_ctx

import "github.com/valyala/fasthttp"

type HttpContext struct {
	ctx *fasthttp.RequestCtx
}

// NewContext 创建Context
func NewContext(ctx *fasthttp.RequestCtx) *HttpContext {
	return &HttpContext{
		ctx: ctx,
	}
}
