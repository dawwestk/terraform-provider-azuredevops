// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/microsoft/azure-devops-go-api/azuredevops/v7/serviceendpoint (interfaces: Client)

// Package azdosdkmocks is a generated GoMock package.
package azdosdkmocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	serviceendpoint "github.com/microsoft/azure-devops-go-api/azuredevops/v7/serviceendpoint"
)

// MockServiceendpointClient is a mock of Client interface.
type MockServiceendpointClient struct {
	ctrl     *gomock.Controller
	recorder *MockServiceendpointClientMockRecorder
}

// MockServiceendpointClientMockRecorder is the mock recorder for MockServiceendpointClient.
type MockServiceendpointClientMockRecorder struct {
	mock *MockServiceendpointClient
}

// NewMockServiceendpointClient creates a new mock instance.
func NewMockServiceendpointClient(ctrl *gomock.Controller) *MockServiceendpointClient {
	mock := &MockServiceendpointClient{ctrl: ctrl}
	mock.recorder = &MockServiceendpointClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceendpointClient) EXPECT() *MockServiceendpointClientMockRecorder {
	return m.recorder
}

// CreateServiceEndpoint mocks base method.
func (m *MockServiceendpointClient) CreateServiceEndpoint(arg0 context.Context, arg1 serviceendpoint.CreateServiceEndpointArgs) (*serviceendpoint.ServiceEndpoint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateServiceEndpoint", arg0, arg1)
	ret0, _ := ret[0].(*serviceendpoint.ServiceEndpoint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateServiceEndpoint indicates an expected call of CreateServiceEndpoint.
func (mr *MockServiceendpointClientMockRecorder) CreateServiceEndpoint(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateServiceEndpoint", reflect.TypeOf((*MockServiceendpointClient)(nil).CreateServiceEndpoint), arg0, arg1)
}

// DeleteServiceEndpoint mocks base method.
func (m *MockServiceendpointClient) DeleteServiceEndpoint(arg0 context.Context, arg1 serviceendpoint.DeleteServiceEndpointArgs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteServiceEndpoint", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteServiceEndpoint indicates an expected call of DeleteServiceEndpoint.
func (mr *MockServiceendpointClientMockRecorder) DeleteServiceEndpoint(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteServiceEndpoint", reflect.TypeOf((*MockServiceendpointClient)(nil).DeleteServiceEndpoint), arg0, arg1)
}

// ExecuteServiceEndpointRequest mocks base method.
func (m *MockServiceendpointClient) ExecuteServiceEndpointRequest(arg0 context.Context, arg1 serviceendpoint.ExecuteServiceEndpointRequestArgs) (*serviceendpoint.ServiceEndpointRequestResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExecuteServiceEndpointRequest", arg0, arg1)
	ret0, _ := ret[0].(*serviceendpoint.ServiceEndpointRequestResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExecuteServiceEndpointRequest indicates an expected call of ExecuteServiceEndpointRequest.
func (mr *MockServiceendpointClientMockRecorder) ExecuteServiceEndpointRequest(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecuteServiceEndpointRequest", reflect.TypeOf((*MockServiceendpointClient)(nil).ExecuteServiceEndpointRequest), arg0, arg1)
}

// GetServiceEndpointDetails mocks base method.
func (m *MockServiceendpointClient) GetServiceEndpointDetails(arg0 context.Context, arg1 serviceendpoint.GetServiceEndpointDetailsArgs) (*serviceendpoint.ServiceEndpoint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServiceEndpointDetails", arg0, arg1)
	ret0, _ := ret[0].(*serviceendpoint.ServiceEndpoint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetServiceEndpointDetails indicates an expected call of GetServiceEndpointDetails.
func (mr *MockServiceendpointClientMockRecorder) GetServiceEndpointDetails(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServiceEndpointDetails", reflect.TypeOf((*MockServiceendpointClient)(nil).GetServiceEndpointDetails), arg0, arg1)
}

// GetServiceEndpointExecutionRecords mocks base method.
func (m *MockServiceendpointClient) GetServiceEndpointExecutionRecords(arg0 context.Context, arg1 serviceendpoint.GetServiceEndpointExecutionRecordsArgs) (*serviceendpoint.GetServiceEndpointExecutionRecordsResponseValue, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServiceEndpointExecutionRecords", arg0, arg1)
	ret0, _ := ret[0].(*serviceendpoint.GetServiceEndpointExecutionRecordsResponseValue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetServiceEndpointExecutionRecords indicates an expected call of GetServiceEndpointExecutionRecords.
func (mr *MockServiceendpointClientMockRecorder) GetServiceEndpointExecutionRecords(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServiceEndpointExecutionRecords", reflect.TypeOf((*MockServiceendpointClient)(nil).GetServiceEndpointExecutionRecords), arg0, arg1)
}

// GetServiceEndpointTypes mocks base method.
func (m *MockServiceendpointClient) GetServiceEndpointTypes(arg0 context.Context, arg1 serviceendpoint.GetServiceEndpointTypesArgs) (*[]serviceendpoint.ServiceEndpointType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServiceEndpointTypes", arg0, arg1)
	ret0, _ := ret[0].(*[]serviceendpoint.ServiceEndpointType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetServiceEndpointTypes indicates an expected call of GetServiceEndpointTypes.
func (mr *MockServiceendpointClientMockRecorder) GetServiceEndpointTypes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServiceEndpointTypes", reflect.TypeOf((*MockServiceendpointClient)(nil).GetServiceEndpointTypes), arg0, arg1)
}

// GetServiceEndpoints mocks base method.
func (m *MockServiceendpointClient) GetServiceEndpoints(arg0 context.Context, arg1 serviceendpoint.GetServiceEndpointsArgs) (*[]serviceendpoint.ServiceEndpoint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServiceEndpoints", arg0, arg1)
	ret0, _ := ret[0].(*[]serviceendpoint.ServiceEndpoint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetServiceEndpoints indicates an expected call of GetServiceEndpoints.
func (mr *MockServiceendpointClientMockRecorder) GetServiceEndpoints(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServiceEndpoints", reflect.TypeOf((*MockServiceendpointClient)(nil).GetServiceEndpoints), arg0, arg1)
}

// GetServiceEndpointsByNames mocks base method.
func (m *MockServiceendpointClient) GetServiceEndpointsByNames(arg0 context.Context, arg1 serviceendpoint.GetServiceEndpointsByNamesArgs) (*[]serviceendpoint.ServiceEndpoint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServiceEndpointsByNames", arg0, arg1)
	ret0, _ := ret[0].(*[]serviceendpoint.ServiceEndpoint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetServiceEndpointsByNames indicates an expected call of GetServiceEndpointsByNames.
func (mr *MockServiceendpointClientMockRecorder) GetServiceEndpointsByNames(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServiceEndpointsByNames", reflect.TypeOf((*MockServiceendpointClient)(nil).GetServiceEndpointsByNames), arg0, arg1)
}

// GetServiceEndpointsWithRefreshedAuthentication mocks base method.
func (m *MockServiceendpointClient) GetServiceEndpointsWithRefreshedAuthentication(arg0 context.Context, arg1 serviceendpoint.GetServiceEndpointsWithRefreshedAuthenticationArgs) (*[]serviceendpoint.ServiceEndpoint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServiceEndpointsWithRefreshedAuthentication", arg0, arg1)
	ret0, _ := ret[0].(*[]serviceendpoint.ServiceEndpoint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetServiceEndpointsWithRefreshedAuthentication indicates an expected call of GetServiceEndpointsWithRefreshedAuthentication.
func (mr *MockServiceendpointClientMockRecorder) GetServiceEndpointsWithRefreshedAuthentication(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServiceEndpointsWithRefreshedAuthentication", reflect.TypeOf((*MockServiceendpointClient)(nil).GetServiceEndpointsWithRefreshedAuthentication), arg0, arg1)
}

// ShareServiceEndpoint mocks base method.
func (m *MockServiceendpointClient) ShareServiceEndpoint(arg0 context.Context, arg1 serviceendpoint.ShareServiceEndpointArgs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShareServiceEndpoint", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ShareServiceEndpoint indicates an expected call of ShareServiceEndpoint.
func (mr *MockServiceendpointClientMockRecorder) ShareServiceEndpoint(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShareServiceEndpoint", reflect.TypeOf((*MockServiceendpointClient)(nil).ShareServiceEndpoint), arg0, arg1)
}

// UpdateServiceEndpoint mocks base method.
func (m *MockServiceendpointClient) UpdateServiceEndpoint(arg0 context.Context, arg1 serviceendpoint.UpdateServiceEndpointArgs) (*serviceendpoint.ServiceEndpoint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateServiceEndpoint", arg0, arg1)
	ret0, _ := ret[0].(*serviceendpoint.ServiceEndpoint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateServiceEndpoint indicates an expected call of UpdateServiceEndpoint.
func (mr *MockServiceendpointClientMockRecorder) UpdateServiceEndpoint(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateServiceEndpoint", reflect.TypeOf((*MockServiceendpointClient)(nil).UpdateServiceEndpoint), arg0, arg1)
}

// UpdateServiceEndpoints mocks base method.
func (m *MockServiceendpointClient) UpdateServiceEndpoints(arg0 context.Context, arg1 serviceendpoint.UpdateServiceEndpointsArgs) (*[]serviceendpoint.ServiceEndpoint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateServiceEndpoints", arg0, arg1)
	ret0, _ := ret[0].(*[]serviceendpoint.ServiceEndpoint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateServiceEndpoints indicates an expected call of UpdateServiceEndpoints.
func (mr *MockServiceendpointClientMockRecorder) UpdateServiceEndpoints(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateServiceEndpoints", reflect.TypeOf((*MockServiceendpointClient)(nil).UpdateServiceEndpoints), arg0, arg1)
}
