// Package dao DAO 层
package dao

import (
	"context"
	"fmt"

	"github.com/baker-yuan/go-blog/all_packaged_library/base/db"
	"github.com/baker-yuan/go-blog/user/infrastructure/persistence/mysql/po"
)

// UserDAO User 的 DAO
type UserDAO struct {
}

func (u UserDAO) GetUserAuthByUsername(ctx context.Context, username string) (*po.UserAuthPO, error) {
	var (
		userAuth po.UserAuthPO
	)
	if tx := db.DB.Raw("select * from tb_user_auth where username = ? ", username).Scan(&userAuth); tx.Error != nil {
		return nil, tx.Error
	}
	return &userAuth, nil
}

func (u UserDAO) GetUserInfoByUserInfoId(ctx context.Context, userInfoId uint32) (*po.UserInfoPO, error) {
	var (
		userInfo po.UserInfoPO
	)
	if tx := db.DB.Raw("select * from tb_user_info where uid = ? ", userInfoId).Scan(&userInfo); tx.Error != nil {
		return nil, tx.Error
	}
	return &userInfo, nil
}

func (u UserDAO) ListUsers(ctx context.Context, current uint32, size uint32, nickname string, loginType uint32) ([]uint32, error) {
	var (
		userIDs = make([]uint32, 0)
	)

	sql := "select info.uid from tb_user_info info join tb_user_auth auth on info.uid = auth.user_info_id where 1 = 1"
	if len(nickname) != 0 {
		sql += " and info.nickname like '%" + nickname + "%'"
	}
	if loginType != 0 {
		sql += fmt.Sprintf(" and auth.login_type = %d ", loginType)
	}
	sql += " order by info.uid desc"
	sql += fmt.Sprintf(" limit %d,%d", (current-1)*size, size)
	if tx := db.DB.Raw(sql).Scan(&userIDs); tx.Error != nil {
		return nil, tx.Error
	}

	return userIDs, nil
}

func (u UserDAO) GetUserAuthByInfoIDs(ctx context.Context, infoIDs []uint32) ([]*po.UserAuthPO, error) {
	var (
		userAuths []*po.UserAuthPO
	)
	if tx := db.DB.Raw("select * from tb_user_auth where user_info_id in ? ", infoIDs).Scan(&userAuths); tx.Error != nil {
		return nil, tx.Error
	}
	return userAuths, nil
}

func (u UserDAO) GetUserInfoByIDs(ctx context.Context, ids []uint32) ([]*po.UserInfoPO, error) {
	var (
		userInfos []*po.UserInfoPO
	)
	if tx := db.DB.Raw("select * from tb_user_info where uid in ? ", ids).Scan(&userInfos); tx.Error != nil {
		return nil, tx.Error
	}
	return userInfos, nil
}
