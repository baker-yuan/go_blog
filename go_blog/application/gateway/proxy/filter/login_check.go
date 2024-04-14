package filter

import (
	"github.com/baker-yuan/go-blog/go_blog/application/blog/gateway/proxy/biz_ctx"
	"github.com/baker-yuan/go-blog/go_blog/application/blog/gateway/proxy/biz_ctx/http"
	"github.com/baker-yuan/go-blog/go_blog/application/blog/gateway/proxy/config"
	"github.com/baker-yuan/go-blog/go_blog/application/blog/gateway/proxy/util"
	auth_pb "github.com/baker-yuan/go-blog/go_blog/protocol/auth"
)

// LoginCheck 登机校验
func LoginCheck(next HandlerFunc) HandlerFunc {
	return func(bizCtx biz_ctx.IBizContext, cfg *config.Config) {
		httpCtx, _ := bizCtx.(*http.HttpContext)

		resource := bizCtx.Value(biz_ctx.ResourceCtxKey).(*auth_pb.Resource)

		// 无需登录的接口也尝试揭秘ticket
		if resource.GetIsNeedLogin() == auth_pb.NeedLogin_LOGIN_NOT_REQUIRED {
			next(bizCtx, cfg)
			return
		}

		// 登录校验
		accessToken := httpCtx.Request().Header().GetHeader("X-Access-Token")
		token, err := util.JwtUtil.ValidateAndExtractUserInfo(accessToken)
		if err != nil {
			return
		}
		httpCtx.Proxy().Header().SetHeader("X-Access-Token", accessToken)
		httpCtx.Proxy().Header().SetHeader("X-User-ID", token.UserID)

		next(bizCtx, cfg)
	}
}
