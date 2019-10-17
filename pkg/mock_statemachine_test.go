// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kevinconway/remouse/pkg (interfaces: StateMachine)

// Package remouse is a generated GoMock package.
package remouse

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockStateMachine is a mock of StateMachine interface
type MockStateMachine struct {
	ctrl     *gomock.Controller
	recorder *MockStateMachineMockRecorder
}

// MockStateMachineMockRecorder is the mock recorder for MockStateMachine
type MockStateMachineMockRecorder struct {
	mock *MockStateMachine
}

// NewMockStateMachine creates a new mock instance
func NewMockStateMachine(ctrl *gomock.Controller) *MockStateMachine {
	mock := &MockStateMachine{ctrl: ctrl}
	mock.recorder = &MockStateMachineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStateMachine) EXPECT() *MockStateMachineMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockStateMachine) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockStateMachineMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockStateMachine)(nil).Close))
}

// Current mocks base method
func (m *MockStateMachine) Current() StateChange {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Current")
	ret0, _ := ret[0].(StateChange)
	return ret0
}

// Current indicates an expected call of Current
func (mr *MockStateMachineMockRecorder) Current() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Current", reflect.TypeOf((*MockStateMachine)(nil).Current))
}

// Next mocks base method
func (m *MockStateMachine) Next() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Next")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Next indicates an expected call of Next
func (mr *MockStateMachineMockRecorder) Next() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockStateMachine)(nil).Next))
}
