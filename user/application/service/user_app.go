package service

import (
	"context"

	"github.com/baker-yuan/go-blog/all_packaged_library/base/log"
	"github.com/baker-yuan/go-blog/user/application/assembler"
	"github.com/baker-yuan/go-blog/user/application/dto"
	"github.com/baker-yuan/go-blog/user/application/dto/command"
	"github.com/baker-yuan/go-blog/user/application/dto/query"
	"github.com/baker-yuan/go-blog/user/domain/entity"
	"github.com/baker-yuan/go-blog/user/domain/port"
)

type AppService struct {
	UserPort port.UserPort
}

// Login 管理员登陆
func (a *AppService) Login(ctx context.Context, cmd *command.NamePassLoginCMD) (*dto.UserDetailDTO, error) {
	var (
		userEntity *entity.User
		err        error
	)
	// 打印日志
	log.Debug(ctx, "application - service - Login cmd: %+v", cmd)

	// 业务编排：取聚合的实体
	userEntity, err = a.UserPort.FindByUsername(ctx, cmd.Username)
	log.Debug(ctx, "application - service - Login userEntity: %+v", userEntity)

	// 判断聚合的实体是否存在
	if err != nil {
		log.Info(ctx, "application - service - Login a.UserPort.FindByUsername err: %+v", err)
		return nil, err
	}

	// 业务编排：登陆
	if err = userEntity.Login(cmd.Password); err != nil {
		return nil, err
	}

	// 业务编排：查询权限

	// 业务编排：转换器转换出 UI 层需要的 DTO
	return assembler.GenUserDetailDTO(userEntity), nil
}

// ChangePassword 修改密码
func (a *AppService) ChangePassword(ctx context.Context, cmd *command.ChangePwdCMD) error {
	var (
		userEntity *entity.User
		err        error
	)
	// 打印日志
	log.Debug(ctx, "application - service - ChangePassword cmd: %+v", cmd)

	// 业务编排：取聚合的实体
	userEntity, err = a.UserPort.FindByUsername(ctx, cmd.Username)
	log.Debug(ctx, "application - service - ChangePassword userEntity: %+v", userEntity)

	// 判断聚合的实体是否存在
	if err != nil {
		log.Info(ctx, "application - service - ChangePassword a.UserPort.FindByUsername err: %+v", err)
		return err
	}

	// 业务编排：修改密码
	if err = userEntity.ChangePassword(ctx, cmd.OldPassword, cmd.NewPassword); err != nil {
		return err
	}

	return nil
}

// ListUsers 查询后台用户列表
func (a *AppService) ListUsers(ctx context.Context, qry *query.GetUserListQRY) ([]*dto.UserDetailDTO, error) {
	// 打印日志
	log.Debug(ctx, "application - service - ListUsers qry: %+v", qry)

	// 业务编排：通过 条件 获取聚合实体
	users, err := a.UserPort.ListUsers(ctx, qry.Current, qry.Size, qry.Nickname, qry.LoginType)
	if err != nil {
		log.Error(ctx,
			"application - service - ListUsers a.FollowPort.ListUsers err: %+v", err)
		return nil, err
	}
	// 打印日志
	log.Debug(ctx, "application - service - ListUsers a.FollowPort.ListUsers: %+v", users)

	userListDTO := assembler.GenUserDetailListDTO(users)
	return userListDTO, nil
}
