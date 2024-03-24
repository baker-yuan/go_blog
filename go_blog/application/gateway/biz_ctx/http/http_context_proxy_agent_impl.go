package http

import (
	"strconv"
	"time"
)

type UrlAgent struct {
	IURIWriter
	host   string
	scheme string
}

func NewUrlAgent(IURIWriter IURIWriter, host string, scheme string) *UrlAgent {
	return &UrlAgent{IURIWriter: IURIWriter, host: host, scheme: scheme}
}

func (u *UrlAgent) SetScheme(scheme string) {
	u.scheme = scheme
}
func (u *UrlAgent) Scheme() string {
	return u.scheme
}

func (u *UrlAgent) Host() string {
	return u.host
}

func (u *UrlAgent) SetHost(host string) {
	u.host = host
}

type requestAgent struct {
	IRequest                 //
	host           string    //
	scheme         string    // http协议
	statusCode     int       //
	status         string    //
	responseLength int       //
	beginTime      time.Time // 请求执行时间
	endTime        time.Time // 请求结束时间
	hostAgent      *UrlAgent //
}

var _ IProxy = (*requestAgent)(nil)

func newRequestAgent(IRequest IRequest, host string, scheme string, beginTime, endTime time.Time) *requestAgent {
	return &requestAgent{
		IRequest:  IRequest,
		host:      host,
		scheme:    scheme,
		beginTime: beginTime,
		endTime:   endTime,
	}
}

func (a *requestAgent) StatusCode() int {
	return a.statusCode
}

func (a *requestAgent) Status() string {
	return a.status
}

func (a *requestAgent) ProxyTime() time.Time {
	return a.beginTime
}

func (a *requestAgent) ResponseLength() int {
	return a.responseLength
}

func (a *requestAgent) ResponseTime() int64 {
	return a.endTime.Sub(a.beginTime).Milliseconds()
}

func (a *requestAgent) URI() IURIWriter {
	if a.hostAgent == nil {
		a.hostAgent = NewUrlAgent(a.IRequest.URI(), a.host, a.scheme)
	}
	return a.hostAgent
}

func (a *requestAgent) setStatusCode(code int) {
	a.statusCode = code
	a.status = strconv.Itoa(code)
}

func (a *requestAgent) setResponseLength(length int) {
	if length > 0 {
		a.responseLength = length
	}
}
