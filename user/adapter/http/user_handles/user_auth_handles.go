package user_handles

import (
	"github.com/baker-yuan/go-blog/middleware"
	"github.com/baker-yuan/go-blog/user/application/cqe"
	"github.com/baker-yuan/go-blog/user/application/dto"
	"github.com/baker-yuan/go-blog/user/application/service"
	"github.com/baker-yuan/go-blog/user/domain/repo"
	"github.com/gin-gonic/gin"
)

type UserAuthController struct {
	userAuthService service.UserAuthService
}

func UserAuthRegister(group *gin.RouterGroup, userRepo repo.UserRepo) {
	user := &UserAuthController{
		userAuthService: service.NewUserAuthService(userRepo),
	}
	group.POST("/login", user.Login)

}

// Login godoc
//
//	@Summary		用户名密码登陆
//	@Description	用户名密码登陆
//	@Tags			用户模块
//	@Accept			json
//	@Produce		json
//	@Param			login	body		vo.PasswordLoginVO	true	"用户名密码登陆"
//	@Success		200		{object}	middleware.Response{data=dto.UserDetailDTO}
//	@Router			/ [get]
func (ua *UserAuthController) Login(ginCtx *gin.Context) {
	var (
		param = &cqe.PasswordLoginVO{}
	)
	var (
		userDetail *dto.UserDetailDTO
		err        error
	)
	if err := param.BindValidParam(ginCtx); err != nil {
		middleware.ResponseError(ginCtx, middleware.FAIL)
		return
	}
	userDetail, err = ua.userAuthService.Login(*param)
	middleware.SendResult(ginCtx, userDetail, err)
}
