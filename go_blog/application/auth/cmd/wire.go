//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/baker-yuan/go-blog/application/auth/internal/biz"
	"github.com/baker-yuan/go-blog/application/auth/internal/data"
	"github.com/baker-yuan/go-blog/application/auth/internal/service"
	pb "github.com/baker-yuan/go-blog/protocol/auth"
	"github.com/google/wire"
	trpc_srv "trpc.group/trpc-go/trpc-go/server"
)

type App struct {
	trpcServer  *trpc_srv.Server
	serviceImpl *service.AuthService
	data        *data.Data
}

func newApp(trpcServer *trpc_srv.Server, serviceImpl *service.AuthService, data *data.Data) *App {
	return &App{
		trpcServer:  trpcServer,
		serviceImpl: serviceImpl,
		data:        data,
	}
}

func (app *App) Run() error {
	pb.RegisterAuthApiService(app.trpcServer, app.serviceImpl)
	data.Init(app.data.GetDB())
	return app.trpcServer.Serve()
}

// wireApp 初始化应用
func wireApp(trpcServer *trpc_srv.Server) (*App, error) {
	panic(wire.Build(data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
