package persistence

import (
	"github.com/baker-yuan/go-blog/go_blog/application/user/domain/entity"
	"github.com/baker-yuan/go-blog/go_blog/common/db"
)

func UserPosToEntity(users []*User) []*entity.User {
	res := make([]*entity.User, 0)
	for _, user := range users {
		res = append(res, UserPoToEntity(user))
	}
	return res
}
func UserPoToEntity(user *User) *entity.User {
	res := &entity.User{
		// 账号信息
		UID:      user.UID,
		Username: user.Username,
		Password: user.Password,
		Salt:     user.Salt,
		UserType: user.UserType,
		// 基本信息
		Email:    user.Email,
		Nickname: user.Nickname,
		Avatar:   user.Avatar,
		Intro:    user.Intro,
		WebSite:  user.WebSite,
		// 状态
		Status: user.Status,
		// 三方登录
		LoginType: user.LoginType,
		UnionID:   user.UnionID,
		// 公共字段
		IsDeleted:    user.IsDeleted,
		CreateUserID: user.CreateUserID,
		UpdateUserID: user.UpdateUserID,
		CreateTime:   uint32(user.CreateTime),
		UpdateTime:   uint32(user.UpdateTime),
	}
	return res
}

func UserEntityToPo(user *entity.User) *User {
	res := &User{
		UID:          user.UID,
		Username:     user.Username,
		Password:     user.Password,
		Salt:         user.Salt,
		UserType:     user.UserType,
		Email:        user.Email,
		Nickname:     user.Nickname,
		Avatar:       user.Avatar,
		Intro:        user.Intro,
		WebSite:      user.WebSite,
		Status:       user.Status,
		LoginType:    user.LoginType,
		UnionID:      user.UnionID,
		IsDeleted:    user.IsDeleted,
		CreateUserID: user.CreateUserID,
		UpdateUserID: user.UpdateUserID,
		CreateTime:   db.Timestamp(user.CreateTime),
		UpdateTime:   db.Timestamp(user.UpdateTime),
	}
	return res
}
