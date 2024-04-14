package persistence

import (
	"context"
	"errors"

	"github.com/baker-yuan/go-blog/go_blog/application/user/domain/entity"
	"github.com/baker-yuan/go-blog/go_blog/application/user/domain/repository"
	"github.com/baker-yuan/go-blog/go_blog/application/user/infrastructure/security"
	"github.com/baker-yuan/go-blog/go_blog/common/db"
	"github.com/baker-yuan/go-blog/go_blog/common/util"
	pb "github.com/baker-yuan/go-blog/go_blog/protocol/user"
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
	Status pb.UserStatus `gorm:"column:status;type:tinyint unsigned;not null;default:0;comment:状态 0-正常 1-禁用"`
	// 三方登录
	LoginType pb.LoginType `gorm:"column:login_type;type:tinyint unsigned;not null;default:0;comment:登录方式 1-用户名 2-GitHub 3-码云 4-QQ 5-微博"`
	UnionID   string       `gorm:"column:union_id;type:varchar(50);not null;default:'';comment:用户唯一标识（第三方网站）"`
	// 逻辑删除
	IsDeleted db.BoolBit `gorm:"column:is_deleted;type:bit(1);not null;default:b'0';comment:是否注销"`
	// 公共字段
	CreateUserID uint32       `gorm:"column:create_user_id;type:int unsigned;not null;default:0;comment:创建人id"`
	UpdateUserID uint32       `gorm:"column:update_user_id;type:int unsigned;not null;default:0;comment:更新人id"`
	CreateTime   db.Timestamp `gorm:"column:create_time;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:插入时间"`
	UpdateTime   db.Timestamp `gorm:"column:update_time;type:timestamp;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:修改时间"`
}

var (
	UserTbName        = "blog_user"
	UserFieldUid      = "uid"
	UserFieldUsername = "username"
	UserFieldPassword = "password"
	UserFieldSalt     = "salt"
	UserFieldUserType = "user_type"
	// 基本信息
	UserFieldEmail    = "email"
	UserFieldNickname = "nickname"
	UserFieldAvatar   = "avatar"
	UserFieldIntro    = "intro"
	UserFieldWebSite  = "web_site"
	// 状态
	UserFieldStatus = "status"
	// 逻辑删除
	UserFieldIsDeleted = "is_deleted"
	// 三方登录
	UserFieldLoginType = "login_type"
	UserFieldUnionID   = "union_id"
	// 公共字段
	UserFieldCreateTime = "create_time"
	UserFieldUpdateTime = "update_time"

	UserTypeAdmin = "ADMIN"
	UserTypeUser  = "USER"

	SaltValueLength = 20 // 盐长度
)

// TableName 设置 User 结构体对应的数据库表名
func (*User) TableName() string {
	return UserTbName
}

type UserRepo struct {
	*db.GenericDao[User, uint32]
}

// UserRepo 强制UserRepo实现repository.UserRepository接口
var _ repository.UserRepository = &UserRepo{}

func NewUserRepository(gormDB *gorm.DB) *UserRepo {
	return &UserRepo{
		GenericDao: db.NewGenericDao[User, uint32](gormDB),
	}
}

func init() {
	registerInitField(initUserField)
}

var (
	// 全字段修改User那些字段不修改
	notUpdateUserField = []string{
		UserFieldUserType,

		UserFieldLoginType,
		UserFieldUnionID,

		UserFieldIsDeleted,
		UserFieldCreateTime,
	}
	updateUserField []string
)

// InitUserField 全字段修改
func initUserField(db *gorm.DB) error {
	columnTypes, err := db.Migrator().ColumnTypes(&User{})
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
	return UserPoToEntity(user), err
}

// GetUserByIDs 根据用户id集合查询用户
func (r *UserRepo) GetUserByIDs(ctx context.Context, ids ...uint32) (entity.Users, error) {
	dbUsers, err := r.GenericDao.GetByIDs(ctx, ids)
	if err != nil {
		return nil, err
	}
	return UserPosToEntity(dbUsers), nil
}

// DeleteByID 根据ID删除用户
func (r *UserRepo) DeleteByID(ctx context.Context, id uint32) error {
	return r.GenericDao.DeleteByID(ctx, id)
}

// SaveAdmin 保存用户
func (r *UserRepo) SaveAdmin(ctx context.Context, user *entity.User) (uint32, error) {
	if user.UID > 0 {
		return 0, errors.New("illegal argument user id exist")
	}
	dbUser := UserEntityToPo(user)
	// 用户类型固定是管理员
	dbUser.UserType = UserTypeAdmin
	// 设置密码+盐
	if err := r.setPasswordAndSalt(ctx, dbUser); err != nil {
		return 0, err
	}
	// 保存
	if err := r.GenericDao.Create(ctx, dbUser); err != nil {
		return 0, err
	}
	return dbUser.UID, nil
}

// UpdateAdminByID 根据ID修改用户
func (r *UserRepo) UpdateAdminByID(ctx context.Context, user *entity.User) error {
	if user.UID == 0 {
		return errors.New("illegal argument user exist")
	}
	dbUser := UserEntityToPo(user)
	//if dbUser.UserType != UserTypeAdmin {
	//	return errors.New("userType not support")
	//}
	// 设置密码+盐
	if err := r.setPasswordAndSalt(ctx, dbUser); err != nil {
		return err
	}
	return r.GenericDao.GetDB(ctx).
		Select(updateUserField).
		Where(UserFieldUid, user.UID).
		Updates(dbUser).
		Error
}

// setPasswordAndSalt 设置密码+盐
func (r *UserRepo) setPasswordAndSalt(ctx context.Context, user *User) error {
	// 密码加密存储
	if len(user.Salt) == 0 {
		user.Salt = util.StrUtils.RandomString(SaltValueLength)
	}
	password, err := security.GeneratePassword(user.Salt)
	if err != nil {
		return err
	}
	user.Password = password
	return nil
}

// SearchUser 用户搜索
func (r *UserRepo) SearchUser(ctx context.Context, req *pb.SearchUserReq) (entity.Users, uint32, error) {
	var (
		res       []*User
		pageTotal int64
	)
	tx, err := db.BuildSearch(
		ctx,
		req.GetSearch(),
		r.GenericDao.GetDB(ctx),
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
	return UserPOsToEntity(res), uint32(pageTotal), nil
}

// GetUserByEmailAndPassword 通过用户名查找用户，并且验证密码
func (r *UserRepo) GetUserByEmailAndPassword(ctx context.Context, username string, password string) (*entity.User, error) {
	// 查询用户
	tx := r.GenericDao.GetDB(ctx).
		Where(UserFieldUsername, username)
	data := &User{}
	if err := tx.Find(&data).Error; err != nil {
		return nil, err
	}
	// 校验密码
	if err := security.VerifyPassword(data.Password, password); err != nil {
		return nil, err
	}
	return UserPoToEntity(data), nil
}
