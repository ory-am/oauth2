// Copyright © 2024 Ory Corp
// SPDX-License-Identifier: Apache-2.0

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ory/fosite/handler/openid (interfaces: OpenIDConnectTokenStrategy)

// Package internal is a generated GoMock package.
package internal

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	fosite "github.com/ory/fosite"
	jwt "github.com/ory/fosite/token/jwt"
)

// MockOpenIDConnectTokenStrategy is a mock of OpenIDConnectTokenStrategy interface.
type MockOpenIDConnectTokenStrategy struct {
	ctrl     *gomock.Controller
	recorder *MockOpenIDConnectTokenStrategyMockRecorder
}

// MockOpenIDConnectTokenStrategyMockRecorder is the mock recorder for MockOpenIDConnectTokenStrategy.
type MockOpenIDConnectTokenStrategyMockRecorder struct {
	mock *MockOpenIDConnectTokenStrategy
}

// NewMockOpenIDConnectTokenStrategy creates a new mock instance.
func NewMockOpenIDConnectTokenStrategy(ctrl *gomock.Controller) *MockOpenIDConnectTokenStrategy {
	mock := &MockOpenIDConnectTokenStrategy{ctrl: ctrl}
	mock.recorder = &MockOpenIDConnectTokenStrategyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOpenIDConnectTokenStrategy) EXPECT() *MockOpenIDConnectTokenStrategyMockRecorder {
	return m.recorder
}

// DecodeIDToken mocks base method.
func (m *MockOpenIDConnectTokenStrategy) DecodeIDToken(arg0 context.Context, arg1 fosite.Requester, arg2 string) (*jwt.Token, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DecodeIDToken", arg0, arg1, arg2)
	ret0, _ := ret[0].(*jwt.Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DecodeIDToken indicates an expected call of DecodeIDToken.
func (mr *MockOpenIDConnectTokenStrategyMockRecorder) DecodeIDToken(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DecodeIDToken", reflect.TypeOf((*MockOpenIDConnectTokenStrategy)(nil).DecodeIDToken), arg0, arg1, arg2)
}

// GenerateIDToken mocks base method.
func (m *MockOpenIDConnectTokenStrategy) GenerateIDToken(arg0 context.Context, arg1 time.Duration, arg2 fosite.Requester) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateIDToken", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateIDToken indicates an expected call of GenerateIDToken.
func (mr *MockOpenIDConnectTokenStrategyMockRecorder) GenerateIDToken(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateIDToken", reflect.TypeOf((*MockOpenIDConnectTokenStrategy)(nil).GenerateIDToken), arg0, arg1, arg2)
}
