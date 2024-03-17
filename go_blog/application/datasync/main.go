package main

import (
	"context"

	"github.com/baker-yuan/go-blog/application/blog/datasync/config"
	"github.com/baker-yuan/go-blog/application/blog/datasync/consumer"
	"github.com/baker-yuan/go-blog/application/blog/datasync/go_mysql"
)

func main() {
	ctx := context.Background()

	// 创建配置
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	// 初始化consumer包
	consumer.Init(cfg.Consumer.BufferSize)

	// 初始化go_mysql包
	if err := go_mysql.Init(ctx, cfg); err != nil {
		panic(err)
	}
}
