package main

import (
	"log"

	"github.com/baker-yuan/go-blog/application/user/application"
	"github.com/baker-yuan/go-blog/application/user/infrastructure/auth"
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
	// trpc服务
	trpcServer := trpc.NewServer()

	// db操作
	repository, err := persistence.NewRepositories()
	if err != nil {
		panic(err)
	}
	defer repository.Close()
	if err := repository.AutoMigrate(); err != nil {
		panic(err)
	}

	// 应用服务
	userApp := application.NewUserApp(repository.User)

	// redis工具
	redisService, err := auth.NewRedisDB()
	if err != nil {
		log.Fatal(err)
	}
	// jwt工具
	tk := auth.NewToken()

	users := interfaces.NewUsers(userApp, redisService.Auth, tk)
	authenticate := interfaces.NewAuthenticate(userApp, redisService.Auth, tk)

	// 注册
	pb.RegisterUserApiService(trpcServer, users)
	pb.RegisterLoginApiService(trpcServer, authenticate)

	// 启动服务
	log.Fatal(trpcServer.Serve())
}
