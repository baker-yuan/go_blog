package cqe

import (
	"github.com/baker-yuan/go-blog/public"
	"github.com/gin-gonic/gin"
)

type PasswordLoginVO struct {
	Username *string `json:"username" form:"username"` // 用户名
	Password *string `json:"password" form:"username"` // 密码
}

// BindValidParam 参数绑定
func (param *PasswordLoginVO) BindValidParam(c *gin.Context) error {
	// 参数绑定 && 参数校验
	return public.DefaultGetValidParams(c, param)
}
