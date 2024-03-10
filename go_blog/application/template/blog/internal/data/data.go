// Package data 业务数据访问，包含 cache、db 等封装，实现了 biz 的 repo 接口。
// 我们可能会把 data 与 dao 混淆在一起，data 偏重业务的含义，它所要做的是将领域对象重新拿出来，我们去掉了 DDD 的 infra层。
package data

import (
	"github.com/google/wire"
	"gorm.io/gorm"
	tgorm "trpc.group/trpc-go/trpc-database/gorm"
	"trpc.group/trpc-go/trpc-go/log"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewArticleRepo)

type Data struct {
	gormDB *gorm.DB
}

// NewData 数据库操作
func NewData() (*Data, error) {
	gormDB, err := tgorm.NewClientProxy("trpc.mysql.blog.template")
	if err != nil {
		log.Errorf("gorm init fail err: %+v", err)
		panic(err)
	}
	return &Data{
		gormDB: gormDB,
	}, nil
}
