package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/baker-yuan/go-blog/application/blog/gateway/config"
	"github.com/baker-yuan/go-blog/application/blog/gateway/http_proxy"
	"github.com/baker-yuan/go-blog/application/blog/gateway/router"
	"github.com/baker-yuan/go-blog/application/blog/gateway/service"
	"github.com/baker-yuan/go-blog/application/blog/gateway/util"
)

func main() {
	ctx := context.Background()

	// 加载配置
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	// jwt工具
	jwtUtil, err := cfg.Jwt.Build(ctx)
	if err != nil {
		panic(err)
	}

	util.Init(jwtUtil)

	// 服务模块初始化
	if err := service.Init(); err != nil {
		panic(err)
	}

	// 初始化路由信息
	if err := router.LoadResourceList(ctx, cfg); err != nil {
		panic(err)
	}

	// 启动http服务
	go func() {
		http_proxy.HttpServerRun(ctx, cfg)
	}()

	// 监听关闭信息
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// 停止服务
	http_proxy.HttpServerStop()
}
