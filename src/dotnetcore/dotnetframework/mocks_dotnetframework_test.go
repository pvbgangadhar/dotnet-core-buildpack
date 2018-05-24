// Code generated by MockGen. DO NOT EDIT.
// Source: dotnetframework.go

// Package dotnetframework_test is a generated GoMock package.
package dotnetframework_test

import (
	libbuildpack "github.com/cloudfoundry/libbuildpack"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockInstaller is a mock of Installer interface
type MockInstaller struct {
	ctrl     *gomock.Controller
	recorder *MockInstallerMockRecorder
}

// MockInstallerMockRecorder is the mock recorder for MockInstaller
type MockInstallerMockRecorder struct {
	mock *MockInstaller
}

// NewMockInstaller creates a new mock instance
func NewMockInstaller(ctrl *gomock.Controller) *MockInstaller {
	mock := &MockInstaller{ctrl: ctrl}
	mock.recorder = &MockInstallerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInstaller) EXPECT() *MockInstallerMockRecorder {
	return m.recorder
}

// InstallDependency mocks base method
func (m *MockInstaller) InstallDependency(arg0 libbuildpack.Dependency, arg1 string) error {
	ret := m.ctrl.Call(m, "InstallDependency", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// InstallDependency indicates an expected call of InstallDependency
func (mr *MockInstallerMockRecorder) InstallDependency(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstallDependency", reflect.TypeOf((*MockInstaller)(nil).InstallDependency), arg0, arg1)
}
