// Package adapter 端口-适配器模式的 adapter
package adapter

import (
	"context"

	"github.com/baker-yuan/go-blog/user/domain/entity"
	"github.com/baker-yuan/go-blog/user/infrastructure/assembler"
	"github.com/baker-yuan/go-blog/user/infrastructure/persistence/mysql/dao"
	"github.com/baker-yuan/go-blog/user/infrastructure/persistence/mysql/po"
)

// UserAdapter User 的 Adapter
type UserAdapter struct {
	dao dao.UserDAO
}

func (u UserAdapter) ListUsers(ctx context.Context, current uint32, size uint32, nickname string, loginType uint32) ([]*entity.User, error) {
	var (
		userIDs   []uint32
		err       error
		userAuths []*po.UserAuthPO
		userInfos []*po.UserInfoPO
	)
	// 查询满足条件的用户ID
	userIDs, err = u.dao.ListUsers(ctx, current, size, nickname, loginType)
	if err != nil {
		return nil, err
	}
	// 查询账号
	userAuths, err = u.dao.GetUserAuthByInfoIDs(ctx, userIDs)
	if err != nil {
		return nil, err
	}
	// 查询账号信息
	userInfos, err = u.dao.GetUserInfoByIDs(ctx, userIDs)
	if err != nil {
		return nil, err
	}
	// 类型转换
	return assembler.GenUserListEntity(ctx, userInfos, userAuths), nil
}

func (u UserAdapter) FindByUsername(ctx context.Context, username string) (*entity.User, error) {
	var (
		userAuth *po.UserAuthPO
		userInfo *po.UserInfoPO
		err      error
	)
	// 查询账号
	userAuth, err = u.dao.GetUserAuthByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	// 查询账号信息
	userInfo, err = u.dao.GetUserInfoByUserInfoId(ctx, userAuth.UserInfoId)

	// 类型转换
	return assembler.GenUserEntity(ctx, userInfo, userAuth), nil
}
