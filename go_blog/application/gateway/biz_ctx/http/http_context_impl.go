package http

import (
	"context"
	"fmt"
	"net"
	"strings"
	"sync"
	"time"

	"github.com/baker-yuan/go-blog/application/blog/gateway/biz_ctx"
	"github.com/baker-yuan/go-blog/application/blog/gateway/config"
	"github.com/baker-yuan/go-blog/application/blog/gateway/fasthttp_client"
	"github.com/baker-yuan/go-blog/common/util"
	"github.com/google/uuid"
	"github.com/valyala/fasthttp"
)

var (
	pool = sync.Pool{
		New: newContext,
	}
	//copyPool = sync.Pool{
	//	New: newCopyContext,
	//}
)

func newContext() interface{} {
	h := new(HttpContext)
	h.proxyRequests = make([]IProxy, 0, 5)
	return h
}

//func newCopyContext() interface{} {
//	h := new(cloneContext)
//	//h.proxyRequests = make([]http_service.IProxy, 0, 5)
//	return h
//}

// 强制HttpContext实现BizContext
var _ biz_ctx.IBizContext = (*HttpContext)(nil)

// 强制HttpContext实现IHttpContext
var _ IHttpContext = (*HttpContext)(nil)

type HttpContext struct {
	fastHttpRequestCtx *fasthttp.RequestCtx // fastHttp
	ctx                context.Context      // ctx
	requestID          string               // 请求id
	labels             map[string]string    // 标签
	port               int                  // 监听端口

	requestReader RequestReader // 请求读取
	response      Response      // 设置响应
	proxyRequest  ProxyRequest  // 组装转发的request
	proxyRequests []IProxy      // 记录转发
}

// NewContext 创建Context
func NewContext(ctx *fasthttp.RequestCtx, cfg *config.Config) biz_ctx.IBizContext {
	// 监听端口
	port := util.TypeConversionUtils.StrToInt(strings.Split(cfg.Http.Addr, ":")[1])
	// 客户端地址
	remoteAddr := ctx.RemoteAddr().String()
	// 从对象池中获取HttpContext
	httpContext := pool.Get().(*HttpContext)
	httpContext.fastHttpRequestCtx = ctx
	httpContext.requestID = uuid.New().String()
	// 原始请求最大读取body为8k，使用clone request
	request := fasthttp.AcquireRequest()
	ctx.Request.CopyTo(request)
	httpContext.requestReader.reset(request, remoteAddr)
	// proxyRequest保留原始请求
	httpContext.proxyRequest.reset(&ctx.Request, remoteAddr)
	httpContext.proxyRequests = httpContext.proxyRequests[:0]
	httpContext.response.reset(&ctx.Response)
	httpContext.labels = make(map[string]string)
	httpContext.port = port
	// 记录请求时间
	httpContext.ctx = context.Background()
	httpContext.WithValue("request_time", ctx.Time())
	return httpContext
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

// Scheme 协议 http、https、grpc、dubbo
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
	if v, ok := i.(*IHttpContext); ok {
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

// RealIP 客户端IP
func (ctx *HttpContext) RealIP() string {
	return ctx.Request().RealIP()
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

func (ctx *HttpContext) Request() IRequestReader {
	return &ctx.requestReader
}

func (ctx *HttpContext) Response() IResponse {
	return &ctx.response
}

func (ctx *HttpContext) Proxy() IRequest {
	return &ctx.proxyRequest
}

func (ctx *HttpContext) Proxies() []IProxy {
	return ctx.proxyRequests
}

// SendTo 发送http请求到下游服务
func (ctx *HttpContext) SendTo(scheme string, node biz_ctx.IInstance, timeout time.Duration) error {
	host := node.Addr()
	request := ctx.proxyRequest.Request() // 这里的请求是从proxyRequest拿的

	beginTime := time.Now()

	// 发送请求，并且吧响应直接塞到fasthttp context里面
	// 1、填充ctx.fastHttpRequestCtx.Response
	// 2、设置responseError
	ctx.response.responseError = fasthttp_client.ProxyTimeout(scheme, node, request, &ctx.fastHttpRequestCtx.Response, timeout)

	agent := newRequestAgent(&ctx.proxyRequest, host, scheme, beginTime, time.Now())

	if ctx.response.responseError != nil {
		// 设置agent响应结果
		agent.setStatusCode(504)
	} else {
		// 前面直接把下游的响应塞到fasthttp context里面，这里需要重放用户手动设置的请求头
		ctx.response.ResponseHeader.refresh()
		// 设置agent响应结果
		agent.setStatusCode(ctx.fastHttpRequestCtx.Response.StatusCode())
	}

	agent.setResponseLength(ctx.fastHttpRequestCtx.Response.Header.ContentLength())

	ctx.proxyRequests = append(ctx.proxyRequests, agent)
	return ctx.response.responseError
}

func (ctx *HttpContext) FastFinish() {

}

// Assert EoContext是否是IHttpContext
func Assert(ctx biz_ctx.IBizContext) (IHttpContext, error) {
	var httpContext IHttpContext
	err := ctx.Assert(&httpContext)
	return httpContext, err
}
