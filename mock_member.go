// Code generated by MockGen. DO NOT EDIT.
// Source: member.go

// Package embedetcd is a generated GoMock package.
package embedetcd

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockMember is a mock of Member interface.
type MockMember struct {
	ctrl     *gomock.Controller
	recorder *MockMemberMockRecorder
}

// MockMemberMockRecorder is the mock recorder for MockMember.
type MockMemberMockRecorder struct {
	mock *MockMember
}

// NewMockMember creates a new mock instance.
func NewMockMember(ctrl *gomock.Controller) *MockMember {
	mock := &MockMember{ctrl: ctrl}
	mock.recorder = &MockMemberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMember) EXPECT() *MockMemberMockRecorder {
	return m.recorder
}

// GetLeaderAddr mocks base method.
func (m *MockMember) GetLeaderAddr() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLeaderAddr")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetLeaderAddr indicates an expected call of GetLeaderAddr.
func (mr *MockMemberMockRecorder) GetLeaderAddr() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLeaderAddr", reflect.TypeOf((*MockMember)(nil).GetLeaderAddr))
}

// GetLeaderID mocks base method.
func (m *MockMember) GetLeaderID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLeaderID")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetLeaderID indicates an expected call of GetLeaderID.
func (mr *MockMemberMockRecorder) GetLeaderID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLeaderID", reflect.TypeOf((*MockMember)(nil).GetLeaderID))
}

// IsLeader mocks base method.
func (m *MockMember) IsLeader() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsLeader")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsLeader indicates an expected call of IsLeader.
func (mr *MockMemberMockRecorder) IsLeader() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsLeader", reflect.TypeOf((*MockMember)(nil).IsLeader))
}

// IsReady mocks base method.
func (m *MockMember) IsReady() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsReady")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsReady indicates an expected call of IsReady.
func (mr *MockMemberMockRecorder) IsReady() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsReady", reflect.TypeOf((*MockMember)(nil).IsReady))
}

// RegisterMembershipChangedProcessor mocks base method.
func (m *MockMember) RegisterMembershipChangedProcessor(arg0 MembershipEventProcessor) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterMembershipChangedProcessor", arg0)
}

// RegisterMembershipChangedProcessor indicates an expected call of RegisterMembershipChangedProcessor.
func (mr *MockMemberMockRecorder) RegisterMembershipChangedProcessor(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterMembershipChangedProcessor", reflect.TypeOf((*MockMember)(nil).RegisterMembershipChangedProcessor), arg0)
}

// ResignIfLeader mocks base method.
func (m *MockMember) ResignIfLeader() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ResignIfLeader")
}

// ResignIfLeader indicates an expected call of ResignIfLeader.
func (mr *MockMemberMockRecorder) ResignIfLeader() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResignIfLeader", reflect.TypeOf((*MockMember)(nil).ResignIfLeader))
}