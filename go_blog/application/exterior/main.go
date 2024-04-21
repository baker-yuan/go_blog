package main

import (
	"github.com/baker-yuan/go-blog/go_blog/application/exterior/config"
	oss_controller "github.com/baker-yuan/go-blog/go_blog/application/exterior/internal/controller/oss"
	"github.com/baker-yuan/go-blog/go_blog/application/exterior/internal/repository"
	oss_repo "github.com/baker-yuan/go-blog/go_blog/application/exterior/internal/repository/oss"
	oss_service "github.com/baker-yuan/go-blog/go_blog/application/exterior/internal/service/oss"
	oss_pb "github.com/baker-yuan/go-blog/go_blog/protocol/exterior"
	"github.com/baker-yuan/go-blog/go_blog/retcode"
	"trpc.group/trpc-go/trpc-database/gorm"
	"trpc.group/trpc-go/trpc-go"
	"trpc.group/trpc-go/trpc-go/http"
	"trpc.group/trpc-go/trpc-go/log"

	_ "github.com/baker-yuan/go-blog/go_blog/filters/httptrace"
	_ "trpc.group/trpc-go/trpc-filter/debuglog"
	_ "trpc.group/trpc-go/trpc-filter/recovery"
	_ "trpc.group/trpc-go/trpc-filter/validation"
	_ "trpc.group/trpc-go/trpc-naming-polarismesh"
)

func init() {
	// todo 临时设置
	trpc.ServerConfigPath = "/Users/yuanyu/code/go-study/go-blog/go_blog/go_blog/application/exterior/trpc_go.yaml"
}

func main() {
	// trpc服务
	trpcServer := trpc.NewServer()

	// 异常处理，统一返回值
	retcode.InitServerCodec([]string{"/external/callback/aliOss"})

	// 数据库连接
	gormDB, err := gorm.NewClientProxy("trpc.mysql.blog.external")
	if err != nil {
		log.Fatalf("gorm init fail err: %+v", err)
		return
	}

	if err := repository.Init(gormDB); err != nil {
		log.Fatalf("repository init fail: %+v", err)
		return
	}
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("newConfig fail: %+v", err)
		return
	}
	var (
		uploadRecordRepo = oss_repo.NewUploadRecordRepo(gormDB)
		modeCfgRepo      = oss_repo.NewOssModeCfgRepo(gormDB)
	)
	var (
		commonUseCase     = oss_service.NewCommonUseCase(cfg, modeCfgRepo, uploadRecordRepo)
		ossUseCase        = oss_service.NewOssUploadRecordUseCase(commonUseCase, cfg, uploadRecordRepo)
		ossModeCfgUseCase = oss_service.NewOssModeCfgUseCase(commonUseCase, cfg, modeCfgRepo)
	)
	var (
		ossApiService = oss_controller.NewOssApiService(ossUseCase, ossModeCfgUseCase)
	)
	oss_pb.RegisterOssApiService(trpcServer, ossApiService)

	http.HandleFunc("/external/callback/aliOss", ossUseCase.AliOssCallbackHandler)
	http.RegisterNoProtocolService(trpcServer.Service("trpc.blog.http_no_protocol.ExternalApi"))

	// 启动
	if err := trpcServer.Serve(); err != nil {
		log.Errorf("sever start fail: %+v", err)
	}
}
