// Package assembler 数据结构转换器
package assembler

import (
	"context"

	"github.com/baker-yuan/go-blog/user/domain/entity"
	"github.com/baker-yuan/go-blog/user/infrastructure/persistence/mysql/po"
)

const (
	ZeroUINT32 uint32 = 0
)

// GenUserEntity 构造 User Entity
func GenUserEntity(ctx context.Context, userInfoPO *po.UserInfoPO, userAuthPO *po.UserAuthPO) *entity.User {
	if userAuthPO.ID == ZeroUINT32 || userInfoPO.UID == ZeroUINT32 {
		return nil
	}
	userBuilder := &entity.UserBuilder{}
	userBuilder.
		UID(userInfoPO.UID).
		Email(userInfoPO.Email).
		Nickname(userInfoPO.Nickname).
		Avatar(userInfoPO.Avatar).
		Intro(userInfoPO.Intro).
		WebSite(userInfoPO.WebSite).
		IsDisable(userInfoPO.IsDisable)
	userBuilder.
		Username(userAuthPO.Username).
		Password(userAuthPO.Password).
		LoginType(userAuthPO.LoginType).
		IpAddress(userAuthPO.IpAddress).
		IpSource(userAuthPO.IpSource).
		LastLoginTime(userAuthPO.LastLoginTime)
	userEntity, _ := userBuilder.Build(ctx)
	return userEntity
}

func GenUserListEntity(ctx context.Context, userInfos []*po.UserInfoPO, userAuths []*po.UserAuthPO) []*entity.User {
	var (
		res = make([]*entity.User, 0)
	)
	userAuthMap := make(map[uint32]*po.UserAuthPO)
	for _, v := range userAuths {
		userAuthMap[v.UserInfoId] = v
	}
	for _, userInfoPO := range userInfos {
		userAuthPO := userAuthMap[userInfoPO.UID]
		res = append(res, GenUserEntity(ctx, userInfoPO, userAuthPO))
	}
	return res
}
