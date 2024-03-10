package persistence

import (
	"github.com/baker-yuan/go-blog/application/user/domain/entity"
	"github.com/baker-yuan/go-blog/application/user/domain/repository"
	"gorm.io/gorm"
	tgorm "trpc.group/trpc-go/trpc-database/gorm"
	"trpc.group/trpc-go/trpc-go/log"

	_ "github.com/jinzhu/gorm/dialects/postgres"
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
	return s.db.AutoMigrate(&entity.User{})
}
