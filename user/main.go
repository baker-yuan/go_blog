package main

import (
	"context"
	"fmt"

	"github.com/baker-yuan/go-blog/all_packaged_library/base"
	"github.com/baker-yuan/go-blog/all_packaged_library/base/config"
	"github.com/baker-yuan/go-blog/all_packaged_library/grpc/server"
	pb "github.com/baker-yuan/go-blog/proto/user"
	"github.com/baker-yuan/go-blog/user/application/service"
	"github.com/baker-yuan/go-blog/user/infrastructure/adapter"
	applicationservice "github.com/baker-yuan/go-blog/user/ui/grpc"
	"google.golang.org/grpc"
)

func registerHTTP(ctx context.Context, s *server.Server) {
	if err := pb.RegisterUserHandler(ctx, s.ServerMux, s.GRPClientConn); err != nil {
		panic(err)
	}
}

func registerGRPC(ctx context.Context, server *server.Server) {

	// 初始化 application 服务
	appService := &service.AppService{
		// 依赖注入
		UserPort: &adapter.UserAdapter{},
	}

	// 注册 gRPC 服务
	gRPCService := &applicationservice.UserServerImpl{
		AppService:  appService,
		MetricsPort: &adapter.MetricsAdapter{},
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServer(grpcServer, gRPCService)
	if err := grpcServer.Serve(server.GRPCListener); err != nil {
		panic(err)
	}
}

func main() {
	base.Init()
	endpoint := fmt.Sprintf(":%d", config.GetHttpConf().Addr)

	grpcServer := server.New(
		server.WithEndpoint(endpoint),
		server.WithGRPCRegisterFunc(registerGRPC),
		server.WithHTTPRegisterFunc(registerHTTP),
	)

	grpcServer.Start()
}
