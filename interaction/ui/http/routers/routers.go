package routers

import (
	middleware "github.com/baker-yuan/go-blog/all_packaged_library/middleware/gin"
	"github.com/baker-yuan/go-blog/interaction/ui/http/follow_handles"
	"github.com/baker-yuan/go-blog/interaction/ui/http/praise_handles"
	"github.com/gin-gonic/gin"
)

func SetRouters(r *gin.Engine, httpService praise_handles.PraiseService, followSrv follow_handles.FollowService) {
	middleware.SetCorsRouters(r)
	middleware.SetRecovery(r)
	//
	r.POST("/api/follow/addFollow", followSrv.AddFollow)
	//
	r.POST("/api/praise", httpService.ObjectPraise)
	r.DELETE("/api/praise", httpService.CancelObjectPraise)
}
