// Copyright © 2025 Ory Corp
// SPDX-License-Identifier: Apache-2.0

// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/ory/fosite/handler/oauth2 (interfaces: AuthorizeCodeGrantStorage)

package internal

import (
	"context"

	gomock "go.uber.org/mock/gomock"

	"github.com/ory/fosite"
)

// Mock of AuthorizeCodeGrantStorage interface
type MockAuthorizeCodeGrantStorage struct {
	ctrl     *gomock.Controller
	recorder *_MockAuthorizeCodeGrantStorageRecorder
}

// Recorder for MockAuthorizeCodeGrantStorage (not exported)
type _MockAuthorizeCodeGrantStorageRecorder struct {
	mock *MockAuthorizeCodeGrantStorage
}

func NewMockAuthorizeCodeGrantStorage(ctrl *gomock.Controller) *MockAuthorizeCodeGrantStorage {
	mock := &MockAuthorizeCodeGrantStorage{ctrl: ctrl}
	mock.recorder = &_MockAuthorizeCodeGrantStorageRecorder{mock}
	return mock
}

func (_m *MockAuthorizeCodeGrantStorage) EXPECT() *_MockAuthorizeCodeGrantStorageRecorder {
	return _m.recorder
}

func (_m *MockAuthorizeCodeGrantStorage) CreateAuthorizeCodeSession(_param0 context.Context, _param1 string, _param2 fosite.Requester) error {
	ret := _m.ctrl.Call(_m, "CreateAuthorizeCodeSession", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockAuthorizeCodeGrantStorageRecorder) CreateAuthorizeCodeSession(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateAuthorizeCodeSession", arg0, arg1, arg2)
}

func (_m *MockAuthorizeCodeGrantStorage) DeleteAuthorizeCodeSession(_param0 context.Context, _param1 string) error {
	ret := _m.ctrl.Call(_m, "DeleteAuthorizeCodeSession", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockAuthorizeCodeGrantStorageRecorder) DeleteAuthorizeCodeSession(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteAuthorizeCodeSession", arg0, arg1)
}

func (_m *MockAuthorizeCodeGrantStorage) GetAuthorizeCodeSession(_param0 context.Context, _param1 string, _param2 fosite.Session) (fosite.Requester, error) {
	ret := _m.ctrl.Call(_m, "GetAuthorizeCodeSession", _param0, _param1, _param2)
	ret0, _ := ret[0].(fosite.Requester)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAuthorizeCodeGrantStorageRecorder) GetAuthorizeCodeSession(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetAuthorizeCodeSession", arg0, arg1, arg2)
}

func (_m *MockAuthorizeCodeGrantStorage) PersistAuthorizeCodeGrantSession(_param0 context.Context, _param1 string, _param2 string, _param3 string, _param4 fosite.Requester) error {
	ret := _m.ctrl.Call(_m, "PersistAuthorizeCodeGrantSession", _param0, _param1, _param2, _param3, _param4)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockAuthorizeCodeGrantStorageRecorder) PersistAuthorizeCodeGrantSession(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PersistAuthorizeCodeGrantSession", arg0, arg1, arg2, arg3, arg4)
}
