package entity

import (
	"context"
	"fmt"

	"github.com/baker-yuan/go-blog/all_packaged_library/base/log"
)

// User 用户实体
type User struct {
	// 用户信息
	uid       uint32 // 主键
	email     string // 邮箱号
	nickname  string // 用户昵称
	avatar    string // 用户头像
	intro     string // 用户简介
	webSite   string // 个人网站
	isDisable uint32 // 是否禁言
	// 用户账号
	username      string // 用户名
	password      string // 密码
	loginType     uint32 // 登录类型
	ipAddress     string // 用户登录ip
	ipSource      string // ip来源
	lastLoginTime uint32 // 最近登录时间
}

// UID 获取 UID
func (u *User) UID() uint32 {
	return u.uid
}

// setUID 设置 UID
func (u *User) setUID(ctx context.Context, id uint32) error {
	u.uid = id
	log.Debug(ctx, "domain - entity - setUID uid: %d", id)
	return nil
}

// Email 获取 email
func (u *User) Email() string {
	return u.email
}

// setUID 设置 email
func (u *User) setEmail(ctx context.Context, email string) error {
	u.email = email
	log.Debug(ctx, "domain - entity - setEmail email: %s", email)
	return nil
}

// Nickname 获取 nickname
func (u *User) Nickname() string {
	return u.nickname
}

// setNickname 设置 nickname
func (u *User) setNickname(ctx context.Context, nickname string) error {
	u.nickname = nickname
	log.Debug(ctx, "domain - entity - setNickname nickname: %s", nickname)
	return nil
}

// Avatar 获取 avatar
func (u *User) Avatar() string {
	return u.avatar
}

// setAvatar 设置 avatar
func (u *User) setAvatar(ctx context.Context, avatar string) error {
	u.avatar = avatar
	log.Debug(ctx, "domain - entity - setAvatar avatar: %s", avatar)
	return nil
}

// Intro 获取 intro
func (u *User) Intro() string {
	return u.intro
}

// setIntro 设置 intro
func (u *User) setIntro(ctx context.Context, intro string) error {
	u.intro = intro
	log.Debug(ctx, "domain - entity - setIntro intro: %s", intro)
	return nil
}

// WebSite 获取 webSite
func (u *User) WebSite() string {
	return u.webSite
}

// setWebSite 设置 webSite
func (u *User) setWebSite(ctx context.Context, webSite string) error {
	u.webSite = webSite
	log.Debug(ctx, "domain - entity - setIntro webSite: %s", webSite)
	return nil
}

// IsDisable 获取 isDisable
func (u *User) IsDisable() uint32 {
	return u.isDisable
}

// setIsDisable 设置 isDisable
func (u *User) setIsDisable(ctx context.Context, isDisable uint32) error {
	u.isDisable = isDisable
	log.Debug(ctx, "domain - entity - setIntro webSite: %+v", isDisable)
	return nil
}

// Username 获取 username
func (u *User) Username() string {
	return u.username
}

// setUsername 设置 username
func (u *User) setUsername(ctx context.Context, username string) error {
	u.username = username
	log.Debug(ctx, "domain - entity - setUsername username: %s", username)
	return nil
}

// Password 获取 Password
func (u *User) Password() string {
	return u.password
}

// setPassword 设置 password
func (u *User) setPassword(ctx context.Context, password string) error {
	u.password = password
	log.Debug(ctx, "domain - entity - setPassword password: %s", password)
	return nil
}

// LoginType 获取 loginType
func (u *User) LoginType() uint32 {
	return u.loginType
}

// setLoginType 设置 loginType
func (u *User) setLoginType(ctx context.Context, loginType uint32) error {
	u.loginType = loginType
	log.Debug(ctx, "domain - entity - setLoginType loginType: %d", loginType)
	return nil
}

// IpAddress 获取 ipAddress
func (u *User) IpAddress() string {
	return u.ipAddress
}

// setIpAddress 设置 ipAddress
func (u *User) setIpAddress(ctx context.Context, ipAddress string) error {
	u.ipAddress = ipAddress
	log.Debug(ctx, "domain - entity - setIpAddress ipAddress: %s", ipAddress)
	return nil
}

// IpSource 获取 ipSource
func (u *User) IpSource() string {
	return u.ipSource
}

// setIpAddress 设置 ipSource
func (u *User) setIpSource(ctx context.Context, ipSource string) error {
	u.ipSource = ipSource
	log.Debug(ctx, "domain - entity - setIpSource ipSource: %s", ipSource)
	return nil
}

// LastLoginTime 获取 lastLoginTime
func (u *User) LastLoginTime() uint32 {
	return u.lastLoginTime
}

// setLastLoginTime 设置 lastLoginTime
func (u *User) setLastLoginTime(ctx context.Context, lastLoginTime uint32) error {
	u.lastLoginTime = lastLoginTime
	log.Debug(ctx, "domain - entity - setLastLoginTime lastLoginTime: %d", lastLoginTime)
	return nil
}

// Login 登陆
func (u *User) Login(password string) error {
	if u.password != password {
		return fmt.Errorf("passworld check fail")
	}
	return nil
}

// ChangePassword 修改密码
func (u *User) ChangePassword(ctx context.Context, oldPassword string, newPassword string) error {
	if u.password != oldPassword {
		return fmt.Errorf("old passworld check fail")
	}
	u.password = newPassword
	return nil
}
