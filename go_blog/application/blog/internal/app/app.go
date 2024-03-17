// Package app configures and runs application.
package app

import (
	"github.com/baker-yuan/go-blog/application/blog/config"
	"github.com/baker-yuan/go-blog/application/blog/internal/controller"
	"github.com/baker-yuan/go-blog/application/blog/internal/usecase"
	"github.com/baker-yuan/go-blog/application/blog/internal/usecase/repo"
	pb "github.com/baker-yuan/go-blog/protocol/blog"
	"trpc.group/trpc-go/trpc-database/gorm"
	"trpc.group/trpc-go/trpc-go/log"
	"trpc.group/trpc-go/trpc-go/server"
)

// Run 启动服务
func Run(trpcServer *server.Server, cfg *config.Config) {
	// 数据库连接
	gormDB, err := gorm.NewClientProxy("trpc.mysql.blog.template")
	if err != nil {
		log.Errorf("gorm init fail err: %+v", err)
		panic(err)
	}

	var (
		articleRepo  = repo.NewArticleRepo(gormDB)
		categoryRepo = repo.NewCategoryRepo(gormDB)
	)
	var (
		commonUseCase   = usecase.NewCommonUseCase()
		articleUseCase  = usecase.NewArticleUseCase(commonUseCase, articleRepo)
		categoryUseCase = usecase.NewCategoryUseCase(commonUseCase, categoryRepo)
	)
	var (
		blogService = controller.NewBlogServiceImpl(articleUseCase, categoryUseCase)
	)

	// 注册服务
	pb.RegisterBlogApiService(trpcServer, blogService)

	// 初始化
	if err := repo.Init(gormDB); err != nil {
		panic(err)
	}

	// 启动
	if err := trpcServer.Serve(); err != nil {
		log.Errorf("sever start fail: %+v", err)
	}
}
