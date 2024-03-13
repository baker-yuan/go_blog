package persistence

import (
	"context"

	"github.com/baker-yuan/go-blog/application/user/domain/entity"
	"github.com/baker-yuan/go-blog/application/user/domain/repository"
	"github.com/baker-yuan/go-blog/common/db"
	pb "github.com/baker-yuan/go-blog/protocol/user"
	"gorm.io/gorm"
)

// User 用户表
type User struct {
	// 账号信息
	UID      uint32 `gorm:"primary_key;column:uid;type:int unsigned auto_increment;comment:主键"`
	Username string `gorm:"unique;column:username;type:varchar(100);not null;comment:用户名(唯一)"`
	Password string `gorm:"column:password;type:varchar(100);not null;comment:登录密码"`
	Salt     string `gorm:"column:salt;type:varchar(20);not null;default:'';comment:盐"`
	UserType string `gorm:"column:user_type;type:enum('ADMIN','USER');not null;default:'USER';comment:用户类型 ADMIN-管理员 USER-普通用户"`
	// 基本信息
	Email    string `gorm:"column:email;type:varchar(50);not null;default:'';comment:邮箱号"`
	Nickname string `gorm:"column:nickname;type:varchar(30);not null;default:'';comment:昵称"`
	Avatar   string `gorm:"column:avatar;type:varchar(255);not null;default:'';comment:头像地址"`
	Intro    string `gorm:"column:intro;type:varchar(255);not null;default:'';comment:用户简介"`
	WebSite  string `gorm:"column:web_site;type:varchar(255);not null;default:'';comment:个人网站"`
	// 状态
	Status uint8 `gorm:"column:status;type:tinyint unsigned;not null;default:0;comment:状态 0-正常 1-禁用"`
	// 三方登录
	LoginType uint8  `gorm:"column:login_type;type:tinyint unsigned;not null;default:0;comment:登录方式 1-用户名 2-GitHub 3-码云 4-QQ 5-微博"`
	UnionID   string `gorm:"column:union_id;type:varchar(50);not null;default:'';comment:用户唯一标识（第三方网站）"`
	// 公共字段
	IsDeleted    db.BoolBit   `gorm:"column:is_deleted;type:bit(1);not null;default:b'0';comment:是否注销"`
	CreateUserID string       `gorm:"column:create_user_id;type:int(10) unsigned;not null;default:0;comment:创建人"`
	UpdateUserID string       `gorm:"column:update_user_id;type:int(10) unsigned;not null;default:0;comment:修改人"`
	CreateTime   db.Timestamp `gorm:"column:create_time;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:创建时间"`
	UpdateTime   db.Timestamp `gorm:"column:update_time;type:timestamp;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:修改时间"`
}

// TableName 设置 User 结构体对应的数据库表名
func (User) TableName() string {
	return "blog_user"
}

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepo {
	return &UserRepo{db}
}

// UserRepo 强制UserRepo实现repository.UserRepository接口
var _ repository.UserRepository = &UserRepo{}

// GetUserByID 根据用户id查询用户
func (r *UserRepo) GetUserByID(ctx context.Context, id uint32) (*entity.User, error) {

	return nil, nil
}

// GetUserByIDs 根据用户id集合查询用户
func (r *UserRepo) GetUserByIDs(ctx context.Context, ids []uint32) (entity.Users, error) {

	return nil, nil
}

// Save 保存用户
func (r *UserRepo) Save(ctx context.Context, user *entity.User) (uint32, error) {

	return 0, nil
}

// UpdateByID 根据ID修改用户
func (r *UserRepo) UpdateByID(ctx context.Context, user *entity.User) error {

	return nil
}

// SearchUser 用户搜索
func (r *UserRepo) SearchUser(ctx context.Context, req *pb.SearchUserReq) (entity.Users, uint32, error) {

	return nil, 0, nil
}
