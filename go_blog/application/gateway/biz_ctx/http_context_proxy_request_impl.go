package biz_ctx

import (
	"bytes"
	"fmt"

	"github.com/valyala/fasthttp"
	"trpc.group/trpc-go/trpc-go/log"
)

// ProxyRequest实现了IRequest
var _ IRequest = (*ProxyRequest)(nil)

// ProxyRequest 组装转发的request，这里直接复用了RequestReader，然后添加了自己特有的方法
type ProxyRequest struct {
	RequestReader // 请求读取
}

func (r *ProxyRequest) Header() IHeaderWriter {
	// RequestHeader也实现了IHeaderWriter
	return &r.headers
}

func (r *ProxyRequest) Body() IBodyDataWriter {
	// BodyRequestHandler也实现了IBodyDataWriter
	return &r.body
}

func (r *ProxyRequest) URI() IURIWriter {
	// URIRequest也实现了IURIWriter
	return &r.uri
}

var (
	xforwardedforKey = []byte("x-forwarded-for")
)

func (r *ProxyRequest) reset(request *fasthttp.Request, remoteAddr string) {
	r.req = request
	// 从 request 的 Header 中获取 "x-forwarded-for" 的值。
	// "x-forwarded-for" 是一个 HTTP 头字段，用于识别经过 HTTP 代理或负载均衡器连接的客户端最初的 IP 地址。
	forwardedFor := r.req.Header.PeekBytes(xforwardedforKey)
	if len(forwardedFor) > 0 {
		// 如果 "x-forwarded-for" 的值不为空，那么检查这个值是否含有逗号（,）。如果含有逗号，那么取出逗号之前的部分作为 realIP。
		// 如果不含有逗号，那么直接将整个 "x-forwarded-for" 的值作为 realIP。然后在 "x-forwarded-for" 的值后追加当前的远程地址 remoteAddr。
		if i := bytes.IndexByte(forwardedFor, ','); i > 0 {
			r.realIP = string(forwardedFor[:i])
		} else {
			r.realIP = string(forwardedFor)
		}
		r.req.Header.Set("x-forwarded-for", fmt.Sprint(string(forwardedFor), ",", r.remoteAddr))
	} else {
		// 如果 "x-forwarded-for" 的值为空，那么直接将远程地址 remoteAddr 设置为 "x-forwarded-for" 的值，并将 remoteAddr 作为 realIP。
		r.req.Header.Set("x-forwarded-for", r.remoteAddr)
		r.realIP = r.remoteAddr
	}
	r.RequestReader.reset(r.req, remoteAddr)
}

func (r *ProxyRequest) SetMethod(s string) {
	r.Request().Header.SetMethod(s)
}

func (r *ProxyRequest) Finish() error {
	err := r.RequestReader.Finish()
	if err != nil {
		log.Warn(err)
	}
	return nil
}
