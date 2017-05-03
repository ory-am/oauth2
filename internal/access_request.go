// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/ory-am/fosite (interfaces: AccessRequester)

package internal

import (
	gomock "github.com/golang/mock/gomock"
	fosite "github.com/ory/fosite"
	url "net/url"
	time "time"
)

// Mock of AccessRequester interface
type MockAccessRequester struct {
	ctrl     *gomock.Controller
	recorder *_MockAccessRequesterRecorder
}

// Recorder for MockAccessRequester (not exported)
type _MockAccessRequesterRecorder struct {
	mock *MockAccessRequester
}

func NewMockAccessRequester(ctrl *gomock.Controller) *MockAccessRequester {
	mock := &MockAccessRequester{ctrl: ctrl}
	mock.recorder = &_MockAccessRequesterRecorder{mock}
	return mock
}

func (_m *MockAccessRequester) EXPECT() *_MockAccessRequesterRecorder {
	return _m.recorder
}

func (_m *MockAccessRequester) AppendRequestedScope(_param0 string) {
	_m.ctrl.Call(_m, "AppendRequestedScope", _param0)
}

func (_mr *_MockAccessRequesterRecorder) AppendRequestedScope(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AppendRequestedScope", arg0)
}

func (_m *MockAccessRequester) GetClient() fosite.Client {
	ret := _m.ctrl.Call(_m, "GetClient")
	ret0, _ := ret[0].(fosite.Client)
	return ret0
}

func (_mr *_MockAccessRequesterRecorder) GetClient() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetClient")
}

func (_m *MockAccessRequester) GetGrantTypes() fosite.Arguments {
	ret := _m.ctrl.Call(_m, "GetGrantTypes")
	ret0, _ := ret[0].(fosite.Arguments)
	return ret0
}

func (_mr *_MockAccessRequesterRecorder) GetGrantTypes() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetGrantTypes")
}

func (_m *MockAccessRequester) GetGrantedScopes() fosite.Arguments {
	ret := _m.ctrl.Call(_m, "GetGrantedScopes")
	ret0, _ := ret[0].(fosite.Arguments)
	return ret0
}

func (_mr *_MockAccessRequesterRecorder) GetGrantedScopes() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetGrantedScopes")
}

func (_m *MockAccessRequester) GetID() string {
	ret := _m.ctrl.Call(_m, "GetID")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockAccessRequesterRecorder) GetID() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetID")
}

func (_m *MockAccessRequester) GetRequestForm() url.Values {
	ret := _m.ctrl.Call(_m, "GetRequestForm")
	ret0, _ := ret[0].(url.Values)
	return ret0
}

func (_mr *_MockAccessRequesterRecorder) GetRequestForm() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetRequestForm")
}

func (_m *MockAccessRequester) GetRequestedAt() time.Time {
	ret := _m.ctrl.Call(_m, "GetRequestedAt")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

func (_mr *_MockAccessRequesterRecorder) GetRequestedAt() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetRequestedAt")
}

func (_m *MockAccessRequester) GetRequestedScopes() fosite.Arguments {
	ret := _m.ctrl.Call(_m, "GetRequestedScopes")
	ret0, _ := ret[0].(fosite.Arguments)
	return ret0
}

func (_mr *_MockAccessRequesterRecorder) GetRequestedScopes() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetRequestedScopes")
}

func (_m *MockAccessRequester) GetSession() fosite.Session {
	ret := _m.ctrl.Call(_m, "GetSession")
	ret0, _ := ret[0].(fosite.Session)
	return ret0
}

func (_mr *_MockAccessRequesterRecorder) GetSession() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetSession")
}

func (_m *MockAccessRequester) GrantScope(_param0 string) {
	_m.ctrl.Call(_m, "GrantScope", _param0)
}

func (_mr *_MockAccessRequesterRecorder) GrantScope(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GrantScope", arg0)
}

func (_m *MockAccessRequester) Merge(_param0 fosite.Requester) {
	_m.ctrl.Call(_m, "Merge", _param0)
}

func (_mr *_MockAccessRequesterRecorder) Merge(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Merge", arg0)
}

func (_m *MockAccessRequester) SetRequestedScopes(_param0 fosite.Arguments) {
	_m.ctrl.Call(_m, "SetRequestedScopes", _param0)
}

func (_mr *_MockAccessRequesterRecorder) SetRequestedScopes(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetRequestedScopes", arg0)
}

func (_m *MockAccessRequester) SetSession(_param0 fosite.Session) {
	_m.ctrl.Call(_m, "SetSession", _param0)
}

func (_mr *_MockAccessRequesterRecorder) SetSession(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetSession", arg0)
}
