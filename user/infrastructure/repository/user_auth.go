package repository

import (
	"fmt"

	"github.com/baker-yuan/go-blog/all_packaged_library/lib"
	"github.com/baker-yuan/go-blog/user/domain/entity"
)

type UserRepoImpl struct {
}

func (u *UserRepoImpl) FindByUsername(username string) (*entity.User, error) {
	var (
		userAuth entity.UserAuth
		userInfo entity.UserInfo
	)
	var (
		user *entity.User
	)
	// 查询账号
	if tx := lib.DB.Raw("select * from tb_user_auth where username = ? ", username).Scan(&userAuth); tx.Error != nil {
		fmt.Println(tx.Error)
	}
	if userAuth.ID == 0 {
		return nil, nil
	}
	// 查询账号信息
	if tx := lib.DB.Raw("select * from tb_user_info where id = ? ", userAuth.UserInfoId).Scan(&userInfo); tx.Error != nil {
		fmt.Println(tx.Error)
	}
	user = &entity.User{
		Info: userInfo,
		Auth: userAuth,
	}
	return user, nil
}
