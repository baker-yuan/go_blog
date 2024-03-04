package main

import (
	"log"

	"github.com/baker-yuan/go-blog/template/config"
	"github.com/baker-yuan/go-blog/template/internal/app"
	"trpc.group/trpc-go/trpc-go"

	_ "trpc.group/trpc-go/trpc-filter/debuglog"
	_ "trpc.group/trpc-go/trpc-filter/recovery"
	_ "trpc.group/trpc-go/trpc-filter/validation"
	_ "trpc.group/trpc-go/trpc-naming-polarismesh"
)

func init() {
	// todo 临时设置
	trpc.ServerConfigPath = "/Users/yuanyu/code/go-study/go-blog/go_blog/application/template/cmd/app/trpc_go.yaml"
}

// https://github.com/trpc-group
// https://github.com/trpc-ecosystem
// https://github.com/trpc-ecosystem/go-filter/blob/main/README.zh_CN.md
// https://github.com/trpc-ecosystem/go-database/blob/main/gorm/README.zh_CN.md
func main() {
	// trpc服务
	trpcServer := trpc.NewServer()

	// 创建配置
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(trpcServer, cfg)
}
