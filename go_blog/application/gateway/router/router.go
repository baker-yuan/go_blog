package router

import (
	"context"
	"fmt"
	"time"

	"github.com/baker-yuan/go-blog/application/blog/gateway/biz_ctx"
	"github.com/baker-yuan/go-blog/application/blog/gateway/config"
	"github.com/baker-yuan/go-blog/application/blog/gateway/service"
	auth_pb "github.com/baker-yuan/go-blog/protocol/auth"
	"trpc.group/trpc-go/trpc-go/client"
	"trpc.group/trpc-go/trpc-go/log"
)

var resourceList = make([]*auth_pb.Resource, 0)
var urlResourceMap = make(map[string]*auth_pb.Resource)

func LoadResourceList(ctx context.Context, cfg *config.Config) error {
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
func MatchResource(ctx biz_ctx.BizContext) *auth_pb.Resource {
	//httpContext, err := biz_ctx.Assert(ctx)
	//if err != nil {
	//	return nil
	//}

	httpCtx, _ := ctx.(*biz_ctx.HttpContext)

	url := string(httpCtx.FastCtx().Request.RequestURI())
	if res, exist := urlResourceMap[url]; exist {
		return res
	}
	return nil
}
