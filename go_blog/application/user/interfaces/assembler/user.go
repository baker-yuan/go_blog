package assembler

import (
	"github.com/baker-yuan/go-blog/application/user/domain/entity"
	"github.com/baker-yuan/go-blog/common/db"
	pb "github.com/baker-yuan/go-blog/protocol/user"
)

// UserEntityToModel entity转pb
func UserEntityToModel(user *entity.User) *pb.User {
	modelRes := &pb.User{
		// 账号信息
		Uid:      user.UID,
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
		// 三方登录
		LoginType: user.LoginType,
		UnionId:   user.UnionID,
		// 状态
		Status:    user.Status,
		IsDeleted: bool(user.IsDeleted),
		// 公共字段
		CreateTime: user.CreateTime,
		UpdateTime: user.UpdateTime,
	}
	return modelRes
}

// AddOrUpdateUserReqToEntity pb转entity
func AddOrUpdateUserReqToEntity(pbUser *pb.AddOrUpdateUserReq) *entity.User {
	entityRes := &entity.User{
		// 账号信息
		UID:      pbUser.Uid,
		Username: pbUser.Username,
		Password: pbUser.Password,
		Salt:     pbUser.Salt,
		UserType: pbUser.UserType,
		// 基本信息
		Email:    pbUser.Email,
		Nickname: pbUser.Nickname,
		Avatar:   pbUser.Avatar,
		Intro:    pbUser.Intro,
		WebSite:  pbUser.WebSite,
		// 三方登录
		LoginType: pbUser.LoginType,
		UnionID:   pbUser.UnionId,
		// 状态
		Status:    pbUser.Status,
		IsDeleted: db.BoolBit(pbUser.IsDeleted),
	}
	return entityRes
}
