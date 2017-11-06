// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/ory/fosite/handler/oauth2 (interfaces: CoreStrategy)

// Copyright © 2017 Aeneas Rekkas <aeneas+oss@aeneas.io>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package internal

import (
	context "context"

	gomock "github.com/golang/mock/gomock"
	fosite "github.com/ory/fosite"
)

// Mock of CoreStrategy interface
type MockCoreStrategy struct {
	ctrl     *gomock.Controller
	recorder *_MockCoreStrategyRecorder
}

// Recorder for MockCoreStrategy (not exported)
type _MockCoreStrategyRecorder struct {
	mock *MockCoreStrategy
}

func NewMockCoreStrategy(ctrl *gomock.Controller) *MockCoreStrategy {
	mock := &MockCoreStrategy{ctrl: ctrl}
	mock.recorder = &_MockCoreStrategyRecorder{mock}
	return mock
}

func (_m *MockCoreStrategy) EXPECT() *_MockCoreStrategyRecorder {
	return _m.recorder
}

func (_m *MockCoreStrategy) AccessTokenSignature(_param0 string) string {
	ret := _m.ctrl.Call(_m, "AccessTokenSignature", _param0)
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockCoreStrategyRecorder) AccessTokenSignature(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AccessTokenSignature", arg0)
}

func (_m *MockCoreStrategy) AuthorizeCodeSignature(_param0 string) string {
	ret := _m.ctrl.Call(_m, "AuthorizeCodeSignature", _param0)
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockCoreStrategyRecorder) AuthorizeCodeSignature(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AuthorizeCodeSignature", arg0)
}

func (_m *MockCoreStrategy) GenerateAccessToken(_param0 context.Context, _param1 fosite.Requester) (string, string, error) {
	ret := _m.ctrl.Call(_m, "GenerateAccessToken", _param0, _param1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockCoreStrategyRecorder) GenerateAccessToken(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GenerateAccessToken", arg0, arg1)
}

func (_m *MockCoreStrategy) GenerateAuthorizeCode(_param0 context.Context, _param1 fosite.Requester) (string, string, error) {
	ret := _m.ctrl.Call(_m, "GenerateAuthorizeCode", _param0, _param1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockCoreStrategyRecorder) GenerateAuthorizeCode(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GenerateAuthorizeCode", arg0, arg1)
}

func (_m *MockCoreStrategy) GenerateRefreshToken(_param0 context.Context, _param1 fosite.Requester) (string, string, error) {
	ret := _m.ctrl.Call(_m, "GenerateRefreshToken", _param0, _param1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockCoreStrategyRecorder) GenerateRefreshToken(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GenerateRefreshToken", arg0, arg1)
}

func (_m *MockCoreStrategy) RefreshTokenSignature(_param0 string) string {
	ret := _m.ctrl.Call(_m, "RefreshTokenSignature", _param0)
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockCoreStrategyRecorder) RefreshTokenSignature(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RefreshTokenSignature", arg0)
}

func (_m *MockCoreStrategy) ValidateAccessToken(_param0 context.Context, _param1 fosite.Requester, _param2 string) error {
	ret := _m.ctrl.Call(_m, "ValidateAccessToken", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockCoreStrategyRecorder) ValidateAccessToken(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ValidateAccessToken", arg0, arg1, arg2)
}

func (_m *MockCoreStrategy) ValidateAuthorizeCode(_param0 context.Context, _param1 fosite.Requester, _param2 string) error {
	ret := _m.ctrl.Call(_m, "ValidateAuthorizeCode", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockCoreStrategyRecorder) ValidateAuthorizeCode(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ValidateAuthorizeCode", arg0, arg1, arg2)
}

func (_m *MockCoreStrategy) ValidateRefreshToken(_param0 context.Context, _param1 fosite.Requester, _param2 string) error {
	ret := _m.ctrl.Call(_m, "ValidateRefreshToken", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockCoreStrategyRecorder) ValidateRefreshToken(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ValidateRefreshToken", arg0, arg1, arg2)
}
