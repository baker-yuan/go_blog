package persistence

import "github.com/baker-yuan/go-blog/go_blog/application/user/domain/entity"

// UserPOsToEntity po转entity
func UserPOsToEntity(dbUsers []*User) []*entity.User {
	res := make([]*entity.User, 0)
	for _, dbUser := range dbUsers {
		res = append(res, UserPOToEntity(dbUser))
	}
	return res
}

// UserPOToEntity po转entity
func UserPOToEntity(dbUser *User) *entity.User {
	poRes := &entity.User{
		// 账号信息
		UID:      dbUser.UID,
		Username: dbUser.Username,
		UserType: dbUser.UserType,
		// 基本信息
		Email:    dbUser.Email,
		Nickname: dbUser.Nickname,
		Avatar:   dbUser.Avatar,
		Intro:    dbUser.Intro,
		WebSite:  dbUser.WebSite,

		// 状态
		Status: dbUser.Status,

		// 三方登录
		LoginType: dbUser.LoginType,
		UnionID:   dbUser.UnionID,
		// 公共字段
		IsDeleted:    dbUser.IsDeleted,
		CreateUserID: dbUser.CreateUserID,
		UpdateUserID: dbUser.UpdateUserID,
		CreateTime:   uint32(dbUser.CreateTime),
		UpdateTime:   uint32(dbUser.UpdateTime),
	}
	return poRes
}
