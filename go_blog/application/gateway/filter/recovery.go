package filter

import (
	"log"

	"github.com/baker-yuan/go-blog/application/blog/gateway/biz_ctx"
)

// Recovery 捕获所有panic，并且返回错误信息
func Recovery(bizCtx *biz_ctx.HttpContext, next func()) {
	log.Println("Middleware1 before")
	next()
	log.Println("Middleware1 after")
}
