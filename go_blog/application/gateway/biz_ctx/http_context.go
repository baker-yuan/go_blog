package biz_ctx

import (
	"context"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/baker-yuan/go-blog/application/blog/gateway/config"
	"github.com/baker-yuan/go-blog/common/util"
	"github.com/google/uuid"
	"github.com/valyala/fasthttp"
)

// 强制HttpContext实现BizContext
var _ BizContext = (*HttpContext)(nil)

type HttpContext struct {
	fastHttpRequestCtx *fasthttp.RequestCtx // fastHttp
	ctx                context.Context      //
	requestID          string               // 请求id
	labels             map[string]string    // 标签
	port               int                  // 监听端口
}

func (ctx *HttpContext) FastCtx() *fasthttp.RequestCtx {
	return ctx.fastHttpRequestCtx
}

func (ctx *HttpContext) Context() context.Context {
	if ctx.ctx == nil {
		ctx.ctx = context.Background()
	}
	return ctx.ctx
}

func (ctx *HttpContext) Value(key interface{}) interface{} {
	return ctx.Context().Value(key)
}

func (ctx *HttpContext) WithValue(key, val interface{}) {
	ctx.ctx = context.WithValue(ctx.Context(), key, val)
}

func (ctx *HttpContext) Scheme() string {
	return string(ctx.fastHttpRequestCtx.Request.URI().Scheme())
}

// RequestId 请求ID
func (ctx *HttpContext) RequestId() string {
	return ctx.requestID
}

// AcceptTime 请求接收时间
func (ctx *HttpContext) AcceptTime() time.Time {
	return ctx.fastHttpRequestCtx.Time()
}

func (ctx *HttpContext) Assert(i interface{}) error {
	if v, ok := i.(*BizContext); ok {
		*v = ctx
		return nil
	}
	return fmt.Errorf("not suport:%s", util.ReflectUtils.TypeNameOf(i))
}

// SetLabel 设置标签
func (ctx *HttpContext) SetLabel(name, value string) {
	ctx.labels[name] = value
}

// GetLabel 获取标签
func (ctx *HttpContext) GetLabel(name string) string {
	return ctx.labels[name]
}

// Labels 返回所有标签
func (ctx *HttpContext) Labels() map[string]string {
	return ctx.labels
}

// LocalIP 本机IP
func (ctx *HttpContext) LocalIP() net.IP {
	return ctx.fastHttpRequestCtx.LocalIP()
}

// LocalAddr 服务器监听的本地地址
func (ctx *HttpContext) LocalAddr() net.Addr {
	return ctx.fastHttpRequestCtx.LocalAddr()
}

// LocalPort 监听端口
func (ctx *HttpContext) LocalPort() int {
	return ctx.port
}

// Assert EoContext是否是IHttpContext
func Assert(ctx BizContext) (HttpContext, error) {
	var httpContext HttpContext
	err := ctx.Assert(&httpContext)
	return httpContext, err
}

// NewContext 创建Context
func NewContext(ctx *fasthttp.RequestCtx, cfg *config.Config) BizContext {
	return &HttpContext{
		fastHttpRequestCtx: ctx,
		requestID:          uuid.New().String(),
		port:               util.TypeConversionUtils.StrToInt(strings.Split(cfg.Http.Addr, ":")[1]),
	}
}
