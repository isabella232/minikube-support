// Code generated by MockGen. DO NOT EDIT.
// Source: manager.go

// Package fake is a generated GoMock package.
package fake

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockManager is a mock of Manager interface
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *MockManagerMockRecorder
}

// MockManagerMockRecorder is the mock recorder for MockManager
type MockManagerMockRecorder struct {
	mock *MockManager
}

// NewMockManager creates a new mock instance
func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &MockManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockManager) EXPECT() *MockManagerMockRecorder {
	return m.recorder
}

// Init mocks base method
func (m *MockManager) Init() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init")
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init
func (mr *MockManagerMockRecorder) Init() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockManager)(nil).Init))
}

// AddRepository mocks base method
func (m *MockManager) AddRepository(name, url string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRepository", name, url)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddRepository indicates an expected call of AddRepository
func (mr *MockManagerMockRecorder) AddRepository(name, url interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRepository", reflect.TypeOf((*MockManager)(nil).AddRepository), name, url)
}

// UpdateRepository mocks base method
func (m *MockManager) UpdateRepository() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRepository")
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateRepository indicates an expected call of UpdateRepository
func (mr *MockManagerMockRecorder) UpdateRepository() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRepository", reflect.TypeOf((*MockManager)(nil).UpdateRepository))
}

// Install mocks base method
func (m *MockManager) Install(chart, release, namespace string, values map[string]interface{}, wait bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Install", chart, release, namespace, values, wait)
}

// Install indicates an expected call of Install
func (mr *MockManagerMockRecorder) Install(chart, release, namespace, values, wait interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Install", reflect.TypeOf((*MockManager)(nil).Install), chart, release, namespace, values, wait)
}

// Uninstall mocks base method
func (m *MockManager) Uninstall(release string, purge bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Uninstall", release, purge)
}

// Uninstall indicates an expected call of Uninstall
func (mr *MockManagerMockRecorder) Uninstall(release, purge interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Uninstall", reflect.TypeOf((*MockManager)(nil).Uninstall), release, purge)
}
