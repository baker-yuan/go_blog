package http

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/valyala/fasthttp"
)

// headerActionHandleFunc http header 暂存/重放
// https://github.com/eolinker/apinto/pull/113
// https://github.com/eolinker/apinto/pull/117
//
// 1、http 转发用了一个快捷方式来复制上游返回的response，就会导致在http协议下，转发完成后，在转发前对响应的header设置被覆盖，
// 针对这种情况，我们将转发前的header操作添加到actions中，在转发完成后执行refresh进行重放。
// 2、其他协议转发(grpc/double)和http转发不一样(不是执行SendTo转发到下游)，不会复制上游返回的response，也不会refresh(SendTo才会)，action也不会进行重放，
// 直接设置的header不存在覆盖的情况。
type headerAction struct {
	Action headerActionHandleFunc
	Key    string
	Value  string
}

type headerActionHandleFunc func(target *ResponseHeader, key string, value ...string)

var (
	headerActionAdd = func(target *ResponseHeader, key string, value ...string) {
		target.cache.Add(key, value[0])
		target.header.Add(key, value[0])
	}
	headerActionSet = func(target *ResponseHeader, key string, value ...string) {
		target.cache.Set(key, value[0])
		target.header.Set(key, value[0])
	}
	headerActionDel = func(target *ResponseHeader, key string, value ...string) {
		target.cache.Del(key)
		target.header.Del(key)
	}
)

type ResponseHeader struct {
	header     *fasthttp.ResponseHeader //
	cache      http.Header              // map[string][]string
	actions    []*headerAction          //
	afterProxy bool                     // NewContext -> false refresh->true
}

func (r *ResponseHeader) reset(header *fasthttp.ResponseHeader) {
	r.header = header
	r.cache = http.Header{}
	r.actions = nil
	r.afterProxy = false
}

func (r *ResponseHeader) GetHeader(name string) string {
	return r.Headers().Get(name)
}

func (r *ResponseHeader) Headers() http.Header {
	return r.cache
}

func (r *ResponseHeader) HeadersString() string {
	return r.header.String()
}

func (r *ResponseHeader) SetHeader(key, value string) {
	r.cache.Set(key, value)
	r.header.Set(key, value)
	if !r.afterProxy {
		r.actions = append(r.actions, &headerAction{
			Key:    key,
			Value:  value,
			Action: headerActionSet,
		})
	}

}

func (r *ResponseHeader) AddHeader(key, value string) {
	r.cache.Add(key, value)
	r.header.Add(key, value)
	if !r.afterProxy {
		r.actions = append(r.actions, &headerAction{
			Key:    key,
			Value:  value,
			Action: headerActionAdd,
		})
	}
}

func (r *ResponseHeader) DelHeader(key string) {
	r.cache.Del(key)
	r.header.Del(key)
	if !r.afterProxy {
		r.actions = append(r.actions, &headerAction{
			Key:    key,
			Action: headerActionDel,
		})
	}
}

// refresh 刷新
func (r *ResponseHeader) refresh() {
	tmp := make(http.Header)
	hs := strings.Split(r.header.String(), "\r\n")
	for _, t := range hs {
		if strings.TrimSpace(t) == "" {
			continue
		}
		vs := strings.Split(t, ":")
		if len(vs) < 2 {
			if vs[0] == "" {
				continue
			}
			tmp[vs[0]] = []string{""}
			continue
		}
		tmp[vs[0]] = []string{strings.TrimSpace(vs[1])}
	}
	r.cache = tmp
	for _, ac := range r.actions {
		ac.Action(r, ac.Key, ac.Value)
	}
	r.afterProxy = true
	r.actions = nil
}

func (r *ResponseHeader) Finish() {
	r.header = nil
	r.cache = nil
	r.actions = nil
}

type Response struct {
	*fasthttp.Response // 客户端和网关之间的响应
	ResponseHeader
	length          int
	responseTime    time.Duration
	proxyStatusCode int
	responseError   error
}

func (r *Response) ResponseError() error {
	return r.responseError
}

func (r *Response) ClearError() {
	r.responseError = nil
}

func (r *Response) SetResponseTime(t time.Duration) {
	r.responseTime = t
}

func (r *Response) ResponseTime() time.Duration {
	return r.responseTime
}

func (r *Response) ContentLength() int {
	if r.length == 0 {
		return r.Response.Header.ContentLength()
	}
	return r.length
}

func (r *Response) ContentType() string {
	return string(r.Response.Header.ContentType())
}

func (r *Response) String() string {
	return r.Response.String()
}

func (r *Response) SetBody(bytes []byte) {
	if strings.Contains(r.GetHeader("Content-Encoding"), "gzip") {
		r.DelHeader("Content-Encoding")
	}
	r.Response.SetBody(bytes)
	r.length = len(bytes)
	r.SetHeader("Content-Length", strconv.Itoa(r.length))
	r.responseError = nil
}

func (r *Response) GetBody() []byte {
	if strings.Contains(r.GetHeader("Content-Encoding"), "gzip") {
		body, _ := r.BodyGunzip()
		r.DelHeader("Content-Encoding")
		r.SetHeader("Content-Length", strconv.Itoa(len(body)))
		r.Response.SetBody(body)
	}
	return r.Response.Body()
}

func (r *Response) BodyLen() int {
	return r.header.ContentLength()
}

func (r *Response) SetStatus(code int, status string) {
	r.Response.SetStatusCode(code)
	r.responseError = nil
}

func (r *Response) SetProxyStatus(code int, status string) {
	r.proxyStatusCode = code
}

func (r *Response) StatusCode() int {
	if r.responseError != nil {
		return 504
	}
	return r.Response.StatusCode()
}

func (r *Response) Status() string {
	return strconv.Itoa(r.StatusCode())
}

func (r *Response) ProxyStatusCode() int {
	return r.proxyStatusCode
}

func (r *Response) ProxyStatus() string {
	return strconv.Itoa(r.proxyStatusCode)
}

func (r *Response) Finish() error {
	r.ResponseHeader.Finish()
	r.Response = nil
	r.responseError = nil
	r.proxyStatusCode = 0
	return nil
}
func (r *Response) reset(resp *fasthttp.Response) {
	r.Response = resp
	r.ResponseHeader.reset(&resp.Header)
	r.responseError = nil
	r.proxyStatusCode = 0
}
