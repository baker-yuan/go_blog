package repo

import (
	"fmt"

	"github.com/baker-yuan/go-blog/application/blog/internal/entity"
	"gorm.io/gorm"
	"trpc.group/trpc-go/trpc-go/log"
)

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
			TableName: "文章表",
			StructPtr: &entity.Article{},
		},
		{
			TableName: "文章标签表",
			StructPtr: &entity.Tag{},
		},
		{
			TableName: "文章标签关联表",
			StructPtr: &entity.ArticleTag{},
		},
		{
			TableName: "文章分类表",
			StructPtr: &entity.Category{},
		},
		{
			TableName: "页面表",
			StructPtr: &entity.Page{},
		},
		{
			TableName: "动态表",
			StructPtr: &entity.Moment{},
		},
		{
			TableName: "友链表",
			StructPtr: &entity.FriendLink{},
		},
		{
			TableName: "网站配置",
			StructPtr: &entity.Config{},
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
