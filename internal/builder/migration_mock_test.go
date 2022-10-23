// Code generated by MockGen. DO NOT EDIT.
// Source: internal/builder/migration.go

// Package builder_test is a generated GoMock package.
package builder_test

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockMigration is a mock of Migration interface.
type MockMigration struct {
	ctrl     *gomock.Controller
	recorder *MockMigrationMockRecorder
}

// MockMigrationMockRecorder is the mock recorder for MockMigration.
type MockMigrationMockRecorder struct {
	mock *MockMigration
}

// NewMockMigration creates a new mock instance.
func NewMockMigration(ctrl *gomock.Controller) *MockMigration {
	mock := &MockMigration{ctrl: ctrl}
	mock.recorder = &MockMigrationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMigration) EXPECT() *MockMigrationMockRecorder {
	return m.recorder
}

// Info mocks base method.
func (m *MockMigration) Info() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Info")
	ret0, _ := ret[0].(string)
	return ret0
}

// Info indicates an expected call of Info.
func (mr *MockMigrationMockRecorder) Info() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockMigration)(nil).Info))
}

// Version mocks base method.
func (m *MockMigration) Version() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Version")
	ret0, _ := ret[0].(int64)
	return ret0
}

// Version indicates an expected call of Version.
func (mr *MockMigrationMockRecorder) Version() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Version", reflect.TypeOf((*MockMigration)(nil).Version))
}