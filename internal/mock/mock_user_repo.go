// Code generated by MockGen. DO NOT EDIT.
// Source: user.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	swagger "github.com/gnarlyman/dbpractice/swagger"
	gomock "github.com/golang/mock/gomock"
)

// MockIUserRepo is a mock of IUserRepo interface.
type MockIUserRepo struct {
	ctrl     *gomock.Controller
	recorder *MockIUserRepoMockRecorder
}

// MockIUserRepoMockRecorder is the mock recorder for MockIUserRepo.
type MockIUserRepoMockRecorder struct {
	mock *MockIUserRepo
}

// NewMockIUserRepo creates a new mock instance.
func NewMockIUserRepo(ctrl *gomock.Controller) *MockIUserRepo {
	mock := &MockIUserRepo{ctrl: ctrl}
	mock.recorder = &MockIUserRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUserRepo) EXPECT() *MockIUserRepoMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockIUserRepo) CreateUser(ctx context.Context, user *swagger.User) (*swagger.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, user)
	ret0, _ := ret[0].(*swagger.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockIUserRepoMockRecorder) CreateUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockIUserRepo)(nil).CreateUser), ctx, user)
}

// DeleteUser mocks base method.
func (m *MockIUserRepo) DeleteUser(ctx context.Context, userID int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", ctx, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockIUserRepoMockRecorder) DeleteUser(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockIUserRepo)(nil).DeleteUser), ctx, userID)
}

// FindUsers mocks base method.
func (m *MockIUserRepo) FindUsers(ctx context.Context) ([]*swagger.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUsers", ctx)
	ret0, _ := ret[0].([]*swagger.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUsers indicates an expected call of FindUsers.
func (mr *MockIUserRepoMockRecorder) FindUsers(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUsers", reflect.TypeOf((*MockIUserRepo)(nil).FindUsers), ctx)
}

// GetUser mocks base method.
func (m *MockIUserRepo) GetUser(ctx context.Context, userID int32) (*swagger.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", ctx, userID)
	ret0, _ := ret[0].(*swagger.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockIUserRepoMockRecorder) GetUser(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockIUserRepo)(nil).GetUser), ctx, userID)
}

// GetUserWithPassword mocks base method.
func (m *MockIUserRepo) GetUserWithPassword(ctx context.Context, userID int32) (*swagger.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserWithPassword", ctx, userID)
	ret0, _ := ret[0].(*swagger.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserWithPassword indicates an expected call of GetUserWithPassword.
func (mr *MockIUserRepoMockRecorder) GetUserWithPassword(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserWithPassword", reflect.TypeOf((*MockIUserRepo)(nil).GetUserWithPassword), ctx, userID)
}

// PatchUser mocks base method.
func (m *MockIUserRepo) PatchUser(ctx context.Context, userID int32, userUpdate *swagger.User) (*swagger.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PatchUser", ctx, userID, userUpdate)
	ret0, _ := ret[0].(*swagger.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PatchUser indicates an expected call of PatchUser.
func (mr *MockIUserRepoMockRecorder) PatchUser(ctx, userID, userUpdate interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchUser", reflect.TypeOf((*MockIUserRepo)(nil).PatchUser), ctx, userID, userUpdate)
}

// UpdateUser mocks base method.
func (m *MockIUserRepo) UpdateUser(ctx context.Context, user *swagger.User) (*swagger.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", ctx, user)
	ret0, _ := ret[0].(*swagger.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockIUserRepoMockRecorder) UpdateUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockIUserRepo)(nil).UpdateUser), ctx, user)
}
