package filter

import (
	"log"

	"github.com/baker-yuan/go-blog/application/blog/gateway/biz_ctx"
	"github.com/baker-yuan/go-blog/application/blog/gateway/config"
)

func Recovery(next HandlerFunc) HandlerFunc {
	return func(bizCtx biz_ctx.IBizContext, cfg *config.Config) {
		log.Println("Recovery before")
		next(bizCtx, cfg)
		log.Println("Recovery after")
	}
}
