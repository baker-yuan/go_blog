package assembler

import (
	pb "github.com/baker-yuan/go-blog/proto/user"
	"github.com/baker-yuan/go-blog/user/application/dto"
)

func dto2po(userDetail *dto.UserDetailDTO) *pb.UserDetail {
	return &pb.UserDetail{
		Id:            userDetail.ID,
		Email:         userDetail.Email,
		Nickname:      userDetail.Nickname,
		Avatar:        userDetail.Avatar,
		Intro:         userDetail.Intro,
		WebSite:       userDetail.WebSite,
		IsDisable:     userDetail.IsDisable,
		Username:      userDetail.Username,
		LoginType:     userDetail.LoginType,
		IpAddress:     userDetail.IpAddress,
		IpSource:      userDetail.IpSource,
		LastLoginTime: userDetail.LastLoginTime,
		RoleList:      userDetail.RoleList,
		Browser:       userDetail.Browser,
		Os:            userDetail.Os,
	}
}

func GenAdminLoginRsp(userDetail *dto.UserDetailDTO) *pb.UserDetail {
	return dto2po(userDetail)
}

func GenAdminListUsersRsp(users []*dto.UserDetailDTO) []*pb.UserDetail {
	res := make([]*pb.UserDetail, 0)
	for _, v := range users {
		res = append(res, dto2po(v))
	}
	return res
}
