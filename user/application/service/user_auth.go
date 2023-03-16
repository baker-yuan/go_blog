package service

import (
	"github.com/baker-yuan/go-blog/user/application/cqe"
	"github.com/baker-yuan/go-blog/user/application/dto"
	"github.com/baker-yuan/go-blog/user/domain/entity"
	"github.com/baker-yuan/go-blog/user/domain/repo"
)

type UserAuthService struct {
	userRepo repo.UserRepo
}

func NewUserAuthService(userRepo repo.UserRepo) UserAuthService {
	userAuthService := UserAuthService{
		userRepo: userRepo,
	}
	return userAuthService
}

func (ua *UserAuthService) Login(param cqe.PasswordLoginVO) (*dto.UserDetailDTO, error) {
	var (
		user          *entity.User
		err           error
		userDetailDTO *dto.UserDetailDTO
	)
	user, err = ua.userRepo.FindByUsername(*param.Username)
	if err != nil {
		return nil, err
	}

	userDetailDTO = &dto.UserDetailDTO{
		ID:         user.Auth.ID,
		LoginType:  user.Auth.LoginType,
		Username:   user.Auth.Username,
		Password:   user.Auth.Password,
		UserInfoId: user.Info.ID,
		Email:      user.Info.Email,
		Nickname:   user.Info.Nickname,
		Avatar:     user.Info.Avatar,
		Intro:      user.Info.Intro,
		WebSite:    user.Info.WebSite,
		IsDisable:  user.Info.IsDisable,
		// RoleList:   labels,
	}

	return userDetailDTO, nil
}

// // 封装登录信息
// return u.convertUserDetail(userAuth), nil
// // 封装用户登录信息
// //
// //	@user	用户账号
// func (u *UserRepoImpl) convertUserDetail(user *entity.UserAuth) *dto.UserDetailDTO {
// 	var (
// 		userInfo      entity.UserInfo
// 		labels        []string
// 		userDetailDTO *dto.UserDetailDTO
// 	)
// 	// 查询账号信息
// 	if tx := lib.DB.Raw("select * from tb_user_info where id = ? ", user.UserInfoId).Scan(&userInfo); tx.Error != nil {
// 		fmt.Println(tx.Error)
// 	}
//
// 	// 查询账号角色
// 	// labels = ua.userRepo.ListRolesByUserInfoId(userInfo.ID)
//
// 	// 查询账号点赞信息
// 	// 获取设备信息
// 	// 封装权限集合
//
// 	userDetailDTO = &dto.UserDetailDTO{
// 		ID:         user.ID,
// 		LoginType:  user.LoginType,
// 		Username:   user.Username,
// 		Password:   user.Password,
// 		UserInfoId: userInfo.ID,
// 		Email:      userInfo.Email,
// 		Nickname:   userInfo.Nickname,
// 		Avatar:     userInfo.Avatar,
// 		Intro:      userInfo.Intro,
// 		WebSite:    userInfo.WebSite,
// 		IsDisable:  userInfo.IsDisable,
// 		RoleList:   labels,
// 	}
//
// 	return userDetailDTO
// }
