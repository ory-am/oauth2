// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/ory/fosite (interfaces: TokenIntrospector)

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

// Mock of TokenIntrospector interface
type MockTokenIntrospector struct {
	ctrl     *gomock.Controller
	recorder *_MockTokenIntrospectorRecorder
}

// Recorder for MockTokenIntrospector (not exported)
type _MockTokenIntrospectorRecorder struct {
	mock *MockTokenIntrospector
}

func NewMockTokenIntrospector(ctrl *gomock.Controller) *MockTokenIntrospector {
	mock := &MockTokenIntrospector{ctrl: ctrl}
	mock.recorder = &_MockTokenIntrospectorRecorder{mock}
	return mock
}

func (_m *MockTokenIntrospector) EXPECT() *_MockTokenIntrospectorRecorder {
	return _m.recorder
}

func (_m *MockTokenIntrospector) IntrospectToken(_param0 context.Context, _param1 string, _param2 fosite.TokenType, _param3 fosite.AccessRequester, _param4 []string) error {
	ret := _m.ctrl.Call(_m, "IntrospectToken", _param0, _param1, _param2, _param3, _param4)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockTokenIntrospectorRecorder) IntrospectToken(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "IntrospectToken", arg0, arg1, arg2, arg3, arg4)
}
