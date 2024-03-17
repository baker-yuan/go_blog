package biz_ctx

import (
	"context"
	"net"
	"time"
)

// IBizContext 上下文，不同的协议实现自己的上下文
type IBizContext interface {
	Context() context.Context          // 原始context
	Value(key interface{}) interface{} // 从原始context中返回键对应的值
	WithValue(key, val interface{})    // 往原始context添加键值对

	Scheme() string // 协议 http、https、grpc、dubbo

	RequestId() string     // 请求id唯一，每次请求随机生成
	AcceptTime() time.Time // 请求接收时间

	Assert(i interface{}) error // context类型断言

	SetLabel(name, value string) // 设置标签
	GetLabel(name string) string // 获取标签
	Labels() map[string]string   // 返回所有标签

	LocalIP() net.IP     // 本机IP
	LocalAddr() net.Addr // 服务器监听的本地地址
	LocalPort() int      // 监听端口
}
