// Copyright © 2025 Ory Corp
// SPDX-License-Identifier: Apache-2.0

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ory/fosite/handler/oauth2 (interfaces: ClientCredentialsGrantStorage)
//
// Generated by this command:
//
//	mockgen -package internal -destination internal/oauth2_client_storage.go github.com/ory/fosite/handler/oauth2 ClientCredentialsGrantStorage
//

// Package internal is a generated GoMock package.
package internal

import (
	context "context"
	reflect "reflect"

	fosite "github.com/ory/fosite"
	gomock "go.uber.org/mock/gomock"
)

// MockClientCredentialsGrantStorage is a mock of ClientCredentialsGrantStorage interface.
type MockClientCredentialsGrantStorage struct {
	ctrl     *gomock.Controller
	recorder *MockClientCredentialsGrantStorageMockRecorder
	isgomock struct{}
}

// MockClientCredentialsGrantStorageMockRecorder is the mock recorder for MockClientCredentialsGrantStorage.
type MockClientCredentialsGrantStorageMockRecorder struct {
	mock *MockClientCredentialsGrantStorage
}

// NewMockClientCredentialsGrantStorage creates a new mock instance.
func NewMockClientCredentialsGrantStorage(ctrl *gomock.Controller) *MockClientCredentialsGrantStorage {
	mock := &MockClientCredentialsGrantStorage{ctrl: ctrl}
	mock.recorder = &MockClientCredentialsGrantStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClientCredentialsGrantStorage) EXPECT() *MockClientCredentialsGrantStorageMockRecorder {
	return m.recorder
}

// CreateAccessTokenSession mocks base method.
func (m *MockClientCredentialsGrantStorage) CreateAccessTokenSession(ctx context.Context, signature string, request fosite.Requester) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccessTokenSession", ctx, signature, request)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateAccessTokenSession indicates an expected call of CreateAccessTokenSession.
func (mr *MockClientCredentialsGrantStorageMockRecorder) CreateAccessTokenSession(ctx, signature, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccessTokenSession", reflect.TypeOf((*MockClientCredentialsGrantStorage)(nil).CreateAccessTokenSession), ctx, signature, request)
}

// DeleteAccessTokenSession mocks base method.
func (m *MockClientCredentialsGrantStorage) DeleteAccessTokenSession(ctx context.Context, signature string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAccessTokenSession", ctx, signature)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAccessTokenSession indicates an expected call of DeleteAccessTokenSession.
func (mr *MockClientCredentialsGrantStorageMockRecorder) DeleteAccessTokenSession(ctx, signature any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAccessTokenSession", reflect.TypeOf((*MockClientCredentialsGrantStorage)(nil).DeleteAccessTokenSession), ctx, signature)
}

// GetAccessTokenSession mocks base method.
func (m *MockClientCredentialsGrantStorage) GetAccessTokenSession(ctx context.Context, signature string, session fosite.Session) (fosite.Requester, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccessTokenSession", ctx, signature, session)
	ret0, _ := ret[0].(fosite.Requester)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccessTokenSession indicates an expected call of GetAccessTokenSession.
func (mr *MockClientCredentialsGrantStorageMockRecorder) GetAccessTokenSession(ctx, signature, session any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccessTokenSession", reflect.TypeOf((*MockClientCredentialsGrantStorage)(nil).GetAccessTokenSession), ctx, signature, session)
}
