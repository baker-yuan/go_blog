// Code generated by trpc-go/trpc-cmdline v1.0.6. DO NOT EDIT.
// source: user/user.proto

package user

import (
	"context"
	"errors"
	"fmt"

	_ "trpc.group/trpc-go/trpc-go"
	"trpc.group/trpc-go/trpc-go/client"
	"trpc.group/trpc-go/trpc-go/codec"
	_ "trpc.group/trpc-go/trpc-go/http"
	"trpc.group/trpc-go/trpc-go/server"
)

// START ======================================= Server Service Definition ======================================= START

// UserApiService defines service.
type UserApiService interface {
	// SearchUser 用户搜索
	SearchUser(ctx context.Context, req *SearchUserReq) (*SearchUserRsp, error)
	// AddOrUpdateUser 添加修改用户
	AddOrUpdateUser(ctx context.Context, req *AddOrUpdateUserReq) (*AddOrUpdateRsp, error)
	// DeleteUser 删除用户
	DeleteUser(ctx context.Context, req *DeleteUserReq) (*EmptyRsp, error)
	// UserDetail 用户详情
	UserDetail(ctx context.Context, req *UserDetailReq) (*User, error)
}

func UserApiService_SearchUser_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {
	req := &SearchUserReq{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}) (interface{}, error) {
		return svr.(UserApiService).SearchUser(ctx, reqbody.(*SearchUserReq))
	}

	var rsp interface{}
	rsp, err = filters.Filter(ctx, req, handleFunc)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func UserApiService_AddOrUpdateUser_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {
	req := &AddOrUpdateUserReq{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}) (interface{}, error) {
		return svr.(UserApiService).AddOrUpdateUser(ctx, reqbody.(*AddOrUpdateUserReq))
	}

	var rsp interface{}
	rsp, err = filters.Filter(ctx, req, handleFunc)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func UserApiService_DeleteUser_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {
	req := &DeleteUserReq{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}) (interface{}, error) {
		return svr.(UserApiService).DeleteUser(ctx, reqbody.(*DeleteUserReq))
	}

	var rsp interface{}
	rsp, err = filters.Filter(ctx, req, handleFunc)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func UserApiService_UserDetail_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {
	req := &UserDetailReq{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}) (interface{}, error) {
		return svr.(UserApiService).UserDetail(ctx, reqbody.(*UserDetailReq))
	}

	var rsp interface{}
	rsp, err = filters.Filter(ctx, req, handleFunc)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

// UserApiServer_ServiceDesc descriptor for server.RegisterService.
var UserApiServer_ServiceDesc = server.ServiceDesc{
	ServiceName: "go_blog.user.UserApi",
	HandlerType: ((*UserApiService)(nil)),
	Methods: []server.Method{
		{
			Name: "/user/admin/searchUser",
			Func: UserApiService_SearchUser_Handler,
		},
		{
			Name: "/user/admin/updateUser",
			Func: UserApiService_AddOrUpdateUser_Handler,
		},
		{
			Name: "/user/admin/deleteUser",
			Func: UserApiService_DeleteUser_Handler,
		},
		{
			Name: "/user/admin/userDetail",
			Func: UserApiService_UserDetail_Handler,
		},
		{
			Name: "/go_blog.user.UserApi/SearchUser",
			Func: UserApiService_SearchUser_Handler,
		},
		{
			Name: "/go_blog.user.UserApi/AddOrUpdateUser",
			Func: UserApiService_AddOrUpdateUser_Handler,
		},
		{
			Name: "/go_blog.user.UserApi/DeleteUser",
			Func: UserApiService_DeleteUser_Handler,
		},
		{
			Name: "/go_blog.user.UserApi/UserDetail",
			Func: UserApiService_UserDetail_Handler,
		},
	},
}

// RegisterUserApiService registers service.
func RegisterUserApiService(s server.Service, svr UserApiService) {
	if err := s.Register(&UserApiServer_ServiceDesc, svr); err != nil {
		panic(fmt.Sprintf("UserApi register error:%v", err))
	}
}

// START --------------------------------- Default Unimplemented Server Service --------------------------------- START

type UnimplementedUserApi struct{}

// SearchUser 用户搜索
func (s *UnimplementedUserApi) SearchUser(ctx context.Context, req *SearchUserReq) (*SearchUserRsp, error) {
	return nil, errors.New("rpc SearchUser of service UserApi is not implemented")
}

