package filter

import (
	"log"

	"github.com/baker-yuan/go-blog/application/blog/gateway/biz_ctx"
)

// RequestLog 日志记录
func RequestLog(bizCtx *biz_ctx.HttpContext, next func()) {
	log.Println("Middleware1 before")
	next()
	log.Println("Middleware1 after")
}
