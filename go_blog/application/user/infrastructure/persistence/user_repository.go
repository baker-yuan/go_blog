package persistence

import (
	"context"
	"errors"

	"github.com/baker-yuan/go-blog/application/user/domain/entity"
	"github.com/baker-yuan/go-blog/application/user/domain/repository"
	"github.com/baker-yuan/go-blog/application/user/infrastructure/assembler"
	"github.com/baker-yuan/go-blog/common/db"
	"github.com/baker-yuan/go-blog/common/util"
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
	*db.GenericDao[User, uint32]
}

// UserRepo 强制UserRepo实现repository.UserRepository接口
var _ repository.UserRepository = &UserRepo{}

func NewUserRepository(gormDB *gorm.DB) *UserRepo {
	return &UserRepo{
		GenericDao: &db.GenericDao[User, uint32]{
			DB: gormDB,
		},
	}
}

func init() {
	registerInitField(initUserField)
}

var (
	// 全字段修改User那些字段不修改
	notUpdateUserField = []string{
		"created_at",
	}
	updateUserField []string
)

// InitUserField 全字段修改
func initUserField(db *gorm.DB) error {
	columnTypes, err := db.Migrator().ColumnTypes(&entity.User{})
	if err != nil {
		return err
	}
	columns := make([]string, 0)
	for _, v := range columnTypes {
		columns = append(columns, v.Name())
	}
	updateUserField = util.NewSliceUtils[string]().SliceRemove(columns, notUpdateUserField...)
	return nil
}

// GetUserByID 根据用户id查询用户
func (r *UserRepo) GetUserByID(ctx context.Context, id uint32) (*entity.User, error) {
	user, err := r.GenericDao.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return assembler.UserPoToEntity(user), err
}

// GetUserByIDs 根据用户id集合查询用户
func (r *UserRepo) GetUserByIDs(ctx context.Context, ids []uint32) (entity.Users, error) {
	dbUsers, err := r.GenericDao.GetByIDs(ctx, ids)
	if err != nil {
		return nil, err
	}
	return assembler.UserPosToEntity(dbUsers), nil
}

// DeleteByID 根据ID删除用户
func (r *UserRepo) DeleteByID(ctx context.Context, id uint32) error {
	return r.GenericDao.DeleteByID(ctx, id)
}

// Save 保存用户
func (r *UserRepo) Save(ctx context.Context, user *entity.User) (uint32, error) {
	if user.UID > 0 {
		return 0, errors.New("illegal argument user id exist")
	}
	dbUser := assembler.UserEntityToPo(user)
	if err := r.GenericDao.Create(ctx, dbUser); err != nil {
		return 0, err
	}
	return dbUser.UID, nil
}

// UpdateByID 根据ID修改用户
func (r *UserRepo) UpdateByID(ctx context.Context, user *entity.User) error {
	if user.UID == 0 {
		return errors.New("illegal argument user exist")
	}
	dbUser := assembler.UserEntityToPo(user)
	return r.GenericDao.DB.WithContext(ctx).Select(updateUserField).Updates(dbUser).Error
}

// SearchUser 用户搜索
func (r *UserRepo) SearchUser(ctx context.Context, req *pb.SearchUserReq) (entity.Users, uint32, error) {
	var (
		res       []*entity.User
		pageTotal int64
	)
	tx, err := db.BuildSearch(
		ctx,
		req.GetSearch(),
		r.GenericDao.DB.WithContext(ctx),
		func(search map[string]*db.SearchValue) {

		},
	)
	if err != nil {
		return nil, 0, err
	}
	tx = tx.Offset(int((req.GetPageNum() - 1) * req.GetPageSize())).
		Limit(int(req.GetPageSize())).Find(&res).
		Offset(-1).Limit(-1).Count(&pageTotal)
	if err := tx.Error; err != nil {
		return nil, 0, err
	}
	return res, uint32(pageTotal), nil
}
