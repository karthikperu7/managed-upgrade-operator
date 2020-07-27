// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/openshift/managed-upgrade-operator/pkg/validation (interfaces: Validator)

// Package validation is a generated GoMock package.
package validation

import (
	logr "github.com/go-logr/logr"
	gomock "github.com/golang/mock/gomock"
	v1 "github.com/openshift/api/config/v1"
	v1alpha1 "github.com/openshift/managed-upgrade-operator/pkg/apis/upgrade/v1alpha1"
	reflect "reflect"
)

// MockValidator is a mock of Validator interface
type MockValidator struct {
	ctrl     *gomock.Controller
	recorder *MockValidatorMockRecorder
}

// MockValidatorMockRecorder is the mock recorder for MockValidator
type MockValidatorMockRecorder struct {
	mock *MockValidator
}

// NewMockValidator creates a new mock instance
func NewMockValidator(ctrl *gomock.Controller) *MockValidator {
	mock := &MockValidator{ctrl: ctrl}
	mock.recorder = &MockValidatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockValidator) EXPECT() *MockValidatorMockRecorder {
	return m.recorder
}

// IsValidUpgradeConfig mocks base method
func (m *MockValidator) IsValidUpgradeConfig(arg0 *v1alpha1.UpgradeConfig, arg1 *v1.ClusterVersion, arg2 logr.Logger) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsValidUpgradeConfig", arg0, arg1, arg2)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsValidUpgradeConfig indicates an expected call of IsValidUpgradeConfig
func (mr *MockValidatorMockRecorder) IsValidUpgradeConfig(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsValidUpgradeConfig", reflect.TypeOf((*MockValidator)(nil).IsValidUpgradeConfig), arg0, arg1, arg2)
}