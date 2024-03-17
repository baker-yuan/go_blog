// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/baker-yuan/go-blog/application/auth/internal/biz"
	"github.com/baker-yuan/go-blog/application/auth/internal/data"
	"github.com/baker-yuan/go-blog/application/auth/internal/service"
	"github.com/baker-yuan/go-blog/protocol/auth"
	"trpc.group/trpc-go/trpc-go/server"
)

import (
	_ "trpc.group/trpc-go/trpc-database/gorm"
	_ "trpc.group/trpc-go/trpc-filter/debuglog"
	_ "trpc.group/trpc-go/trpc-filter/recovery"
	_ "trpc.group/trpc-go/trpc-filter/validation"
	_ "trpc.group/trpc-go/trpc-naming-polarismesh"
)

// Injectors from wire.go:

// wireApp 初始化应用
func wireApp(trpcServer *server.Server) (*App, error) {
	commonUseCase := biz.NewCommonUseCase()
	dataData, err := data.NewData()
	if err != nil {
		return nil, err
	}
	menuRepo := data.NewMenuRepo(dataData)
	menuUsecase := biz.NewMenuUsecase(commonUseCase, menuRepo)
	resourceRepo := data.NewResourceRepo(dataData)
	resourceUseCase := biz.NewResourceUseCase(commonUseCase, resourceRepo)
	authService := service.NewAuthService(menuUsecase, resourceUseCase)
	app := newApp(trpcServer, authService, dataData)
	return app, nil
}

// wire.go:

type App struct {
	trpcServer  *server.Server
	serviceImpl *service.AuthService
	data        *data.Data
}

func newApp(trpcServer *server.Server, serviceImpl *service.AuthService, data2 *data.Data) *App {
	return &App{
		trpcServer:  trpcServer,
		serviceImpl: serviceImpl,
		data:        data2,
	}
}

func (app *App) Run() error {
	auth.RegisterAuthApiService(app.trpcServer, app.serviceImpl)
	data.Init(app.data.GetDB())
	return app.trpcServer.Serve()
}
