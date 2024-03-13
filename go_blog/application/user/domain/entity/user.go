package entity

import (
	"github.com/baker-yuan/go-blog/common/db"
	pb "github.com/baker-yuan/go-blog/protocol/user"
)

// User 用户
type User struct {
	// 账号信息
	UID      uint32 // 主键
	Username string // 用户名(唯一)
	Password string // 登录密码
	Salt     string // 盐
	UserType string // 用户类型 ADMIN-管理员 USER-普通用户
	// 基本信息
	Email    string // 邮箱号
	Nickname string // 昵称
	Avatar   string // 头像地址
	Intro    string // 用户简介
	WebSite  string // 个人网站
	// 状态
	Status pb.UserStatus // 状态 0-正常 1-禁用
	// 三方登录
	LoginType pb.LoginType // 登录方式 1-用户名 2-GitHub 3-码云 4-QQ 5-微博
	UnionID   string       // 用户唯一标识（第三方网站）
	// 公共字段
	IsDeleted    db.BoolBit // 是否注销
	CreateUserID uint32     // 创建人
	UpdateUserID uint32     // 修改人
	CreateTime   uint32     // 创建时间
	UpdateTime   uint32     // 修改时间
}

func (u *User) Delete() {
	u.IsDeleted = true
}

type Users []*User
