package biz_ctx

import (
	"context"
	"net"
	"time"
)

//// ContextKey 定义一个类型安全的键类型
//type ContextKey[T any] struct {
//	name string
//}
//
//// NewContextKey 创建一个新的类型安全的键
//func NewContextKey[T any](name string) ContextKey[T] {
//	return ContextKey[T]{name: name}
//}
//
//// SetValue 设置类型化的键值对
//func SetValue[T any](ctx context.Context, key ContextKey[T], val T) context.Context {
//	return context.WithValue(ctx, key, val)
//}
//
//// GetValue 获取类型化的键对应的值
//func GetValue[T any](ctx context.Context, key ContextKey[T]) (T, bool) {
//	val, ok := ctx.Value(key).(T)
//	return val, ok
//}

var (
	ResourceCtxKey = "resource" // 资源，类型 *auth_pb.Resource
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

	RealIP() string      // 客户端IP
	LocalIP() net.IP     // 本机IP
	LocalAddr() net.Addr // 服务器监听的本地地址
	LocalPort() int      // 监听端口
}
