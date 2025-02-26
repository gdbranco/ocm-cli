// Code generated by MockGen. DO NOT EDIT.
// Source: cluster.go
//
// Generated by this command:
//
//	mockgen -source=cluster.go -package=ocm -destination=mock_cluster.go
//

// Package ocm is a generated GoMock package.
package ocm

import (
	reflect "reflect"

	v1 "github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1"
	gomock "go.uber.org/mock/gomock"
)

// MockCluster is a mock of Cluster interface.
type MockCluster struct {
	ctrl     *gomock.Controller
	recorder *MockClusterMockRecorder
}

// MockClusterMockRecorder is the mock recorder for MockCluster.
type MockClusterMockRecorder struct {
	mock *MockCluster
}

// NewMockCluster creates a new mock instance.
func NewMockCluster(ctrl *gomock.Controller) *MockCluster {
	mock := &MockCluster{ctrl: ctrl}
	mock.recorder = &MockClusterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCluster) EXPECT() *MockClusterMockRecorder {
	return m.recorder
}

// CcsEnabled mocks base method.
func (m *MockCluster) CcsEnabled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CcsEnabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// CcsEnabled indicates an expected call of CcsEnabled.
func (mr *MockClusterMockRecorder) CcsEnabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CcsEnabled", reflect.TypeOf((*MockCluster)(nil).CcsEnabled))
}

// CloudProviderId mocks base method.
func (m *MockCluster) CloudProviderId() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloudProviderId")
	ret0, _ := ret[0].(string)
	return ret0
}

// CloudProviderId indicates an expected call of CloudProviderId.
func (mr *MockClusterMockRecorder) CloudProviderId() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudProviderId", reflect.TypeOf((*MockCluster)(nil).CloudProviderId))
}

// Id mocks base method.
func (m *MockCluster) Id() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Id")
	ret0, _ := ret[0].(string)
	return ret0
}

// Id indicates an expected call of Id.
func (mr *MockClusterMockRecorder) Id() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Id", reflect.TypeOf((*MockCluster)(nil).Id))
}

// State mocks base method.
func (m *MockCluster) State() v1.ClusterState {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "State")
	ret0, _ := ret[0].(v1.ClusterState)
	return ret0
}

// State indicates an expected call of State.
func (mr *MockClusterMockRecorder) State() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "State", reflect.TypeOf((*MockCluster)(nil).State))
}
