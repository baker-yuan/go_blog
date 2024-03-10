package server

import (
	"kratos-layout/internal/service"

	"trpc.group/trpc-go/trpc-go"
	"trpc.group/trpc-go/trpc-go/server"
)

// NewTRPCServer new a tRPC server.
func NewTRPCServer(greeter *service.GreeterService) *server.Server {
	trpcServer := trpc.NewServer()

	//v1.RegisterGreeterServer(srv, greeter)
	return trpcServer
}
