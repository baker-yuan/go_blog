// Code generated by MockGen. DO NOT EDIT.
// Source: datasync.trpc.go

// Package datasync is a generated GoMock package.
package datasync

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	client "trpc.group/trpc-go/trpc-go/client"
)

// MockDataSyncApiService is a mock of DataSyncApiService interface.
type MockDataSyncApiService struct {
	ctrl     *gomock.Controller
	recorder *MockDataSyncApiServiceMockRecorder
}

// MockDataSyncApiServiceMockRecorder is the mock recorder for MockDataSyncApiService.
type MockDataSyncApiServiceMockRecorder struct {
	mock *MockDataSyncApiService
}

// NewMockDataSyncApiService creates a new mock instance.
func NewMockDataSyncApiService(ctrl *gomock.Controller) *MockDataSyncApiService {
	mock := &MockDataSyncApiService{ctrl: ctrl}
	mock.recorder = &MockDataSyncApiServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDataSyncApiService) EXPECT() *MockDataSyncApiServiceMockRecorder {
	return m.recorder
}

// DataChange mocks base method.
func (m *MockDataSyncApiService) DataChange(arg0 DataSyncApi_DataChangeServer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DataChange", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DataChange indicates an expected call of DataChange.
func (mr *MockDataSyncApiServiceMockRecorder) DataChange(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DataChange", reflect.TypeOf((*MockDataSyncApiService)(nil).DataChange), arg0)
}

// MockDataSyncApi_DataChangeServer is a mock of DataSyncApi_DataChangeServer interface.
type MockDataSyncApi_DataChangeServer struct {
	ctrl     *gomock.Controller
	recorder *MockDataSyncApi_DataChangeServerMockRecorder
}

// MockDataSyncApi_DataChangeServerMockRecorder is the mock recorder for MockDataSyncApi_DataChangeServer.
type MockDataSyncApi_DataChangeServerMockRecorder struct {
	mock *MockDataSyncApi_DataChangeServer
}

// NewMockDataSyncApi_DataChangeServer creates a new mock instance.
func NewMockDataSyncApi_DataChangeServer(ctrl *gomock.Controller) *MockDataSyncApi_DataChangeServer {
	mock := &MockDataSyncApi_DataChangeServer{ctrl: ctrl}
	mock.recorder = &MockDataSyncApi_DataChangeServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDataSyncApi_DataChangeServer) EXPECT() *MockDataSyncApi_DataChangeServerMockRecorder {
	return m.recorder
}

// Context mocks base method.
func (m *MockDataSyncApi_DataChangeServer) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockDataSyncApi_DataChangeServerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockDataSyncApi_DataChangeServer)(nil).Context))
}

// Recv mocks base method.
func (m *MockDataSyncApi_DataChangeServer) Recv() (*TableChange, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*TableChange)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv.
func (mr *MockDataSyncApi_DataChangeServerMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockDataSyncApi_DataChangeServer)(nil).Recv))
}

// RecvMsg mocks base method.
func (m_2 *MockDataSyncApi_DataChangeServer) RecvMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "RecvMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockDataSyncApi_DataChangeServerMockRecorder) RecvMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockDataSyncApi_DataChangeServer)(nil).RecvMsg), m)
}

// SendAndClose mocks base method.
func (m *MockDataSyncApi_DataChangeServer) SendAndClose(arg0 *DataChangeRsp) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendAndClose", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendAndClose indicates an expected call of SendAndClose.
func (mr *MockDataSyncApi_DataChangeServerMockRecorder) SendAndClose(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendAndClose", reflect.TypeOf((*MockDataSyncApi_DataChangeServer)(nil).SendAndClose), arg0)
}

// SendMsg mocks base method.
func (m_2 *MockDataSyncApi_DataChangeServer) SendMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockDataSyncApi_DataChangeServerMockRecorder) SendMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockDataSyncApi_DataChangeServer)(nil).SendMsg), m)
}

// MockDataSyncApiClientProxy is a mock of DataSyncApiClientProxy interface.
type MockDataSyncApiClientProxy struct {
	ctrl     *gomock.Controller
	recorder *MockDataSyncApiClientProxyMockRecorder
}

