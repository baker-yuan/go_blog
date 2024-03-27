package filter

import (
	"log"

	"github.com/baker-yuan/go-blog/application/blog/gateway/proxy/biz_ctx"
	"github.com/baker-yuan/go-blog/application/blog/gateway/proxy/config"
)

// RequestLog 日志记录
func RequestLog(next HandlerFunc) HandlerFunc {
	return func(bizCtx biz_ctx.IBizContext, cfg *config.Config) {
		log.Println("RequestLog before")
		next(bizCtx, cfg)
		log.Println("RequestLog after")
	}
}
