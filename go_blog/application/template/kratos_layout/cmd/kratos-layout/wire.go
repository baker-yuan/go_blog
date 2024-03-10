//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"kratos-layout/internal/biz"
	"kratos-layout/internal/data"
	"kratos-layout/internal/server"
	"kratos-layout/internal/service"

	"github.com/google/wire"
	trpc_srv "trpc.group/trpc-go/trpc-go/server"
)

type App struct {
	trpcServer *trpc_srv.Server
}

func newApp(trpcServer *trpc_srv.Server) *App {
	return &App{
		trpcServer: trpcServer,
	}
}

func (app *App) Run() error {
	//gormDB, _ := gorm.NewClientProxy("trpc.mysql.blog.template")
	//if err != nil {
	//	log.Errorf("gorm init fail err: %+v", err)
	//	panic(err)
	//}
	//gormDB = gormDB
	return app.trpcServer.Serve()
}

// wireApp 初始化应用
func wireApp() (*App, error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
