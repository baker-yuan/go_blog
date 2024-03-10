// Code generated by MockGen. DO NOT EDIT.
// Source: user.trpc.go

// Package user is a generated GoMock package.
package user

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	client "trpc.group/trpc-go/trpc-go/client"
)

// MockUserApiService is a mock of UserApiService interface.
type MockUserApiService struct {
	ctrl     *gomock.Controller
	recorder *MockUserApiServiceMockRecorder
}

// MockUserApiServiceMockRecorder is the mock recorder for MockUserApiService.
type MockUserApiServiceMockRecorder struct {
	mock *MockUserApiService
}

// NewMockUserApiService creates a new mock instance.
func NewMockUserApiService(ctrl *gomock.Controller) *MockUserApiService {
	mock := &MockUserApiService{ctrl: ctrl}
	mock.recorder = &MockUserApiServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserApiService) EXPECT() *MockUserApiServiceMockRecorder {
	return m.recorder
}

// AddOrUpdateUser mocks base method.
func (m *MockUserApiService) AddOrUpdateUser(ctx context.Context, req *AddOrUpdateUserReq) (*AddOrUpdateRsp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddOrUpdateUser", ctx, req)
	ret0, _ := ret[0].(*AddOrUpdateRsp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddOrUpdateUser indicates an expected call of AddOrUpdateUser.
func (mr *MockUserApiServiceMockRecorder) AddOrUpdateUser(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddOrUpdateUser", reflect.TypeOf((*MockUserApiService)(nil).AddOrUpdateUser), ctx, req)
}

// DeleteUser mocks base method.
func (m *MockUserApiService) DeleteUser(ctx context.Context, req *DeleteUserReq) (*EmptyRsp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", ctx, req)
	ret0, _ := ret[0].(*EmptyRsp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockUserApiServiceMockRecorder) DeleteUser(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockUserApiService)(nil).DeleteUser), ctx, req)
}

// SearchUser mocks base method.
func (m *MockUserApiService) SearchUser(ctx context.Context, req *SearchUserReq) (*SearchUserRsp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchUser", ctx, req)
	ret0, _ := ret[0].(*SearchUserRsp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchUser indicates an expected call of SearchUser.
func (mr *MockUserApiServiceMockRecorder) SearchUser(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchUser", reflect.TypeOf((*MockUserApiService)(nil).SearchUser), ctx, req)
}

// UserDetail mocks base method.
func (m *MockUserApiService) UserDetail(ctx context.Context, req *UserDetailReq) (*User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserDetail", ctx, req)
	ret0, _ := ret[0].(*User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserDetail indicates an expected call of UserDetail.
func (mr *MockUserApiServiceMockRecorder) UserDetail(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserDetail", reflect.TypeOf((*MockUserApiService)(nil).UserDetail), ctx, req)
}

// MockUserApiClientProxy is a mock of UserApiClientProxy interface.
type MockUserApiClientProxy struct {
	ctrl     *gomock.Controller
	recorder *MockUserApiClientProxyMockRecorder
}

// MockUserApiClientProxyMockRecorder is the mock recorder for MockUserApiClientProxy.
type MockUserApiClientProxyMockRecorder struct {
	mock *MockUserApiClientProxy
}

// NewMockUserApiClientProxy creates a new mock instance.
func NewMockUserApiClientProxy(ctrl *gomock.Controller) *MockUserApiClientProxy {
	mock := &MockUserApiClientProxy{ctrl: ctrl}
	mock.recorder = &MockUserApiClientProxyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserApiClientProxy) EXPECT() *MockUserApiClientProxyMockRecorder {
	return m.recorder
}

// AddOrUpdateUser mocks base method.
func (m *MockUserApiClientProxy) AddOrUpdateUser(ctx context.Context, req *AddOrUpdateUserReq, opts ...client.Option) (*AddOrUpdateRsp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, req}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddOrUpdateUser", varargs...)
	ret0, _ := ret[0].(*AddOrUpdateRsp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddOrUpdateUser indicates an expected call of AddOrUpdateUser.
func (mr *MockUserApiClientProxyMockRecorder) AddOrUpdateUser(ctx, req interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, req}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddOrUpdateUser", reflect.TypeOf((*MockUserApiClientProxy)(nil).AddOrUpdateUser), varargs...)
}

// DeleteUser mocks base method.
func (m *MockUserApiClientProxy) DeleteUser(ctx context.Context, req *DeleteUserReq, opts ...client.Option) (*EmptyRsp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, req}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteUser", varargs...)
	ret0, _ := ret[0].(*EmptyRsp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockUserApiClientProxyMockRecorder) DeleteUser(ctx, req interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, req}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockUserApiClientProxy)(nil).DeleteUser), varargs...)
}

// SearchUser mocks base method.
func (m *MockUserApiClientProxy) SearchUser(ctx context.Context, req *SearchUserReq, opts ...client.Option) (*SearchUserRsp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, req}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SearchUser", varargs...)
	ret0, _ := ret[0].(*SearchUserRsp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchUser indicates an expected call of SearchUser.
func (mr *MockUserApiClientProxyMockRecorder) SearchUser(ctx, req interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, req}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchUser", reflect.TypeOf((*MockUserApiClientProxy)(nil).SearchUser), varargs...)
}

// UserDetail mocks base method.
func (m *MockUserApiClientProxy) UserDetail(ctx context.Context, req *UserDetailReq, opts ...client.Option) (*User, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, req}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UserDetail", varargs...)
	ret0, _ := ret[0].(*User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserDetail indicates an expected call of UserDetail.
func (mr *MockUserApiClientProxyMockRecorder) UserDetail(ctx, req interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, req}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserDetail", reflect.TypeOf((*MockUserApiClientProxy)(nil).UserDetail), varargs...)
}
