package main

import (
	"context"
	"log"

	"github.com/baker-yuan/go-blog/application/user/application"
	"github.com/baker-yuan/go-blog/application/user/infrastructure/config"
	"github.com/baker-yuan/go-blog/application/user/infrastructure/persistence"
	"github.com/baker-yuan/go-blog/application/user/interfaces"
	pb "github.com/baker-yuan/go-blog/protocol/user"
	"trpc.group/trpc-go/trpc-go"

	_ "trpc.group/trpc-go/trpc-database/gorm"
	_ "trpc.group/trpc-go/trpc-filter/debuglog"
	_ "trpc.group/trpc-go/trpc-filter/recovery"
	_ "trpc.group/trpc-go/trpc-filter/validation"
	_ "trpc.group/trpc-go/trpc-naming-polarismesh"
)

func init() {
	// todo 临时设置
	trpc.ServerConfigPath = "/Users/yuanyu/code/go-study/go-blog/go_blog/go_blog/application/user/trpc_go.yaml"
}

func main() {
	ctx := context.Background()

	// todo
	configPath := "/Users/yuanyu/code/go-study/go-blog/go_blog/go_blog/application/user/config.yaml"

	// trpc服务
	trpcServer := trpc.NewServer()

	// 加载应用程序配置。
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		panic(err)
	}

	// db操作
	repository, err := persistence.NewRepositories()
	if err != nil {
		panic(err)
	}
	defer repository.Close()
	if err := repository.AutoMigrate(); err != nil {
		panic(err)
	}

	jwtUtil, err := cfg.JWT.Build(ctx)
	if err != nil {
		panic(err)
	}

	// 应用服务
	userApp := application.NewUserApp(repository.User, jwtUtil)

	// 接口实现
	userApiImpl := interfaces.NewUsers(userApp)
	loginApiImpl := interfaces.NewAuthenticate(userApp)

	// 注册
	pb.RegisterUserApiService(trpcServer, userApiImpl)
	pb.RegisterLoginApiService(trpcServer, loginApiImpl)

	// 启动服务
	log.Fatal(trpcServer.Serve())
}
