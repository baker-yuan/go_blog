package persistence

import (
	"github.com/baker-yuan/go-blog/go_blog/application/user/domain/repository"
	"gorm.io/gorm"
	tgorm "trpc.group/trpc-go/trpc-database/gorm"
	"trpc.group/trpc-go/trpc-go/log"
)

type Repositories struct {
	User repository.UserRepository
	db   *gorm.DB
}

func NewRepositories() (*Repositories, error) {
	db, err := tgorm.NewClientProxy("trpc.mysql.blog.user")
	if err != nil {
		log.Errorf("gorm init fail err: %+v", err)
		panic(err)
	}

	return &Repositories{
		User: NewUserRepository(db),
		db:   db,
	}, nil
}

// Close 关闭数据库连接
func (s *Repositories) Close() error {
	return nil
}

// AutoMigrate 自动建表
func (s *Repositories) AutoMigrate() error {
	return s.db.AutoMigrate(&User{})
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
func (s *Repositories) InitField() error {
	for _, funcType := range initFieldFuncTypes {
		if err := funcType(s.db); err != nil {
			return err
		}
	}
	return nil
}
