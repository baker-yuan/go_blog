// Package app configures and runs application.
package app

import (
	"github.com/baker-yuan/go-blog/go_blog/application/blog/config"
	"github.com/baker-yuan/go-blog/go_blog/application/blog/internal/controller"
	"github.com/baker-yuan/go-blog/go_blog/application/blog/internal/usecase"
	"github.com/baker-yuan/go-blog/go_blog/application/blog/internal/usecase/repo"
	pb "github.com/baker-yuan/go-blog/go_blog/protocol/blog"
	"github.com/baker-yuan/go-blog/go_blog/retcode"
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
		articleRepo       = repo.NewArticleRepo(gormDB)
		categoryRepo      = repo.NewCategoryRepo(gormDB)
		tagRepo           = repo.NewTagRepo(gormDB)
		friendLinkRepo    = repo.NewFriendLinkRepo(gormDB)
		pageRepo          = repo.NewPageRepo(gormDB)
		dynamicConfigRepo = repo.NewDynamicConfigRepo(gormDB)
	)
	var (
		commonUseCase        = usecase.NewCommonUseCase()
		articleUseCase       = usecase.NewArticleUseCase(commonUseCase, articleRepo)
		categoryUseCase      = usecase.NewCategoryUseCase(commonUseCase, categoryRepo)
		tagUseCase           = usecase.NewTagUseCase(commonUseCase, tagRepo)
		friendLinkUseCase    = usecase.NewFriendLinkUseCase(commonUseCase, friendLinkRepo)
		pageUseCase          = usecase.NewPageUseCase(commonUseCase, pageRepo)
		dynamicConfigUseCase = usecase.NewDynamicConfigUseCase(commonUseCase, dynamicConfigRepo)
		portalUseCase        = usecase.NewPortalUseCase(commonUseCase)
	)
	var (
		blogService = controller.NewBlogServiceImpl(
			articleUseCase,
			categoryUseCase,
			tagUseCase,
			friendLinkUseCase,
			pageUseCase,
			dynamicConfigUseCase,
			portalUseCase,
		)
	)

	// 注册服务
	pb.RegisterBlogApiService(trpcServer, blogService)
	pb.RegisterBlogPortalApiService(trpcServer, blogService)

	// 初始化
	if err := repo.Init(gormDB); err != nil {
		panic(err)
	}

	// 异常处理，统一返回值
	retcode.InitServerCodec(nil)

	// 启动
	if err := trpcServer.Serve(); err != nil {
		log.Errorf("sever start fail: %+v", err)
	}
}
