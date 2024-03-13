package filter

import (
	"log"

	"github.com/baker-yuan/go-blog/application/blog/gateway/biz_ctx"
)

// ForwardHandler 请求转发处理
func ForwardHandler(bizCtx *biz_ctx.HttpContext) {
	// 这里实现请求转发的逻辑
	log.Println("Forwarding request")
	// ... 请求转发代码 ...
}
