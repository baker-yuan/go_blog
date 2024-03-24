package filter

import (
	"github.com/baker-yuan/go-blog/application/blog/gateway/biz_ctx"
	"github.com/baker-yuan/go-blog/application/blog/gateway/config"
	auth_pb "github.com/baker-yuan/go-blog/protocol/auth"
)

// LoginCheck 登机校验
func LoginCheck(next HandlerFunc) HandlerFunc {
	return func(bizCtx biz_ctx.IBizContext, cfg *config.Config) {
		resource := bizCtx.Value(biz_ctx.ResourceCtxKey).(*auth_pb.Resource)

		// 无需登录的接口也尝试揭秘ticket
		if resource.GetIsNeedLogin() == auth_pb.NeedLogin_LOGIN_NOT_REQUIRED {
			next(bizCtx, cfg)
			return
		}

		// 登录校验
		//util.JwtUtil.ValidateAndExtractUserInfo()
		next(bizCtx, cfg)
	}
}
