// Copyright © 2024 Ory Corp
// SPDX-License-Identifier: Apache-2.0

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ory/fosite/handler/rfc7523 (interfaces: RFC7523KeyStorage)

// Package internal is a generated GoMock package.
package internal

import (
	context "context"
	reflect "reflect"
	time "time"

	v3 "github.com/go-jose/go-jose/v3"
	gomock "github.com/golang/mock/gomock"
)

// MockRFC7523KeyStorage is a mock of RFC7523KeyStorage interface
type MockRFC7523KeyStorage struct {
	ctrl     *gomock.Controller
	recorder *MockRFC7523KeyStorageMockRecorder
}

// MockRFC7523KeyStorageMockRecorder is the mock recorder for MockRFC7523KeyStorage
type MockRFC7523KeyStorageMockRecorder struct {
	mock *MockRFC7523KeyStorage
}

// NewMockRFC7523KeyStorage creates a new mock instance
func NewMockRFC7523KeyStorage(ctrl *gomock.Controller) *MockRFC7523KeyStorage {
	mock := &MockRFC7523KeyStorage{ctrl: ctrl}
	mock.recorder = &MockRFC7523KeyStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRFC7523KeyStorage) EXPECT() *MockRFC7523KeyStorageMockRecorder {
	return m.recorder
}

// GetPublicKey mocks base method
func (m *MockRFC7523KeyStorage) GetPublicKey(arg0 context.Context, arg1, arg2, arg3 string) (*v3.JSONWebKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPublicKey", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*v3.JSONWebKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPublicKey indicates an expected call of GetPublicKey
func (mr *MockRFC7523KeyStorageMockRecorder) GetPublicKey(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPublicKey", reflect.TypeOf((*MockRFC7523KeyStorage)(nil).GetPublicKey), arg0, arg1, arg2, arg3)
}

// GetPublicKeyScopes mocks base method
func (m *MockRFC7523KeyStorage) GetPublicKeyScopes(arg0 context.Context, arg1, arg2, arg3 string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPublicKeyScopes", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPublicKeyScopes indicates an expected call of GetPublicKeyScopes
func (mr *MockRFC7523KeyStorageMockRecorder) GetPublicKeyScopes(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPublicKeyScopes", reflect.TypeOf((*MockRFC7523KeyStorage)(nil).GetPublicKeyScopes), arg0, arg1, arg2, arg3)
}

// GetPublicKeys mocks base method
func (m *MockRFC7523KeyStorage) GetPublicKeys(arg0 context.Context, arg1, arg2 string) (*v3.JSONWebKeySet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPublicKeys", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v3.JSONWebKeySet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPublicKeys indicates an expected call of GetPublicKeys
func (mr *MockRFC7523KeyStorageMockRecorder) GetPublicKeys(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPublicKeys", reflect.TypeOf((*MockRFC7523KeyStorage)(nil).GetPublicKeys), arg0, arg1, arg2)
}

// IsJWTUsed mocks base method
func (m *MockRFC7523KeyStorage) IsJWTUsed(arg0 context.Context, arg1 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsJWTUsed", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsJWTUsed indicates an expected call of IsJWTUsed
func (mr *MockRFC7523KeyStorageMockRecorder) IsJWTUsed(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsJWTUsed", reflect.TypeOf((*MockRFC7523KeyStorage)(nil).IsJWTUsed), arg0, arg1)
}

// MarkJWTUsedForTime mocks base method
func (m *MockRFC7523KeyStorage) MarkJWTUsedForTime(arg0 context.Context, arg1 string, arg2 time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarkJWTUsedForTime", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// MarkJWTUsedForTime indicates an expected call of MarkJWTUsedForTime
func (mr *MockRFC7523KeyStorageMockRecorder) MarkJWTUsedForTime(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarkJWTUsedForTime", reflect.TypeOf((*MockRFC7523KeyStorage)(nil).MarkJWTUsedForTime), arg0, arg1, arg2)
}
