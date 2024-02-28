package grpc

import (
	"context"
	"net/http"

	"github.com/baker-yuan/go-blog/all_packaged_library/base/log"
	"github.com/baker-yuan/go-blog/all_packaged_library/grpc/auth"
	pb "github.com/baker-yuan/go-blog/proto/user"
	"github.com/baker-yuan/go-blog/user/application/service"
	"github.com/baker-yuan/go-blog/user/ui/grpc/assembler"
	"github.com/baker-yuan/go-blog/user/ui/port"
)

type UserServerImpl struct {
	AppService  *service.AppService
	MetricsPort port.MetricsPort
}

// AdminUpdatePwd 管理员修改密码
func (u UserServerImpl) AdminUpdatePwd(ctx context.Context, req *pb.AdminUpdatePwdReq) (*pb.BaseResponse, error) {
	var (
		rsp     = &pb.BaseResponse{}
		err     error
		current *auth.CurrentUser
	)

	current, err = auth.GetCurrentUser(ctx)
	if err != nil {
		// 记录日志
		log.Error(ctx, "ui - service - gRPC AdminUpdatePwd err: %v", err)
		// 上报监控
		u.MetricsPort.CounterIncr("gRPC.AdminUpdatePwd.GetCurrentUser.ERR")
		rsp.Code, rsp.Message = http.StatusForbidden, http.StatusText(http.StatusForbidden)
		return nil, err
	}

	// 转换数据: 得到 changePwdCMD
	changePwdCMD := assembler.GenChangePwdCMD(current.Username, req)

	// 调用 AppService
	err = u.AppService.ChangePassword(ctx, changePwdCMD)
	if err != nil {
		// 记录日志
		log.Error(ctx, "ui - service - gRPC AdminUpdatePwd err: %v", err)
		// 上报监控
		u.MetricsPort.CounterIncr("gRPC.AdminUpdatePwd.ChangePassword.ERR")
		rsp.Code, rsp.Message = http.StatusBadRequest, http.StatusText(http.StatusBadRequest)
		return nil, err
	}

	return rsp, nil
}

// AdminLogin 管理员登陆
func (u UserServerImpl) AdminLogin(ctx context.Context, req *pb.AdminLoginReq) (*pb.AdminLoginRsp, error) {
	var (
		rsp = &pb.AdminLoginRsp{}
	)
	// 转换数据: 得到 loginCMD
	loginCMD := assembler.GenNamePassLoginCMD(req)
	// 调用 AppService
	userDetail, err := u.AppService.Login(ctx, loginCMD)
	if err != nil {
		// 记录日志
		log.Error(ctx, "ui - service - gRPC Login err: %v", err)
		// 上报监控
		u.MetricsPort.CounterIncr("gRPC.Login.ERR")
		rsp.Code, rsp.Message = http.StatusBadRequest, http.StatusText(http.StatusBadRequest)
		return nil, err
	}

	rsp.Code, rsp.Message = http.StatusOK, http.StatusText(http.StatusOK)
	// 转换数据
	rsp.Data = assembler.GenAdminLoginRsp(userDetail)

	// 打印日志
	log.Debug(ctx, "ui - service - gRPC Login rsp: %+v", rsp)
	return rsp, nil
}

// AdminListUsers 管理员查询后台用户列表
func (u UserServerImpl) AdminListUsers(ctx context.Context, req *pb.AdminListUserReq) (*pb.AdminListUserRsp, error) {
	var (
		rsp = &pb.AdminListUserRsp{}
	)
	// 转换数据: 得到 userListQRY
	userListQRY := assembler.GenGetUserListQRY(req)

	// 调用 AppService
	users, err := u.AppService.ListUsers(ctx, userListQRY)
	if err != nil {
		// 记录日志
		log.Error(ctx, "ui - service - AdminListUsers err: %v", err)
		// 上报监控
		u.MetricsPort.CounterIncr("tRPC.AdminListUsers.ERR")
		return nil, err
	}

	// 转换数据
	rsp.Data = assembler.GenAdminListUsersRsp(users)

	// 打印日志
	log.Debug(ctx, "ui - service - gRPC AdminListUsersRsp rsp: %+v", rsp)

	return rsp, nil
}
