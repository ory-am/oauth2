// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ory/fosite/handler/rfc8628 (interfaces: DeviceCodeStorage)

// Package internal is a generated GoMock package.
package internal

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	fosite "github.com/ory/fosite"
)

// MockDeviceCodeStorage is a mock of DeviceCodeStorage interface.
type MockDeviceCodeStorage struct {
	ctrl     *gomock.Controller
	recorder *MockDeviceCodeStorageMockRecorder
}

// MockDeviceCodeStorageMockRecorder is the mock recorder for MockDeviceCodeStorage.
type MockDeviceCodeStorageMockRecorder struct {
	mock *MockDeviceCodeStorage
}

// NewMockDeviceCodeStorage creates a new mock instance.
func NewMockDeviceCodeStorage(ctrl *gomock.Controller) *MockDeviceCodeStorage {
	mock := &MockDeviceCodeStorage{ctrl: ctrl}
	mock.recorder = &MockDeviceCodeStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDeviceCodeStorage) EXPECT() *MockDeviceCodeStorageMockRecorder {
	return m.recorder
}

// CreateDeviceCodeSession mocks base method.
func (m *MockDeviceCodeStorage) CreateDeviceCodeSession(arg0 context.Context, arg1 string, arg2 fosite.Requester) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDeviceCodeSession", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateDeviceCodeSession indicates an expected call of CreateDeviceCodeSession.
func (mr *MockDeviceCodeStorageMockRecorder) CreateDeviceCodeSession(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDeviceCodeSession", reflect.TypeOf((*MockDeviceCodeStorage)(nil).CreateDeviceCodeSession), arg0, arg1, arg2)
}

// GetDeviceCodeSession mocks base method.
func (m *MockDeviceCodeStorage) GetDeviceCodeSession(arg0 context.Context, arg1 string, arg2 fosite.Session) (fosite.Requester, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeviceCodeSession", arg0, arg1, arg2)
	ret0, _ := ret[0].(fosite.Requester)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeviceCodeSession indicates an expected call of GetDeviceCodeSession.
func (mr *MockDeviceCodeStorageMockRecorder) GetDeviceCodeSession(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeviceCodeSession", reflect.TypeOf((*MockDeviceCodeStorage)(nil).GetDeviceCodeSession), arg0, arg1, arg2)
}

// InvalidateDeviceCodeSession mocks base method.
func (m *MockDeviceCodeStorage) InvalidateDeviceCodeSession(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InvalidateDeviceCodeSession", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// InvalidateDeviceCodeSession indicates an expected call of InvalidateDeviceCodeSession.
func (mr *MockDeviceCodeStorageMockRecorder) InvalidateDeviceCodeSession(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InvalidateDeviceCodeSession", reflect.TypeOf((*MockDeviceCodeStorage)(nil).InvalidateDeviceCodeSession), arg0, arg1)
}

// UpdateDeviceCodeSession mocks base method.
func (m *MockDeviceCodeStorage) UpdateDeviceCodeSession(arg0 context.Context, arg1 string, arg2 fosite.Requester) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateDeviceCodeSession", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateDeviceCodeSession indicates an expected call of UpdateDeviceCodeSession.
func (mr *MockDeviceCodeStorageMockRecorder) UpdateDeviceCodeSession(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDeviceCodeSession", reflect.TypeOf((*MockDeviceCodeStorage)(nil).UpdateDeviceCodeSession), arg0, arg1, arg2)
}
