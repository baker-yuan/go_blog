package util

import (
	"github.com/baker-yuan/go-blog/common/util/conversion"
	"github.com/baker-yuan/go-blog/common/util/obj"
	"github.com/baker-yuan/go-blog/common/util/page"
	"github.com/baker-yuan/go-blog/common/util/session"
	"github.com/baker-yuan/go-blog/common/util/slice"
	"github.com/baker-yuan/go-blog/common/util/stack"
	"github.com/baker-yuan/go-blog/common/util/str"
)

var (
	TypeConversionUtils conversion.ConversionUtils // 类型转换，用于string，int，int64，float等数据转换，免去err的接收，和设置默认值
	StackUtils          stack.StackUtils           // 堆栈处理
	StrUtils            str.StrUtils               // 字符串处理
	SessionUtils        session.SessionUtils       // 会话
)

// NewSliceUtils 是 SliceUtils 的构造函数，返回一个 SliceUtils 的实例。
func NewSliceUtils[T comparable]() slice.SliceUtils[T] {
	return slice.SliceUtils[T]{}
}

// NewFieldExtractor 创建并返回一个新的 FieldExtractor 实例。
func NewFieldExtractor[T any, K comparable]() slice.FieldExtractor[T, K] {
	return slice.FieldExtractor[T, K]{}
}

// NewPageUtils 内存分页
func NewPageUtils[T any]() page.PageUtils[T] {
	return page.PageUtils[T]{}
}

// NewObjectUtils 对象操作
func NewObjectUtils[T any]() obj.ObjectUtils[T] {
	return obj.ObjectUtils[T]{}
}
