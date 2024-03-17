// Package main 整个项目启动的入口文件
package main

import (
	"trpc.group/trpc-go/trpc-go"

	_ "trpc.group/trpc-go/trpc-database/gorm"
	_ "trpc.group/trpc-go/trpc-filter/debuglog"
	_ "trpc.group/trpc-go/trpc-filter/recovery"
	_ "trpc.group/trpc-go/trpc-filter/validation"
	_ "trpc.group/trpc-go/trpc-naming-polarismesh"
)

func init() {
	// todo 临时设置
	trpc.ServerConfigPath = "/Users/yuanyu/code/go-study/go-blog/go_blog/go_blog/application/auth/cmd/trpc_go.yaml"
}

func main() {
	trpcServer := trpc.NewServer()

	app, err := wireApp(trpcServer)
	if err != nil {
		panic(err)
	}

	// 启动服务
	if err := app.Run(); err != nil {
		panic(err)
	}
}
