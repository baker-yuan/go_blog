package main

import (
	"log"
	"os"

	"github.com/baker-yuan/go-blog/application/user/infrastructure/auth"
	"github.com/baker-yuan/go-blog/application/user/infrastructure/persistence"
	"github.com/baker-yuan/go-blog/application/user/interfaces"
	pb "github.com/baker-yuan/go-blog/protocol/user"
	"trpc.group/trpc-go/trpc-go"

	_ "trpc.group/trpc-go/trpc-filter/debuglog"
	_ "trpc.group/trpc-go/trpc-filter/recovery"
	_ "trpc.group/trpc-go/trpc-filter/validation"
	_ "trpc.group/trpc-go/trpc-naming-polarismesh"
)

func init() {
	// todo 临时设置
	trpc.ServerConfigPath = "/Users/yuanyu/code/go-study/go-blog/go_blog/application/blog/cmd/app/trpc_go.yaml"
}

func main() {
	// redis连接配置
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	redisPassword := os.Getenv("REDIS_PASSWORD")

	// db操作
	services, err := persistence.NewRepositories()
	if err != nil {
		panic(err)
	}
	defer services.Close()
	if err := services.AutoMigrate(); err != nil {
		panic(err)
	}

	// redis操作
	redisService, err := auth.NewRedisDB(redisHost, redisPort, redisPassword)
	if err != nil {
		log.Fatal(err)
	}

	tk := auth.NewToken()
	users := interfaces.NewUsers(services.User, redisService.Auth, tk)
	//authenticate := interfaces.NewAuthenticate(services.User, redisService.Auth, tk)

	// trpc服务
	trpcServer := trpc.NewServer()

	// 注册
	pb.RegisterUserApiService(trpcServer, users)

	// 启动服务
	log.Fatal(trpcServer.Serve())
}