// AddOrUpdateUser 添加修改用户
func (s *UnimplementedUserApi) AddOrUpdateUser(ctx context.Context, req *AddOrUpdateUserReq) (*AddOrUpdateRsp, error) {
	return nil, errors.New("rpc AddOrUpdateUser of service UserApi is not implemented")
}

// DeleteUser 删除用户
func (s *UnimplementedUserApi) DeleteUser(ctx context.Context, req *DeleteUserReq) (*EmptyRsp, error) {
	return nil, errors.New("rpc DeleteUser of service UserApi is not implemented")
}

// UserDetail 用户详情
func (s *UnimplementedUserApi) UserDetail(ctx context.Context, req *UserDetailReq) (*User, error) {
	return nil, errors.New("rpc UserDetail of service UserApi is not implemented")
}

// END --------------------------------- Default Unimplemented Server Service --------------------------------- END

// END ======================================= Server Service Definition ======================================= END

// START ======================================= Client Service Definition ======================================= START

// UserApiClientProxy defines service client proxy
type UserApiClientProxy interface {
	// SearchUser 用户搜索
	SearchUser(ctx context.Context, req *SearchUserReq, opts ...client.Option) (rsp *SearchUserRsp, err error)
	// AddOrUpdateUser 添加修改用户
	AddOrUpdateUser(ctx context.Context, req *AddOrUpdateUserReq, opts ...client.Option) (rsp *AddOrUpdateRsp, err error)
	// DeleteUser 删除用户
	DeleteUser(ctx context.Context, req *DeleteUserReq, opts ...client.Option) (rsp *EmptyRsp, err error)
	// UserDetail 用户详情
	UserDetail(ctx context.Context, req *UserDetailReq, opts ...client.Option) (rsp *User, err error)
}

type UserApiClientProxyImpl struct {
	client client.Client
	opts   []client.Option
}

var NewUserApiClientProxy = func(opts ...client.Option) UserApiClientProxy {
	return &UserApiClientProxyImpl{client: client.DefaultClient, opts: opts}
}

func (c *UserApiClientProxyImpl) SearchUser(ctx context.Context, req *SearchUserReq, opts ...client.Option) (*SearchUserRsp, error) {
	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)
	msg.WithClientRPCName("/user/admin/searchUser")
	msg.WithCalleeServiceName(UserApiServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("")
	msg.WithCalleeServer("")
	msg.WithCalleeService("UserApi")
	msg.WithCalleeMethod("SearchUser")
	msg.WithSerializationType(codec.SerializationTypePB)
	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)
	rsp := &SearchUserRsp{}
	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}
	return rsp, nil
}

func (c *UserApiClientProxyImpl) AddOrUpdateUser(ctx context.Context, req *AddOrUpdateUserReq, opts ...client.Option) (*AddOrUpdateRsp, error) {
	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)
	msg.WithClientRPCName("/user/admin/updateUser")
	msg.WithCalleeServiceName(UserApiServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("")
	msg.WithCalleeServer("")
	msg.WithCalleeService("UserApi")
	msg.WithCalleeMethod("AddOrUpdateUser")
	msg.WithSerializationType(codec.SerializationTypePB)
	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)
	rsp := &AddOrUpdateRsp{}
	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}
	return rsp, nil
}

func (c *UserApiClientProxyImpl) DeleteUser(ctx context.Context, req *DeleteUserReq, opts ...client.Option) (*EmptyRsp, error) {
	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)
	msg.WithClientRPCName("/user/admin/deleteUser")
	msg.WithCalleeServiceName(UserApiServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("")
	msg.WithCalleeServer("")
	msg.WithCalleeService("UserApi")
	msg.WithCalleeMethod("DeleteUser")
	msg.WithSerializationType(codec.SerializationTypePB)
	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)
	rsp := &EmptyRsp{}
	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}
	return rsp, nil
}

func (c *UserApiClientProxyImpl) UserDetail(ctx context.Context, req *UserDetailReq, opts ...client.Option) (*User, error) {
	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)
	msg.WithClientRPCName("/user/admin/userDetail")
	msg.WithCalleeServiceName(UserApiServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("")
	msg.WithCalleeServer("")
	msg.WithCalleeService("UserApi")
	msg.WithCalleeMethod("UserDetail")
	msg.WithSerializationType(codec.SerializationTypePB)
	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)
	rsp := &User{}
	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}
	return rsp, nil
}

// END ======================================= Client Service Definition ======================================= END
