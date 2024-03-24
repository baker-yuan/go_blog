package application

import (
	"context"
	"time"

	"github.com/baker-yuan/go-blog/application/user/domain/entity"
	"github.com/baker-yuan/go-blog/application/user/domain/repository"
	"github.com/baker-yuan/go-blog/application/user/interfaces/assembler"
	"github.com/baker-yuan/go-blog/common/jwt"
	"github.com/baker-yuan/go-blog/common/util"
	pb "github.com/baker-yuan/go-blog/protocol/user"
)

type UserAppInterface interface {
	// SearchUser 查询用户
	SearchUser(ctx context.Context, req *pb.SearchUserReq) ([]*pb.User, uint32, error)
	// UserDetail 用户详情
	UserDetail(ctx context.Context, req *pb.UserDetailReq) (*pb.User, error)
	// AddOrUpdateUser 添加修改用户
	AddOrUpdateUser(ctx context.Context, req *pb.AddOrUpdateUserReq) (uint32, error)
	// DeleteUser 删除用户
	DeleteUser(ctx context.Context, req *pb.DeleteUserReq) error

	// Login 登录
	Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginRsp, error)
	// Logout 退出
	Logout(ctx context.Context, req *pb.LogoutReq) (*pb.LogoutRsp, error)
	// Refresh 刷新token
	Refresh(ctx context.Context, req *pb.RefreshReq) (*pb.RefreshRsp, error)
}

type userApp struct {
	us      repository.UserRepository
	jwtUtil *jwt.TokenGenerator
}

// UserApp 强制userApp实现UserAppInterface
var _ UserAppInterface = &userApp{}

func NewUserApp(us repository.UserRepository, jwtUtil *jwt.TokenGenerator) UserAppInterface {
	return &userApp{
		us:      us,
		jwtUtil: jwtUtil,
	}
}

// UserDetail 用户详情
func (c *userApp) UserDetail(ctx context.Context, req *pb.UserDetailReq) (*pb.User, error) {
	user, err := c.us.GetUserByID(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	pbUser := assembler.UserEntityToModel(user)
	return pbUser, nil
}

// SearchUser 用户搜索
func (c *userApp) SearchUser(ctx context.Context, req *pb.SearchUserReq) ([]*pb.User, uint32, error) {
	users, total, err := c.us.SearchUser(ctx, req)
	if err != nil {
		return nil, 0, err
	}
	data := make([]*pb.User, 0)
	for _, user := range users {
		data = append(data, assembler.UserEntityToModel(user))
	}
	return data, total, nil
}

// AddOrUpdateUser 添加修改用户
func (c *userApp) AddOrUpdateUser(ctx context.Context, req *pb.AddOrUpdateUserReq) (uint32, error) {
	userID, err := util.SessionUtils.GetLoginUserID(ctx)
	if err != nil {
		return 0, err
	}
	if req.GetUid() == 0 {
		return c.addUser(ctx, userID, req)
	} else {
		dbUser, err := c.us.GetUserByID(ctx, req.GetUid())
		if err != nil {
			return 0, err
		}
		return c.updateUser(ctx, dbUser, userID, req)
	}
}

func (c *userApp) addUser(ctx context.Context, userID uint32, req *pb.AddOrUpdateUserReq) (uint32, error) {
	user := assembler.AddOrUpdateUserReqToEntity(req)
	user.CreateTime = uint32(time.Now().Unix())
	user.UpdateTime = uint32(time.Now().Unix())
	user.CreateUserID = userID
	user.UpdateUserID = userID
	lastInsertID, err := c.us.Save(ctx, user)
	if err != nil {
		return 0, err
	}
	return lastInsertID, nil
}

func (c *userApp) updateUser(ctx context.Context, dbUser *entity.User, userID uint32, req *pb.AddOrUpdateUserReq) (uint32, error) {
	saveUser := assembler.AddOrUpdateUserReqToEntity(req)
	saveUser.CreateTime = dbUser.CreateTime
	saveUser.CreateUserID = dbUser.CreateUserID
	saveUser.UpdateUserID = userID
	saveUser.UpdateTime = uint32(time.Now().Unix())
	if err := c.us.UpdateByID(ctx, saveUser); err != nil {
		return 0, err
	}
	return req.GetUid(), nil
}

// DeleteUser 删除用户
func (c *userApp) DeleteUser(ctx context.Context, req *pb.DeleteUserReq) error {
	// 获取用户
	user, err := c.us.GetUserByID(ctx, req.GetId())
	if err != nil {
		return err
	}
	// 删除用户
	user.Delete()
	// 持久化
	if err := c.us.UpdateByID(ctx, user); err != nil {
		return err
	}
	return nil
}

// Login 登录
func (c *userApp) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginRsp, error) {
	// 用户名+密码校验
	user, err := c.us.GetUserByEmailAndPassword(ctx, req.GetUsername(), req.GetPassword())
	if err != nil {
		return nil, err
	}
	// 生成token
	token, err := c.jwtUtil.CreateToken(ctx, jwt.NewUserInfoBuilder().UserID(user.UID).Build())
	if err != nil {
		return nil, err
	}
	// 返回
	return &pb.LoginRsp{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		AtExpires:    token.AtExpires,
		RtExpires:    token.RtExpires,
	}, nil
}

// Logout 退出
func (c *userApp) Logout(ctx context.Context, req *pb.LogoutReq) (*pb.LogoutRsp, error) {
	//TODO implement me
	panic("implement me")
}

// Refresh 刷新token
func (c *userApp) Refresh(ctx context.Context, req *pb.RefreshReq) (*pb.RefreshRsp, error) {
	//TODO implement me
	panic("implement me")
}
