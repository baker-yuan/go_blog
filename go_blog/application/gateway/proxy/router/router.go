package router

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/baker-yuan/go-blog/application/blog/gateway/proxy/biz_ctx"
	"github.com/baker-yuan/go-blog/application/blog/gateway/proxy/biz_ctx/http"
	"github.com/baker-yuan/go-blog/application/blog/gateway/proxy/config"
	"github.com/baker-yuan/go-blog/application/blog/gateway/proxy/service"
	auth_pb "github.com/baker-yuan/go-blog/protocol/auth"
	datasync_pb "github.com/baker-yuan/go-blog/protocol/datasync"
	"trpc.group/trpc-go/trpc-go/client"
	"trpc.group/trpc-go/trpc-go/log"
)

var resourceList = make([]*auth_pb.Resource, 0)
var urlResourceMap = make(map[string]*auth_pb.Resource)

func LoadResourceList(ctx context.Context, cfg *config.Config) error {
	// 首次加载
	if err := doLoadResourceList(ctx, cfg); err != nil {
		return err
	}
	// 定时加载
	go func() {
		ticker := time.Tick(10 * time.Minute)
		for range ticker {
			if err := doLoadResourceList(ctx, cfg); err != nil {
				log.ErrorContextf(ctx, "#doLoadResourceList fail err: %+v", err)
			}
			log.DebugContextf(ctx, "#doLoadResourceList success")
		}
	}()
	return nil
}

func doLoadResourceList(ctx context.Context, cfg *config.Config) error {
	// 获取权限服务地址
	instance, err := service.GetOneInstance(cfg.Global.Namespace, "trpc.blog.auth.BlogApiTrpc")
	if err != nil {
		panic(err)
	}
	// 读取接口配置
	proxy := auth_pb.NewAuthApiClientProxy(
		client.WithTarget(fmt.Sprintf("ip://%s", instance)),
		client.WithTimeout(2*time.Second),
	)
	resourceRsp, err := proxy.GetEffectiveResource(ctx, &auth_pb.GetEffectiveResourceReq{})
	if err != nil {
		log.ErrorContextf(ctx, "#httpServerRun getEffectiveResource fail err: %+v", err)
		return err
	}
	// 赋值
	resourceList = resourceRsp.Data
	for _, item := range resourceList {
		urlResourceMap[item.Url] = item
	}
	return nil
}

// MatchResource 路由匹配
func MatchResource(ctx biz_ctx.IBizContext) *auth_pb.Resource {
	httpCtx, _ := ctx.(*http.HttpContext)
	url := httpCtx.Request().URI().RequestURI()
	if res, exist := urlResourceMap[url]; exist {
		return res
	}
	return nil
}

type HandDataChange struct {
}

// 强制HandDataChange实现DataSyncApiService
var _ datasync_pb.DataSyncApiService = (*HandDataChange)(nil)

// DataChange 数据发送变化
func (h *HandDataChange) DataChange(stream datasync_pb.DataSyncApi_DataChangeServer) error {
	for {
		// 循环接受客户端数据
		tableChange, err := stream.Recv()
		// 客户端流已经结束
		if err == io.EOF {
			return stream.SendAndClose(&datasync_pb.DataChangeRsp{})
		}
		// 流发生异常，需要返回
		if err != nil {
			return err
		}

		tableChange = tableChange
	}
}
