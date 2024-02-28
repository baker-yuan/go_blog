// Package entity 领域聚合实体以及实体
package entity

import "context"

// UserBuilder User Entity 的 Builder
type UserBuilder struct {
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

// UID build uid
func (u *UserBuilder) UID(uid uint32) *UserBuilder {
	u.uid = uid
	return u
}

// Email build email
func (u *UserBuilder) Email(email string) *UserBuilder {
	u.email = email
	return u
}

// Nickname build nickname
func (u *UserBuilder) Nickname(nickname string) *UserBuilder {
	u.nickname = nickname
	return u
}

// Avatar build avatar
func (u *UserBuilder) Avatar(avatar string) *UserBuilder {
	u.avatar = avatar
	return u
}

// Intro build intro
func (u *UserBuilder) Intro(intro string) *UserBuilder {
	u.intro = intro
	return u
}

// WebSite build webSite
func (u *UserBuilder) WebSite(webSite string) *UserBuilder {
	u.webSite = webSite
	return u
}

// IsDisable build isDisable
func (u *UserBuilder) IsDisable(isDisable uint32) *UserBuilder {
	u.isDisable = isDisable
	return u
}

// Username build username
func (u *UserBuilder) Username(username string) *UserBuilder {
	u.username = username
	return u
}

// Password build password
func (u *UserBuilder) Password(password string) *UserBuilder {
	u.password = password
	return u
}

// LoginType build loginType
func (u *UserBuilder) LoginType(loginType uint32) *UserBuilder {
	u.loginType = loginType
	return u
}

// IpAddress build ipAddress
func (u *UserBuilder) IpAddress(ipAddress string) *UserBuilder {
	u.ipAddress = ipAddress
	return u
}

// IpSource build ipSource
func (u *UserBuilder) IpSource(ipSource string) *UserBuilder {
	u.ipSource = ipSource
	return u
}

// LastLoginTime build lastLoginTime
func (u *UserBuilder) LastLoginTime(lastLoginTime uint32) *UserBuilder {
	u.lastLoginTime = lastLoginTime
	return u
}

// Build build一个实体
func (u *UserBuilder) Build(ctx context.Context) (*User, error) {
	user := &User{}
	if err := user.setUID(ctx, u.uid); err != nil {
		return nil, err
	}

	if err := user.setEmail(ctx, u.email); err != nil {
		return nil, err
	}

	if err := user.setNickname(ctx, u.nickname); err != nil {
		return nil, err
	}

	if err := user.setAvatar(ctx, u.avatar); err != nil {
		return nil, err
	}
	if err := user.setIntro(ctx, u.intro); err != nil {
		return nil, err
	}

	if err := user.setWebSite(ctx, u.webSite); err != nil {
		return nil, err
	}

	if err := user.setIsDisable(ctx, u.isDisable); err != nil {
		return nil, err
	}

	if err := user.setUsername(ctx, u.username); err != nil {
		return nil, err
	}

	if err := user.setPassword(ctx, u.password); err != nil {
		return nil, err
	}
	if err := user.setLoginType(ctx, u.loginType); err != nil {
		return nil, err
	}

	if err := user.setIpAddress(ctx, u.ipAddress); err != nil {
		return nil, err
	}
	if err := user.setIpSource(ctx, u.ipSource); err != nil {
		return nil, err
	}
	if err := user.setLastLoginTime(ctx, u.lastLoginTime); err != nil {
		return nil, err
	}

	return user, nil
}
