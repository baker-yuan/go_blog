// Package app configures and runs application.
package app

import (
	pb "github.com/baker-yuan/go-blog/protocol/template"
	"github.com/baker-yuan/go-blog/template/go_clean_template/config"
	"github.com/baker-yuan/go-blog/template/go_clean_template/internal/controller"
	"github.com/baker-yuan/go-blog/template/go_clean_template/internal/usecase"
	"github.com/baker-yuan/go-blog/template/go_clean_template/internal/usecase/repo"
	"trpc.group/trpc-go/trpc-database/gorm"
	"trpc.group/trpc-go/trpc-go/log"
	"trpc.group/trpc-go/trpc-go/server"
)

// Run 启动服务
func Run(trpcServer *server.Server, cfg *config.Config) {

	gormDB, err := gorm.NewClientProxy("trpc.mysql.blog.template")
	if err != nil {
		log.Errorf("gorm init fail err: %+v", err)
		panic(err)
	}

	var (
		friendLinkRepo = repo.NewFriendLinkRepo(gormDB)
	)
	var (
		commonUseCase     = usecase.NewCommonUseCase()
		friendLinkUseCase = usecase.NewFriendLinkUseCase(commonUseCase, friendLinkRepo)
	)

	templateService := controller.NewTemplateServiceImpl(friendLinkUseCase)

	// 注册服务
	pb.RegisterTemplateApiService(trpcServer, templateService)

	// 启动
	if err := trpcServer.Serve(); err != nil {
		log.Errorf("sever start fail: %+v", err)
	}
}
