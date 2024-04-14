package assembler

import (
	"github.com/baker-yuan/go-blog/go_blog/application/user/domain/entity"
	pb "github.com/baker-yuan/go-blog/go_blog/protocol/user"
)

// UserEntityListToModel entity转pb
func UserEntityListToModel(users []*entity.User, uidToNameMap map[uint32]string) []*pb.User {
	res := make([]*pb.User, 0)
	for _, user := range users {
		res = append(res, UserEntityToModel(user, uidToNameMap))
	}
	return res
}

// UserEntityToModel entity转pb
func UserEntityToModel(user *entity.User, uidToNameMap map[uint32]string) *pb.User {
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
		CreateUserId: user.CreateUserID,
		UpdateUserId: user.UpdateUserID,
		CreateTime:   user.CreateTime,
		UpdateTime:   user.UpdateTime,
		// 查询补充
	}
	modelRes.CreateUser = uidToNameMap[user.CreateUserID]
	modelRes.UpdateUser = uidToNameMap[user.UpdateUserID]
	return modelRes
}

// AddOrUpdateUserReqToEntity pb转entity
func AddOrUpdateUserReqToEntity(pbUser *pb.AddOrUpdateUserReq) *entity.User {
	entityRes := &entity.User{
		// 账号信息
		UID:      pbUser.Uid,
		Username: pbUser.Username,
		Password: pbUser.Password,
		// 基本信息
		Email:    pbUser.Email,
		Nickname: pbUser.Nickname,
		Avatar:   pbUser.Avatar,
		Intro:    pbUser.Intro,
		WebSite:  pbUser.WebSite,
		// 状态
		Status: pbUser.Status,
	}
	return entityRes
}
