// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/pb/users/users_grpc.pb.go

// Package users_mock is a generated GoMock package.
package users_mock

import (
	context "context"
	users "heisenbug/identification/internal/pb/users"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockUsersClient is a mock of UsersClient interface.
type MockUsersClient struct {
	ctrl     *gomock.Controller
	recorder *MockUsersClientMockRecorder
}

// MockUsersClientMockRecorder is the mock recorder for MockUsersClient.
type MockUsersClientMockRecorder struct {
	mock *MockUsersClient
}

// NewMockUsersClient creates a new mock instance.
func NewMockUsersClient(ctrl *gomock.Controller) *MockUsersClient {
	mock := &MockUsersClient{ctrl: ctrl}
	mock.recorder = &MockUsersClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUsersClient) EXPECT() *MockUsersClientMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockUsersClient) CreateUser(ctx context.Context, in *users.CreateUserRequest, opts ...grpc.CallOption) (*users.CreateUserResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateUser", varargs...)
	ret0, _ := ret[0].(*users.CreateUserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUsersClientMockRecorder) CreateUser(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUsersClient)(nil).CreateUser), varargs...)
}

// GetInfo mocks base method.
func (m *MockUsersClient) GetInfo(ctx context.Context, in *users.GetInfoRequest, opts ...grpc.CallOption) (*users.GetInfoResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetInfo", varargs...)
	ret0, _ := ret[0].(*users.GetInfoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInfo indicates an expected call of GetInfo.
func (mr *MockUsersClientMockRecorder) GetInfo(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInfo", reflect.TypeOf((*MockUsersClient)(nil).GetInfo), varargs...)
}

// RemoveUser mocks base method.
func (m *MockUsersClient) RemoveUser(ctx context.Context, in *users.RemoveUserRequest, opts ...grpc.CallOption) (*users.RemoveUserResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RemoveUser", varargs...)
	ret0, _ := ret[0].(*users.RemoveUserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveUser indicates an expected call of RemoveUser.
func (mr *MockUsersClientMockRecorder) RemoveUser(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveUser", reflect.TypeOf((*MockUsersClient)(nil).RemoveUser), varargs...)
}

// UpdateUserLevel mocks base method.
func (m *MockUsersClient) UpdateUserLevel(ctx context.Context, in *users.UpdateUserLevelRequest, opts ...grpc.CallOption) (*users.UpdateUserLevelResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateUserLevel", varargs...)
	ret0, _ := ret[0].(*users.UpdateUserLevelResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUserLevel indicates an expected call of UpdateUserLevel.
func (mr *MockUsersClientMockRecorder) UpdateUserLevel(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserLevel", reflect.TypeOf((*MockUsersClient)(nil).UpdateUserLevel), varargs...)
}

// MockUsersServer is a mock of UsersServer interface.
type MockUsersServer struct {
	ctrl     *gomock.Controller
	recorder *MockUsersServerMockRecorder
}

// MockUsersServerMockRecorder is the mock recorder for MockUsersServer.
type MockUsersServerMockRecorder struct {
	mock *MockUsersServer
}

// NewMockUsersServer creates a new mock instance.
func NewMockUsersServer(ctrl *gomock.Controller) *MockUsersServer {
	mock := &MockUsersServer{ctrl: ctrl}
	mock.recorder = &MockUsersServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUsersServer) EXPECT() *MockUsersServerMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockUsersServer) CreateUser(arg0 context.Context, arg1 *users.CreateUserRequest) (*users.CreateUserResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(*users.CreateUserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUsersServerMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUsersServer)(nil).CreateUser), arg0, arg1)
}

// GetInfo mocks base method.
func (m *MockUsersServer) GetInfo(arg0 context.Context, arg1 *users.GetInfoRequest) (*users.GetInfoResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInfo", arg0, arg1)
	ret0, _ := ret[0].(*users.GetInfoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInfo indicates an expected call of GetInfo.
func (mr *MockUsersServerMockRecorder) GetInfo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInfo", reflect.TypeOf((*MockUsersServer)(nil).GetInfo), arg0, arg1)
}

// RemoveUser mocks base method.
func (m *MockUsersServer) RemoveUser(arg0 context.Context, arg1 *users.RemoveUserRequest) (*users.RemoveUserResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveUser", arg0, arg1)
	ret0, _ := ret[0].(*users.RemoveUserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveUser indicates an expected call of RemoveUser.
func (mr *MockUsersServerMockRecorder) RemoveUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveUser", reflect.TypeOf((*MockUsersServer)(nil).RemoveUser), arg0, arg1)
}

// UpdateUserLevel mocks base method.
func (m *MockUsersServer) UpdateUserLevel(arg0 context.Context, arg1 *users.UpdateUserLevelRequest) (*users.UpdateUserLevelResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserLevel", arg0, arg1)
	ret0, _ := ret[0].(*users.UpdateUserLevelResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUserLevel indicates an expected call of UpdateUserLevel.
func (mr *MockUsersServerMockRecorder) UpdateUserLevel(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserLevel", reflect.TypeOf((*MockUsersServer)(nil).UpdateUserLevel), arg0, arg1)
}

// mustEmbedUnimplementedUsersServer mocks base method.
func (m *MockUsersServer) mustEmbedUnimplementedUsersServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedUsersServer")
}

// mustEmbedUnimplementedUsersServer indicates an expected call of mustEmbedUnimplementedUsersServer.
func (mr *MockUsersServerMockRecorder) mustEmbedUnimplementedUsersServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedUsersServer", reflect.TypeOf((*MockUsersServer)(nil).mustEmbedUnimplementedUsersServer))
}

// MockUnsafeUsersServer is a mock of UnsafeUsersServer interface.
type MockUnsafeUsersServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeUsersServerMockRecorder
}

// MockUnsafeUsersServerMockRecorder is the mock recorder for MockUnsafeUsersServer.
type MockUnsafeUsersServerMockRecorder struct {
	mock *MockUnsafeUsersServer
}

// NewMockUnsafeUsersServer creates a new mock instance.
func NewMockUnsafeUsersServer(ctrl *gomock.Controller) *MockUnsafeUsersServer {
	mock := &MockUnsafeUsersServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeUsersServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeUsersServer) EXPECT() *MockUnsafeUsersServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedUsersServer mocks base method.
func (m *MockUnsafeUsersServer) mustEmbedUnimplementedUsersServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedUsersServer")
}

// mustEmbedUnimplementedUsersServer indicates an expected call of mustEmbedUnimplementedUsersServer.
func (mr *MockUnsafeUsersServerMockRecorder) mustEmbedUnimplementedUsersServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedUsersServer", reflect.TypeOf((*MockUnsafeUsersServer)(nil).mustEmbedUnimplementedUsersServer))
}