package server

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

// Server 服务
type Server struct {
	endpoint      string            // 端口
	registerGRPC  registerFunc      // 注册grpc
	registerHTTP  registerFunc      // 注册http
	HTTPListener  net.Listener      //
	GRPCListener  net.Listener      //
	httpServer    *http.Server      //
	router        *http.ServeMux    //
	GRPClientConn *grpc.ClientConn  //
	ServerMux     *runtime.ServeMux //
	tcpMux        cmux.CMux         //
}

type registerFunc func(ctx context.Context, s *Server)

type Option func(*Server)

func WithEndpoint(endpoint string) Option {
	return func(s *Server) { s.endpoint = endpoint }
}

func WithHTTPRegisterFunc(registerHTTP registerFunc) Option {
	return func(s *Server) { s.registerHTTP = registerHTTP }
}

func WithGRPCRegisterFunc(registerGRPC registerFunc) Option {
	return func(s *Server) { s.registerGRPC = registerGRPC }
}

// New returns a Server instance
func New(opts ...Option) *Server {
	s := &Server{}

	for _, o := range opts {
		o(s)
	}

	listener, err := net.Listen("tcp", s.endpoint)
	if err != nil {
		panic(err)
	}
	s.tcpMux = cmux.New(listener)

	return s
}

// Start 启动服务
func (s *Server) Start() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	s.GRPCListener = s.tcpMux.MatchWithWriters(cmux.HTTP2MatchHeaderFieldPrefixSendSettings("content-type", "application/grpc"))
	s.HTTPListener = s.tcpMux.Match(cmux.HTTP1Fast())

	go func() {
		s.registerGRPC(ctx, s)
	}()

	go func() {
		if err := s.initGateway(ctx); err != nil {
			panic(err)
		}
		s.registerHTTP(ctx, s)
		s.startGateway()
	}()

	return s.tcpMux.Serve()
}

func (s *Server) Stop() {
	s.tcpMux.Close()
}

func (s *Server) initGateway(ctx context.Context) error {
	var err error

	s.router = http.NewServeMux()

	s.GRPClientConn, err = grpc.Dial(s.endpoint, []grpc.DialOption{
		grpc.WithTimeout(10 * time.Second),
		grpc.WithBlock(),
		grpc.WithInsecure(),
	}...)
	if err != nil {
		return fmt.Errorf("fail to dial: %v", err)
	}

	s.ServerMux = runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				UseProtoNames:  true,
				UseEnumNumbers: true,
			},
		}),
		runtime.WithErrorHandler(runtime.DefaultHTTPErrorHandler),
	)

	return nil
}

func (s *Server) startGateway() {
	s.router.Handle("/", s.ServerMux)

	s.httpServer = &http.Server{
		Addr:         s.endpoint,
		Handler:      s.router,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
		IdleTimeout:  time.Second,
	}

	if err := s.httpServer.Serve(s.HTTPListener); err != nil {
		panic(err)
	}
}
