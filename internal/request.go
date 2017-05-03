// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/ory-am/fosite (interfaces: Requester)

package internal

import (
	gomock "github.com/golang/mock/gomock"
	fosite "github.com/ory/fosite"
	url "net/url"
	time "time"
)

// Mock of Requester interface
type MockRequester struct {
	ctrl     *gomock.Controller
	recorder *_MockRequesterRecorder
}

// Recorder for MockRequester (not exported)
type _MockRequesterRecorder struct {
	mock *MockRequester
}

func NewMockRequester(ctrl *gomock.Controller) *MockRequester {
	mock := &MockRequester{ctrl: ctrl}
	mock.recorder = &_MockRequesterRecorder{mock}
	return mock
}

func (_m *MockRequester) EXPECT() *_MockRequesterRecorder {
	return _m.recorder
}

func (_m *MockRequester) AppendRequestedScope(_param0 string) {
	_m.ctrl.Call(_m, "AppendRequestedScope", _param0)
}

func (_mr *_MockRequesterRecorder) AppendRequestedScope(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AppendRequestedScope", arg0)
}

func (_m *MockRequester) GetClient() fosite.Client {
	ret := _m.ctrl.Call(_m, "GetClient")
	ret0, _ := ret[0].(fosite.Client)
	return ret0
}

func (_mr *_MockRequesterRecorder) GetClient() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetClient")
}

func (_m *MockRequester) GetGrantedScopes() fosite.Arguments {
	ret := _m.ctrl.Call(_m, "GetGrantedScopes")
	ret0, _ := ret[0].(fosite.Arguments)
	return ret0
}

func (_mr *_MockRequesterRecorder) GetGrantedScopes() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetGrantedScopes")
}

func (_m *MockRequester) GetID() string {
	ret := _m.ctrl.Call(_m, "GetID")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockRequesterRecorder) GetID() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetID")
}

func (_m *MockRequester) GetRequestForm() url.Values {
	ret := _m.ctrl.Call(_m, "GetRequestForm")
	ret0, _ := ret[0].(url.Values)
	return ret0
}

func (_mr *_MockRequesterRecorder) GetRequestForm() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetRequestForm")
}

func (_m *MockRequester) GetRequestedAt() time.Time {
	ret := _m.ctrl.Call(_m, "GetRequestedAt")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

func (_mr *_MockRequesterRecorder) GetRequestedAt() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetRequestedAt")
}

func (_m *MockRequester) GetRequestedScopes() fosite.Arguments {
	ret := _m.ctrl.Call(_m, "GetRequestedScopes")
	ret0, _ := ret[0].(fosite.Arguments)
	return ret0
}

func (_mr *_MockRequesterRecorder) GetRequestedScopes() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetRequestedScopes")
}

func (_m *MockRequester) GetSession() fosite.Session {
	ret := _m.ctrl.Call(_m, "GetSession")
	ret0, _ := ret[0].(fosite.Session)
	return ret0
}

func (_mr *_MockRequesterRecorder) GetSession() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetSession")
}

func (_m *MockRequester) GrantScope(_param0 string) {
	_m.ctrl.Call(_m, "GrantScope", _param0)
}

func (_mr *_MockRequesterRecorder) GrantScope(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GrantScope", arg0)
}

func (_m *MockRequester) Merge(_param0 fosite.Requester) {
	_m.ctrl.Call(_m, "Merge", _param0)
}

func (_mr *_MockRequesterRecorder) Merge(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Merge", arg0)
}

func (_m *MockRequester) SetRequestedScopes(_param0 fosite.Arguments) {
	_m.ctrl.Call(_m, "SetRequestedScopes", _param0)
}

func (_mr *_MockRequesterRecorder) SetRequestedScopes(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetRequestedScopes", arg0)
}

func (_m *MockRequester) SetSession(_param0 fosite.Session) {
	_m.ctrl.Call(_m, "SetSession", _param0)
}

func (_mr *_MockRequesterRecorder) SetSession(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetSession", arg0)
}
