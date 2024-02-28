package follow_handles

import (
	"context"

	"github.com/baker-yuan/go-blog/all_packaged_library/base/log"
	middleware "github.com/baker-yuan/go-blog/all_packaged_library/middleware/gin"
	follow3 "github.com/baker-yuan/go-blog/interaction/application/service/follow"
	"github.com/baker-yuan/go-blog/interaction/ui/http/assembler/follow"
	follow2 "github.com/baker-yuan/go-blog/interaction/ui/http/dto/follow"
	"github.com/baker-yuan/go-blog/interaction/ui/port"
	"github.com/gin-gonic/gin"
)

// FollowService 关注
type FollowService interface {
	// AddFollow 新增关注关系
	AddFollow(g *gin.Context)
}

type FollowServiceImpl struct {
	AppService  *follow3.AppService
	MetricsPort port.MetricsPort
}

func (f *FollowServiceImpl) AddFollow(g *gin.Context) {
	var (
		req = &follow2.AddFollowReq{}
		err error
		ctx = context.Background()
	)
	// gin参数绑定
	if err = g.BindJSON(req); err != nil {
		log.Error(ctx, "bind addFollow req fail err: %v", err)
		middleware.ResponseError(g)
		return
	}

	// 转换数据: 得到 addFollowCMD
	addFollowCMD := follow.GenAddFollowCMD(req)

	// 上报监控
	f.MetricsPort.CounterIncr("gin.AddFollow")

	// 调用 AppService
	err = f.AppService.AddFollow(ctx, addFollowCMD)
	if err != nil {
		// 记录日志
		log.Error(ctx, "addFollow fail err: %v", err)
		// 上报监控
		f.MetricsPort.CounterIncr("tRPC.AddFollow.ERR")
		middleware.ResponseError(g)
		return
	}
	middleware.ResponseSuccess(g)
}
