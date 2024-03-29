// Code generated by MockGen. DO NOT EDIT.
// Source: apperr.go

// Package apperr is a generated GoMock package.
package apperr

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	apperr "github.com/krtsato/go-server-templates/202105-twtr/pkg/apperr"
)

// MockAppErr is a mock of AppErr interface.
type MockAppErr struct {
	ctrl     *gomock.Controller
	recorder *MockAppErrMockRecorder
}

// MockAppErrMockRecorder is the mock recorder for MockAppErr.
type MockAppErrMockRecorder struct {
	mock *MockAppErr
}

// NewMockAppErr creates a new mock instance.
func NewMockAppErr(ctrl *gomock.Controller) *MockAppErr {
	mock := &MockAppErr{ctrl: ctrl}
	mock.recorder = &MockAppErrMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAppErr) EXPECT() *MockAppErrMockRecorder {
	return m.recorder
}

// Error mocks base method.
func (m *MockAppErr) Error() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Error")
	ret0, _ := ret[0].(string)
	return ret0
}

// Error indicates an expected call of Error.
func (mr *MockAppErrMockRecorder) Error() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockAppErr)(nil).Error))
}

// ErrorCode mocks base method.
func (m *MockAppErr) ErrorCode() apperr.Code {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ErrorCode")
	ret0, _ := ret[0].(apperr.Code)
	return ret0
}

// ErrorCode indicates an expected call of ErrorCode.
func (mr *MockAppErrMockRecorder) ErrorCode() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ErrorCode", reflect.TypeOf((*MockAppErr)(nil).ErrorCode))
}

// Unwrap mocks base method.
func (m *MockAppErr) Unwrap() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unwrap")
	ret0, _ := ret[0].(error)
	return ret0
}

// Unwrap indicates an expected call of Unwrap.
func (mr *MockAppErrMockRecorder) Unwrap() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unwrap", reflect.TypeOf((*MockAppErr)(nil).Unwrap))
}
