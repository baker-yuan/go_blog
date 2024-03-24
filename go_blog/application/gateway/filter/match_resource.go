package filter

import (
	"github.com/baker-yuan/go-blog/application/blog/gateway/biz_ctx"
	"github.com/baker-yuan/go-blog/application/blog/gateway/config"
	"github.com/baker-yuan/go-blog/application/blog/gateway/router"
)

// MatchResource 匹配资源
func MatchResource(next HandlerFunc) HandlerFunc {
	return func(bizCtx biz_ctx.IBizContext, cfg *config.Config) {
		resource := router.MatchResource(bizCtx)
		if resource == nil {
			return
		}
		bizCtx.WithValue(biz_ctx.ResourceCtxKey, resource)
		next(bizCtx, cfg)
	}
}
