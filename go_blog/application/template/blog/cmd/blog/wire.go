//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	pb "github.com/baker-yuan/go-blog/application/blog/api/blog/v1"
	"github.com/baker-yuan/go-blog/application/blog/internal/biz"
	"github.com/baker-yuan/go-blog/application/blog/internal/data"
	"github.com/baker-yuan/go-blog/application/blog/internal/service"
	"github.com/google/wire"
	trpc_srv "trpc.group/trpc-go/trpc-go/server"
)

type App struct {
	trpcServer *trpc_srv.Server
	greeter    *service.BlogService
}

func newApp(trpcServer *trpc_srv.Server, greeter *service.BlogService) *App {
	return &App{
		trpcServer: trpcServer,
		greeter:    greeter,
	}
}

func (app *App) Run() error {
	pb.RegisterBlogApiService(app.trpcServer, app.greeter)
	return app.trpcServer.Serve()
}

// wireApp 初始化应用
func wireApp(trpcServer *trpc_srv.Server) (*App, error) {
	panic(wire.Build(data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
