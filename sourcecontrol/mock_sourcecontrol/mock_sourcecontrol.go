// Automatically generated by MockGen. DO NOT EDIT!
// Source: bitbucket.org/michaellockwood/github-repos/sourcecontrol (interfaces: SourceControlGateway)

package mock_sourcecontrol

import (
	sourcecontrol "github.com/michaellockwood/github-repos/sourcecontrol"
	gomock "github.com/golang/mock/gomock"
)

// Mock of SourceControlGateway interface
type MockSourceControlGateway struct {
	ctrl     *gomock.Controller
	recorder *_MockSourceControlGatewayRecorder
}

// Recorder for MockSourceControlGateway (not exported)
type _MockSourceControlGatewayRecorder struct {
	mock *MockSourceControlGateway
}

func NewMockSourceControlGateway(ctrl *gomock.Controller) *MockSourceControlGateway {
	mock := &MockSourceControlGateway{ctrl: ctrl}
	mock.recorder = &_MockSourceControlGatewayRecorder{mock}
	return mock
}

func (_m *MockSourceControlGateway) EXPECT() *_MockSourceControlGatewayRecorder {
	return _m.recorder
}

func (_m *MockSourceControlGateway) GetCommits(_param0 string, _param1 string, _param2 int) ([]sourcecontrol.Commit, error) {
	ret := _m.ctrl.Call(_m, "GetCommits", _param0, _param1, _param2)
	ret0, _ := ret[0].([]sourcecontrol.Commit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockSourceControlGatewayRecorder) GetCommits(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetCommits", arg0, arg1, arg2)
}

func (_m *MockSourceControlGateway) GetRepositories(_param0 string, _param1 int) ([]sourcecontrol.Repository, error) {
	ret := _m.ctrl.Call(_m, "GetRepositories", _param0, _param1)
	ret0, _ := ret[0].([]sourcecontrol.Repository)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockSourceControlGatewayRecorder) GetRepositories(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetRepositories", arg0, arg1)
}
