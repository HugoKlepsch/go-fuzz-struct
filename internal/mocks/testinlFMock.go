// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/hugoklepsch/go-fuzz-all/fuzzing (interfaces: TestingF)
//
// Generated by this command:
//
//	mockgen -destination ../internal/mocks/testinlFMock.go -package mocks github.com/hugoklepsch/go-fuzz-all/fuzzing TestingF
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockTestingF is a mock of TestingF interface.
type MockTestingF struct {
	ctrl     *gomock.Controller
	recorder *MockTestingFMockRecorder
}

// MockTestingFMockRecorder is the mock recorder for MockTestingF.
type MockTestingFMockRecorder struct {
	mock *MockTestingF
}

// NewMockTestingF creates a new mock instance.
func NewMockTestingF(ctrl *gomock.Controller) *MockTestingF {
	mock := &MockTestingF{ctrl: ctrl}
	mock.recorder = &MockTestingFMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTestingF) EXPECT() *MockTestingFMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockTestingF) Add(arg0 ...any) {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Add", varargs...)
}

// Add indicates an expected call of Add.
func (mr *MockTestingFMockRecorder) Add(arg0 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockTestingF)(nil).Add), arg0...)
}

// Fuzz mocks base method.
func (m *MockTestingF) Fuzz(arg0 any) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Fuzz", arg0)
}

// Fuzz indicates an expected call of Fuzz.
func (mr *MockTestingFMockRecorder) Fuzz(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fuzz", reflect.TypeOf((*MockTestingF)(nil).Fuzz), arg0)
}
