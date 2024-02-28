package main

import (
	"context"

	"github.com/baker-yuan/go-blog/all_packaged_library/base"
	"github.com/baker-yuan/go-blog/all_packaged_library/base/log"
	"github.com/baker-yuan/go-blog/all_packaged_library/util"
	follow_srv "github.com/baker-yuan/go-blog/interaction/application/service/follow"
	praise_srv "github.com/baker-yuan/go-blog/interaction/application/service/praise"
	follow_adapter "github.com/baker-yuan/go-blog/interaction/infrastructure/adapter/follow"
	praise_adapter "github.com/baker-yuan/go-blog/interaction/infrastructure/adapter/praise"
	"github.com/baker-yuan/go-blog/interaction/ui/http"
	"github.com/baker-yuan/go-blog/interaction/ui/http/follow_handles"
	"github.com/baker-yuan/go-blog/interaction/ui/http/praise_handles"
)

func initPraiseService() praise_handles.PraiseService {
	// 初始化 application 服务
	appService := &praise_srv.AppService{
		// 依赖注入
		PraisePort: &praise_adapter.PraiseAdapter{},
	}
	// 注册 http 服务
	httpService := &praise_handles.PraiseServiceImpl{
		AppService:  appService,                       // 应用服务
		MetricsPort: &praise_adapter.MetricsAdapter{}, // 数据上报
	}
	return httpService
}
func initFollowService() follow_handles.FollowService {
	// 初始化 application 服务
	appService := &follow_srv.AppService{
		// 依赖注入
		FollowPort: &follow_adapter.FollowAdapter{},
	}
	// 注册 http 服务
	httpService := &follow_handles.FollowServiceImpl{
		AppService:  appService,                       // 应用服务
		MetricsPort: &follow_adapter.MetricsAdapter{}, // 数据上报
	}
	return httpService
}

func main() {
	base.Init()
	http.NewHttp(initPraiseService(), initFollowService())
	util.QuitSignal(func() {
		log.Info(context.Background(), "server exit")
	})
}
