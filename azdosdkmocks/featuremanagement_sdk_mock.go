// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/microsoft/azure-devops-go-api/azuredevops/v7/featuremanagement (interfaces: Client)

// Package azdosdkmocks is a generated GoMock package.
package azdosdkmocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	featuremanagement "github.com/microsoft/azure-devops-go-api/azuredevops/v7/featuremanagement"
)

// MockFeaturemanagementClient is a mock of Client interface.
type MockFeaturemanagementClient struct {
	ctrl     *gomock.Controller
	recorder *MockFeaturemanagementClientMockRecorder
}

// MockFeaturemanagementClientMockRecorder is the mock recorder for MockFeaturemanagementClient.
type MockFeaturemanagementClientMockRecorder struct {
	mock *MockFeaturemanagementClient
}

// NewMockFeaturemanagementClient creates a new mock instance.
func NewMockFeaturemanagementClient(ctrl *gomock.Controller) *MockFeaturemanagementClient {
	mock := &MockFeaturemanagementClient{ctrl: ctrl}
	mock.recorder = &MockFeaturemanagementClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFeaturemanagementClient) EXPECT() *MockFeaturemanagementClientMockRecorder {
	return m.recorder
}

// GetFeature mocks base method.
func (m *MockFeaturemanagementClient) GetFeature(arg0 context.Context, arg1 featuremanagement.GetFeatureArgs) (*featuremanagement.ContributedFeature, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFeature", arg0, arg1)
	ret0, _ := ret[0].(*featuremanagement.ContributedFeature)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFeature indicates an expected call of GetFeature.
func (mr *MockFeaturemanagementClientMockRecorder) GetFeature(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFeature", reflect.TypeOf((*MockFeaturemanagementClient)(nil).GetFeature), arg0, arg1)
}

// GetFeatureState mocks base method.
func (m *MockFeaturemanagementClient) GetFeatureState(arg0 context.Context, arg1 featuremanagement.GetFeatureStateArgs) (*featuremanagement.ContributedFeatureState, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFeatureState", arg0, arg1)
	ret0, _ := ret[0].(*featuremanagement.ContributedFeatureState)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFeatureState indicates an expected call of GetFeatureState.
func (mr *MockFeaturemanagementClientMockRecorder) GetFeatureState(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFeatureState", reflect.TypeOf((*MockFeaturemanagementClient)(nil).GetFeatureState), arg0, arg1)
}

// GetFeatureStateForScope mocks base method.
func (m *MockFeaturemanagementClient) GetFeatureStateForScope(arg0 context.Context, arg1 featuremanagement.GetFeatureStateForScopeArgs) (*featuremanagement.ContributedFeatureState, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFeatureStateForScope", arg0, arg1)
	ret0, _ := ret[0].(*featuremanagement.ContributedFeatureState)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFeatureStateForScope indicates an expected call of GetFeatureStateForScope.
func (mr *MockFeaturemanagementClientMockRecorder) GetFeatureStateForScope(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFeatureStateForScope", reflect.TypeOf((*MockFeaturemanagementClient)(nil).GetFeatureStateForScope), arg0, arg1)
}

// GetFeatures mocks base method.
func (m *MockFeaturemanagementClient) GetFeatures(arg0 context.Context, arg1 featuremanagement.GetFeaturesArgs) (*[]featuremanagement.ContributedFeature, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFeatures", arg0, arg1)
	ret0, _ := ret[0].(*[]featuremanagement.ContributedFeature)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFeatures indicates an expected call of GetFeatures.
func (mr *MockFeaturemanagementClientMockRecorder) GetFeatures(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFeatures", reflect.TypeOf((*MockFeaturemanagementClient)(nil).GetFeatures), arg0, arg1)
}

// QueryFeatureStates mocks base method.
func (m *MockFeaturemanagementClient) QueryFeatureStates(arg0 context.Context, arg1 featuremanagement.QueryFeatureStatesArgs) (*featuremanagement.ContributedFeatureStateQuery, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryFeatureStates", arg0, arg1)
	ret0, _ := ret[0].(*featuremanagement.ContributedFeatureStateQuery)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryFeatureStates indicates an expected call of QueryFeatureStates.
func (mr *MockFeaturemanagementClientMockRecorder) QueryFeatureStates(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryFeatureStates", reflect.TypeOf((*MockFeaturemanagementClient)(nil).QueryFeatureStates), arg0, arg1)
}

// QueryFeatureStatesForDefaultScope mocks base method.
func (m *MockFeaturemanagementClient) QueryFeatureStatesForDefaultScope(arg0 context.Context, arg1 featuremanagement.QueryFeatureStatesForDefaultScopeArgs) (*featuremanagement.ContributedFeatureStateQuery, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryFeatureStatesForDefaultScope", arg0, arg1)
	ret0, _ := ret[0].(*featuremanagement.ContributedFeatureStateQuery)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryFeatureStatesForDefaultScope indicates an expected call of QueryFeatureStatesForDefaultScope.
func (mr *MockFeaturemanagementClientMockRecorder) QueryFeatureStatesForDefaultScope(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryFeatureStatesForDefaultScope", reflect.TypeOf((*MockFeaturemanagementClient)(nil).QueryFeatureStatesForDefaultScope), arg0, arg1)
}

// QueryFeatureStatesForNamedScope mocks base method.
func (m *MockFeaturemanagementClient) QueryFeatureStatesForNamedScope(arg0 context.Context, arg1 featuremanagement.QueryFeatureStatesForNamedScopeArgs) (*featuremanagement.ContributedFeatureStateQuery, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryFeatureStatesForNamedScope", arg0, arg1)
	ret0, _ := ret[0].(*featuremanagement.ContributedFeatureStateQuery)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryFeatureStatesForNamedScope indicates an expected call of QueryFeatureStatesForNamedScope.
func (mr *MockFeaturemanagementClientMockRecorder) QueryFeatureStatesForNamedScope(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryFeatureStatesForNamedScope", reflect.TypeOf((*MockFeaturemanagementClient)(nil).QueryFeatureStatesForNamedScope), arg0, arg1)
}

// SetFeatureState mocks base method.
func (m *MockFeaturemanagementClient) SetFeatureState(arg0 context.Context, arg1 featuremanagement.SetFeatureStateArgs) (*featuremanagement.ContributedFeatureState, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetFeatureState", arg0, arg1)
	ret0, _ := ret[0].(*featuremanagement.ContributedFeatureState)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetFeatureState indicates an expected call of SetFeatureState.
func (mr *MockFeaturemanagementClientMockRecorder) SetFeatureState(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetFeatureState", reflect.TypeOf((*MockFeaturemanagementClient)(nil).SetFeatureState), arg0, arg1)
}

// SetFeatureStateForScope mocks base method.
func (m *MockFeaturemanagementClient) SetFeatureStateForScope(arg0 context.Context, arg1 featuremanagement.SetFeatureStateForScopeArgs) (*featuremanagement.ContributedFeatureState, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetFeatureStateForScope", arg0, arg1)
	ret0, _ := ret[0].(*featuremanagement.ContributedFeatureState)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetFeatureStateForScope indicates an expected call of SetFeatureStateForScope.
func (mr *MockFeaturemanagementClientMockRecorder) SetFeatureStateForScope(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetFeatureStateForScope", reflect.TypeOf((*MockFeaturemanagementClient)(nil).SetFeatureStateForScope), arg0, arg1)
}
