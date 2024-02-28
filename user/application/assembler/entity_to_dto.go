// Package assembler 转换器
package assembler

import (
	"github.com/baker-yuan/go-blog/user/application/dto"
	"github.com/baker-yuan/go-blog/user/domain/entity"
)

// GenUserDetailDTO 把 Entity 转换成 DTO
func GenUserDetailDTO(userEntity *entity.User) *dto.UserDetailDTO {
	userDetailDTO := &dto.UserDetailDTO{}
	// 信息
	userDetailDTO.ID = userEntity.UID()
	userDetailDTO.Email = userEntity.Email()
	userDetailDTO.Nickname = userEntity.Nickname()
	userDetailDTO.Avatar = userEntity.Avatar()
	userDetailDTO.Intro = userEntity.Intro()
	userDetailDTO.WebSite = userEntity.WebSite()
	userDetailDTO.IsDisable = userEntity.IsDisable()
	// 账户
	userDetailDTO.Username = userEntity.Username()
	userDetailDTO.LoginType = userEntity.LoginType()
	userDetailDTO.IpAddress = userEntity.IpAddress()
	userDetailDTO.IpSource = userEntity.IpSource()
	userDetailDTO.LastLoginTime = userEntity.LastLoginTime()
	return userDetailDTO
}

func GenUserDetailListDTO(users []*entity.User) []*dto.UserDetailDTO {
	var res = make([]*dto.UserDetailDTO, 0)
	for _, v := range users {
		res = append(res, GenUserDetailDTO(v))
	}
	return res
}
