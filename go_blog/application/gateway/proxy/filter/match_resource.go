package filter

import (
	"github.com/baker-yuan/go-blog/go_blog/application/blog/gateway/proxy/biz_ctx"
	"github.com/baker-yuan/go-blog/go_blog/application/blog/gateway/proxy/biz_ctx/http"
	"github.com/baker-yuan/go-blog/go_blog/application/blog/gateway/proxy/config"
	"github.com/baker-yuan/go-blog/go_blog/application/blog/gateway/proxy/ret"
	"github.com/baker-yuan/go-blog/go_blog/application/blog/gateway/proxy/router"
)

// MatchResource 匹配资源
func MatchResource(next HandlerFunc) HandlerFunc {
	return func(bizCtx biz_ctx.IBizContext, cfg *config.Config) {
		httpCtx, _ := bizCtx.(*http.HttpContext)
		resource := router.MatchResource(bizCtx)
		if resource == nil {
			ret.SetFilResponse(httpCtx, "match resource fail")
			return
		}
		bizCtx.WithValue(biz_ctx.ResourceCtxKey, resource)
		next(bizCtx, cfg)
	}
}
