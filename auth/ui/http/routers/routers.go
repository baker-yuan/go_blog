package routers

import (
	middleware "github.com/baker-yuan/go-blog/all_packaged_library/middleware/gin"
	"github.com/baker-yuan/go-blog/auth/ui/http/menu_handles"
	"github.com/gin-gonic/gin"
)

func SetRouters(r *gin.Engine, menuSrv menu_handles.MenuUI) {
	middleware.SetCorsRouters(r)
	middleware.SetRecovery(r)
	r.GET("/admin/user/menus", menuSrv.ListUserMenus)
}
