// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/upbound/universal-crossplane/internal/clients/upbound (interfaces: Client)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	upbound "github.com/upbound/universal-crossplane/internal/clients/upbound"
)

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// GetGatewayCerts mocks base method
func (m *MockClient) GetGatewayCerts(arg0 string) (upbound.PublicCerts, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGatewayCerts", arg0)
	ret0, _ := ret[0].(upbound.PublicCerts)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGatewayCerts indicates an expected call of GetGatewayCerts
func (mr *MockClientMockRecorder) GetGatewayCerts(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGatewayCerts", reflect.TypeOf((*MockClient)(nil).GetGatewayCerts), arg0)
}