// MockDataSyncApiClientProxyMockRecorder is the mock recorder for MockDataSyncApiClientProxy.
type MockDataSyncApiClientProxyMockRecorder struct {
	mock *MockDataSyncApiClientProxy
}

// NewMockDataSyncApiClientProxy creates a new mock instance.
func NewMockDataSyncApiClientProxy(ctrl *gomock.Controller) *MockDataSyncApiClientProxy {
	mock := &MockDataSyncApiClientProxy{ctrl: ctrl}
	mock.recorder = &MockDataSyncApiClientProxyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDataSyncApiClientProxy) EXPECT() *MockDataSyncApiClientProxyMockRecorder {
	return m.recorder
}

// DataChange mocks base method.
func (m *MockDataSyncApiClientProxy) DataChange(ctx context.Context, opts ...client.Option) (DataSyncApi_DataChangeClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DataChange", varargs...)
	ret0, _ := ret[0].(DataSyncApi_DataChangeClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DataChange indicates an expected call of DataChange.
func (mr *MockDataSyncApiClientProxyMockRecorder) DataChange(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DataChange", reflect.TypeOf((*MockDataSyncApiClientProxy)(nil).DataChange), varargs...)
}

// MockDataSyncApi_DataChangeClient is a mock of DataSyncApi_DataChangeClient interface.
type MockDataSyncApi_DataChangeClient struct {
	ctrl     *gomock.Controller
	recorder *MockDataSyncApi_DataChangeClientMockRecorder
}

// MockDataSyncApi_DataChangeClientMockRecorder is the mock recorder for MockDataSyncApi_DataChangeClient.
type MockDataSyncApi_DataChangeClientMockRecorder struct {
	mock *MockDataSyncApi_DataChangeClient
}

// NewMockDataSyncApi_DataChangeClient creates a new mock instance.
func NewMockDataSyncApi_DataChangeClient(ctrl *gomock.Controller) *MockDataSyncApi_DataChangeClient {
	mock := &MockDataSyncApi_DataChangeClient{ctrl: ctrl}
	mock.recorder = &MockDataSyncApi_DataChangeClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDataSyncApi_DataChangeClient) EXPECT() *MockDataSyncApi_DataChangeClientMockRecorder {
	return m.recorder
}

// CloseAndRecv mocks base method.
func (m *MockDataSyncApi_DataChangeClient) CloseAndRecv() (*DataChangeRsp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseAndRecv")
	ret0, _ := ret[0].(*DataChangeRsp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CloseAndRecv indicates an expected call of CloseAndRecv.
func (mr *MockDataSyncApi_DataChangeClientMockRecorder) CloseAndRecv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseAndRecv", reflect.TypeOf((*MockDataSyncApi_DataChangeClient)(nil).CloseAndRecv))
}

// CloseSend mocks base method.
func (m *MockDataSyncApi_DataChangeClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend.
func (mr *MockDataSyncApi_DataChangeClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockDataSyncApi_DataChangeClient)(nil).CloseSend))
}

// Context mocks base method.
func (m *MockDataSyncApi_DataChangeClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockDataSyncApi_DataChangeClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockDataSyncApi_DataChangeClient)(nil).Context))
}

// RecvMsg mocks base method.
func (m_2 *MockDataSyncApi_DataChangeClient) RecvMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "RecvMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockDataSyncApi_DataChangeClientMockRecorder) RecvMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockDataSyncApi_DataChangeClient)(nil).RecvMsg), m)
}

// Send mocks base method.
func (m *MockDataSyncApi_DataChangeClient) Send(arg0 *TableChange) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockDataSyncApi_DataChangeClientMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockDataSyncApi_DataChangeClient)(nil).Send), arg0)
}

// SendMsg mocks base method.
func (m_2 *MockDataSyncApi_DataChangeClient) SendMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockDataSyncApi_DataChangeClientMockRecorder) SendMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockDataSyncApi_DataChangeClient)(nil).SendMsg), m)
}
