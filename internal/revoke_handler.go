// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/ory/fosite (interfaces: RevocationHandler)

package internal

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	fosite "github.com/ory/fosite"
)

// Mock of RevocationHandler interface
type MockRevocationHandler struct {
	ctrl     *gomock.Controller
	recorder *_MockRevocationHandlerRecorder
}

// Recorder for MockRevocationHandler (not exported)
type _MockRevocationHandlerRecorder struct {
	mock *MockRevocationHandler
}

func NewMockRevocationHandler(ctrl *gomock.Controller) *MockRevocationHandler {
	mock := &MockRevocationHandler{ctrl: ctrl}
	mock.recorder = &_MockRevocationHandlerRecorder{mock}
	return mock
}

func (_m *MockRevocationHandler) EXPECT() *_MockRevocationHandlerRecorder {
	return _m.recorder
}

func (_m *MockRevocationHandler) RevokeToken(_param0 context.Context, _param1 string, _param2 fosite.TokenType) error {
	ret := _m.ctrl.Call(_m, "RevokeToken", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockRevocationHandlerRecorder) RevokeToken(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RevokeToken", arg0, arg1, arg2)
}
