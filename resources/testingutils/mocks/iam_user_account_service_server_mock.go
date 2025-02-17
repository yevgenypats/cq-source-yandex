// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/yandex-cloud/go-genproto/yandex/cloud/iam/v1 (interfaces: UserAccountServiceServer)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	iam "github.com/yandex-cloud/go-genproto/yandex/cloud/iam/v1"
	reflect "reflect"
)

// MockUserAccountServiceServer is a mock of UserAccountServiceServer interface
type MockUserAccountServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockUserAccountServiceServerMockRecorder
}

// MockUserAccountServiceServerMockRecorder is the mock recorder for MockUserAccountServiceServer
type MockUserAccountServiceServerMockRecorder struct {
	mock *MockUserAccountServiceServer
}

// NewMockUserAccountServiceServer creates a new mock instance
func NewMockUserAccountServiceServer(ctrl *gomock.Controller) *MockUserAccountServiceServer {
	mock := &MockUserAccountServiceServer{ctrl: ctrl}
	mock.recorder = &MockUserAccountServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserAccountServiceServer) EXPECT() *MockUserAccountServiceServerMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockUserAccountServiceServer) Get(arg0 context.Context, arg1 *iam.GetUserAccountRequest) (*iam.UserAccount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*iam.UserAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockUserAccountServiceServerMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockUserAccountServiceServer)(nil).Get), arg0, arg1)
}
