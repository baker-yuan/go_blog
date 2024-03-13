package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/baker-yuan/go-blog/application/blog/gateway/http_proxy"
)

func main() {
	// 加载配置

	// 启动http服务
	go func() {
		http_proxy.HttpServerRun()
	}()

	// 监听关闭信息
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// 停止服务
	http_proxy.HttpServerStop()
}
