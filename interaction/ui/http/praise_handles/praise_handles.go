package praise_handles

import (
	"context"

	"github.com/baker-yuan/go-blog/all_packaged_library/base/log"
	middleware "github.com/baker-yuan/go-blog/all_packaged_library/middleware/gin"
	service "github.com/baker-yuan/go-blog/interaction/application/service/praise"
	"github.com/baker-yuan/go-blog/interaction/ui/http/assembler/praise"
	praise2 "github.com/baker-yuan/go-blog/interaction/ui/http/dto/praise"
	"github.com/baker-yuan/go-blog/interaction/ui/port"
	"github.com/gin-gonic/gin"
)

// PraiseService 点赞
type PraiseService interface {
	// ObjectPraise 点赞
	ObjectPraise(g *gin.Context)
	// CancelObjectPraise 取消点赞
	CancelObjectPraise(g *gin.Context)
}

type PraiseServiceImpl struct {
	AppService  *service.AppService
	MetricsPort port.MetricsPort
}

// ObjectPraise 点赞
func (f *PraiseServiceImpl) ObjectPraise(g *gin.Context) {
	var (
		req = &praise2.ObjectPraiseReq{}
		err error
		ctx = context.Background()
	)
	// gin参数绑定
	if err = g.BindJSON(req); err != nil {
		log.Error(ctx, "ui - bind addPraise req fail err: %v", err)
		middleware.ResponseError(g)
		return
	}

	// 转换数据: 得到 addFollowCMD
	addPraiseCMD := praise.GenObjectPraiseCMD(req)

	// 上报监控
	f.MetricsPort.CounterIncr("gin.ObjectPraise")

	// 调用 AppService
	err = f.AppService.ObjectPraise(ctx, addPraiseCMD)
	if err != nil {
		// 记录日志
		log.Error(ctx, "addFollow fail err: %v", err)
		// 上报监控
		f.MetricsPort.CounterIncr("gin.ObjectPraise.ERR")
		middleware.ResponseError(g)
		return
	}
	middleware.ResponseSuccess(g)
}

// CancelObjectPraise 取消点赞
func (f *PraiseServiceImpl) CancelObjectPraise(g *gin.Context) {
	var (
		req = &praise2.CancelObjectPraiseReq{}
		err error
		ctx = context.Background()
	)
	// gin参数绑定
	if err = g.BindJSON(req); err != nil {
		log.Error(ctx, "ui - bind addPraise req fail err: %v", err)
		middleware.ResponseError(g)
		return
	}

	// 转换数据: 得到 cancelCMD
	cancelCMD := praise.GenCancelObjectPraiseCMD(req)

	// 上报监控
	f.MetricsPort.CounterIncr("gin.CancelObjectPraise")

	// 调用 AppService
	err = f.AppService.CancelObjectPraise(ctx, cancelCMD)
	if err != nil {
		// 记录日志
		log.Error(ctx, "ui - service - CancelObjectPraise err: %v", err)
		// 上报监控
		f.MetricsPort.CounterIncr("gin.CancelObjectPraise.ERR")
		middleware.ResponseError(g)
		return
	}
	middleware.ResponseSuccess(g)
}
