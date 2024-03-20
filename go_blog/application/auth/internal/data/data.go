// Package data 业务数据访问，包含 cache、db 等封装，实现了 biz 的 repo 接口。
// 我们可能会把 data 与 dao 混淆在一起，data 偏重业务的含义，它所要做的是将领域对象重新拿出来，我们去掉了 DDD 的 infra层。
package data

import (
	"fmt"

	"github.com/google/wire"
	"gorm.io/gorm"
	tgorm "trpc.group/trpc-go/trpc-database/gorm"
	"trpc.group/trpc-go/trpc-go/log"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewMenuRepo, NewResourceRepo, NewRoleRepo)

// Data 数据操作
type Data struct {
	gormDB *gorm.DB
}

func (d *Data) GetDB() *gorm.DB {
	return d.gormDB
}

// NewData 数据库操作
func NewData() (*Data, error) {
	gormDB, err := tgorm.NewClientProxy("trpc.mysql.blog.auth")
	if err != nil {
		log.Errorf("gorm init fail err: %+v", err)
		panic(err)
	}
	return &Data{
		gormDB: gormDB,
	}, nil
}

func Init(db *gorm.DB) error {
	// 自动建表
	//if err := InitSchema(db); err != nil {
	//	return err
	//}
	// 全字段更新
	if err := InitField(db); err != nil {
		return err
	}
	return nil
}

type RegisterFuncType func(db *gorm.DB) error

var (
	initFieldFuncTypes []RegisterFuncType
)

// 注册全字段更新初始化函数回调
func registerInitField(funcType RegisterFuncType) {
	initFieldFuncTypes = append(initFieldFuncTypes, funcType)
}

// InitField 全字段更新，初始化那些字段不更新，那些字段需要更新
func InitField(db *gorm.DB) error {
	for _, funcType := range initFieldFuncTypes {
		if err := funcType(db); err != nil {
			return err
		}
	}
	return nil
}

// InitSchema 初始化
func InitSchema(db *gorm.DB) error {
	schemas := []tableSchema{
		{
			TableName: "菜单资源表",
			StructPtr: &Menu{},
		},
		{
			TableName: "接口资源表",
			StructPtr: &Resource{},
		},
		{
			TableName: "角色表",
			StructPtr: &Role{},
		},
		{
			TableName: "角色关联的资源和目录表",
			StructPtr: &RoleAuthority{},
		},
		{
			TableName: "账号角色绑定表",
			StructPtr: &UserRole{},
		},
	}

	return autoMigrate(db, schemas)
}

// tableSchema 自动建表描述信息
type tableSchema struct {
	TableName string      // 表名
	StructPtr interface{} // 结构体指针
}

func autoMigrate(db *gorm.DB, schemas []tableSchema) error {
	for _, schema := range schemas {
		if err := db.
			Set("gorm:table_options", fmt.Sprintf("CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT='%s'", schema.TableName)).
			AutoMigrate(schema.StructPtr); err != nil {
			log.Errorf("#autoMigrate fail schema:%+v", schema, err)
			return err
		}
	}
	return nil
}
