// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/microsoft/azure-devops-go-api/azuredevops/v7/test (interfaces: Client)

// Package azdosdkmocks is a generated GoMock package.
package azdosdkmocks

import (
	context "context"
	io "io"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	test "github.com/microsoft/azure-devops-go-api/azuredevops/v7/test"
)

// MockTestClient is a mock of Client interface.
type MockTestClient struct {
	ctrl     *gomock.Controller
	recorder *MockTestClientMockRecorder
}

// MockTestClientMockRecorder is the mock recorder for MockTestClient.
type MockTestClientMockRecorder struct {
	mock *MockTestClient
}

// NewMockTestClient creates a new mock instance.
func NewMockTestClient(ctrl *gomock.Controller) *MockTestClient {
	mock := &MockTestClient{ctrl: ctrl}
	mock.recorder = &MockTestClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTestClient) EXPECT() *MockTestClientMockRecorder {
	return m.recorder
}

// AddTestCasesToSuite mocks base method.
func (m *MockTestClient) AddTestCasesToSuite(arg0 context.Context, arg1 test.AddTestCasesToSuiteArgs) (*[]test.SuiteTestCase, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTestCasesToSuite", arg0, arg1)
	ret0, _ := ret[0].(*[]test.SuiteTestCase)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddTestCasesToSuite indicates an expected call of AddTestCasesToSuite.
func (mr *MockTestClientMockRecorder) AddTestCasesToSuite(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTestCasesToSuite", reflect.TypeOf((*MockTestClient)(nil).AddTestCasesToSuite), arg0, arg1)
}

// AddTestResultsToTestRun mocks base method.
func (m *MockTestClient) AddTestResultsToTestRun(arg0 context.Context, arg1 test.AddTestResultsToTestRunArgs) (*[]test.TestCaseResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTestResultsToTestRun", arg0, arg1)
	ret0, _ := ret[0].(*[]test.TestCaseResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddTestResultsToTestRun indicates an expected call of AddTestResultsToTestRun.
func (mr *MockTestClientMockRecorder) AddTestResultsToTestRun(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTestResultsToTestRun", reflect.TypeOf((*MockTestClient)(nil).AddTestResultsToTestRun), arg0, arg1)
}

// CreateTestResultAttachment mocks base method.
func (m *MockTestClient) CreateTestResultAttachment(arg0 context.Context, arg1 test.CreateTestResultAttachmentArgs) (*test.TestAttachmentReference, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTestResultAttachment", arg0, arg1)
	ret0, _ := ret[0].(*test.TestAttachmentReference)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTestResultAttachment indicates an expected call of CreateTestResultAttachment.
func (mr *MockTestClientMockRecorder) CreateTestResultAttachment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTestResultAttachment", reflect.TypeOf((*MockTestClient)(nil).CreateTestResultAttachment), arg0, arg1)
}

// CreateTestRun mocks base method.
func (m *MockTestClient) CreateTestRun(arg0 context.Context, arg1 test.CreateTestRunArgs) (*test.TestRun, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTestRun", arg0, arg1)
	ret0, _ := ret[0].(*test.TestRun)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTestRun indicates an expected call of CreateTestRun.
func (mr *MockTestClientMockRecorder) CreateTestRun(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTestRun", reflect.TypeOf((*MockTestClient)(nil).CreateTestRun), arg0, arg1)
}

// CreateTestRunAttachment mocks base method.
func (m *MockTestClient) CreateTestRunAttachment(arg0 context.Context, arg1 test.CreateTestRunAttachmentArgs) (*test.TestAttachmentReference, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTestRunAttachment", arg0, arg1)
	ret0, _ := ret[0].(*test.TestAttachmentReference)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTestRunAttachment indicates an expected call of CreateTestRunAttachment.
func (mr *MockTestClientMockRecorder) CreateTestRunAttachment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTestRunAttachment", reflect.TypeOf((*MockTestClient)(nil).CreateTestRunAttachment), arg0, arg1)
}

// CreateTestSession mocks base method.
func (m *MockTestClient) CreateTestSession(arg0 context.Context, arg1 test.CreateTestSessionArgs) (*test.TestSession, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTestSession", arg0, arg1)
	ret0, _ := ret[0].(*test.TestSession)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTestSession indicates an expected call of CreateTestSession.
func (mr *MockTestClientMockRecorder) CreateTestSession(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTestSession", reflect.TypeOf((*MockTestClient)(nil).CreateTestSession), arg0, arg1)
}

// CreateTestSubResultAttachment mocks base method.
func (m *MockTestClient) CreateTestSubResultAttachment(arg0 context.Context, arg1 test.CreateTestSubResultAttachmentArgs) (*test.TestAttachmentReference, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTestSubResultAttachment", arg0, arg1)
	ret0, _ := ret[0].(*test.TestAttachmentReference)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTestSubResultAttachment indicates an expected call of CreateTestSubResultAttachment.
func (mr *MockTestClientMockRecorder) CreateTestSubResultAttachment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTestSubResultAttachment", reflect.TypeOf((*MockTestClient)(nil).CreateTestSubResultAttachment), arg0, arg1)
}

// DeleteTestCase mocks base method.
func (m *MockTestClient) DeleteTestCase(arg0 context.Context, arg1 test.DeleteTestCaseArgs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTestCase", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTestCase indicates an expected call of DeleteTestCase.
func (mr *MockTestClientMockRecorder) DeleteTestCase(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTestCase", reflect.TypeOf((*MockTestClient)(nil).DeleteTestCase), arg0, arg1)
}

// DeleteTestRun mocks base method.
func (m *MockTestClient) DeleteTestRun(arg0 context.Context, arg1 test.DeleteTestRunArgs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTestRun", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTestRun indicates an expected call of DeleteTestRun.
func (mr *MockTestClientMockRecorder) DeleteTestRun(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTestRun", reflect.TypeOf((*MockTestClient)(nil).DeleteTestRun), arg0, arg1)
}

// GetBuildCodeCoverage mocks base method.
func (m *MockTestClient) GetBuildCodeCoverage(arg0 context.Context, arg1 test.GetBuildCodeCoverageArgs) (*[]test.BuildCoverage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBuildCodeCoverage", arg0, arg1)
	ret0, _ := ret[0].(*[]test.BuildCoverage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBuildCodeCoverage indicates an expected call of GetBuildCodeCoverage.
func (mr *MockTestClientMockRecorder) GetBuildCodeCoverage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBuildCodeCoverage", reflect.TypeOf((*MockTestClient)(nil).GetBuildCodeCoverage), arg0, arg1)
}

// GetPoint mocks base method.
func (m *MockTestClient) GetPoint(arg0 context.Context, arg1 test.GetPointArgs) (*test.TestPoint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPoint", arg0, arg1)
	ret0, _ := ret[0].(*test.TestPoint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPoint indicates an expected call of GetPoint.
func (mr *MockTestClientMockRecorder) GetPoint(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPoint", reflect.TypeOf((*MockTestClient)(nil).GetPoint), arg0, arg1)
}

// GetPoints mocks base method.
func (m *MockTestClient) GetPoints(arg0 context.Context, arg1 test.GetPointsArgs) (*[]test.TestPoint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPoints", arg0, arg1)
	ret0, _ := ret[0].(*[]test.TestPoint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPoints indicates an expected call of GetPoints.
func (mr *MockTestClientMockRecorder) GetPoints(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPoints", reflect.TypeOf((*MockTestClient)(nil).GetPoints), arg0, arg1)
}

// GetPointsByQuery mocks base method.
func (m *MockTestClient) GetPointsByQuery(arg0 context.Context, arg1 test.GetPointsByQueryArgs) (*test.TestPointsQuery, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPointsByQuery", arg0, arg1)
	ret0, _ := ret[0].(*test.TestPointsQuery)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPointsByQuery indicates an expected call of GetPointsByQuery.
func (mr *MockTestClientMockRecorder) GetPointsByQuery(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPointsByQuery", reflect.TypeOf((*MockTestClient)(nil).GetPointsByQuery), arg0, arg1)
}

// GetResultRetentionSettings mocks base method.
func (m *MockTestClient) GetResultRetentionSettings(arg0 context.Context, arg1 test.GetResultRetentionSettingsArgs) (*test.ResultRetentionSettings, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResultRetentionSettings", arg0, arg1)
	ret0, _ := ret[0].(*test.ResultRetentionSettings)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetResultRetentionSettings indicates an expected call of GetResultRetentionSettings.
func (mr *MockTestClientMockRecorder) GetResultRetentionSettings(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResultRetentionSettings", reflect.TypeOf((*MockTestClient)(nil).GetResultRetentionSettings), arg0, arg1)
}

// GetTestCaseById mocks base method.
func (m *MockTestClient) GetTestCaseById(arg0 context.Context, arg1 test.GetTestCaseByIdArgs) (*test.SuiteTestCase, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTestCaseById", arg0, arg1)
	ret0, _ := ret[0].(*test.SuiteTestCase)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTestCaseById indicates an expected call of GetTestCaseById.
func (mr *MockTestClientMockRecorder) GetTestCaseById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTestCaseById", reflect.TypeOf((*MockTestClient)(nil).GetTestCaseById), arg0, arg1)
}

// GetTestCases mocks base method.
func (m *MockTestClient) GetTestCases(arg0 context.Context, arg1 test.GetTestCasesArgs) (*[]test.SuiteTestCase, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTestCases", arg0, arg1)
	ret0, _ := ret[0].(*[]test.SuiteTestCase)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTestCases indicates an expected call of GetTestCases.
func (mr *MockTestClientMockRecorder) GetTestCases(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTestCases", reflect.TypeOf((*MockTestClient)(nil).GetTestCases), arg0, arg1)
}

// GetTestIteration mocks base method.
func (m *MockTestClient) GetTestIteration(arg0 context.Context, arg1 test.GetTestIterationArgs) (*test.TestIterationDetailsModel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTestIteration", arg0, arg1)
	ret0, _ := ret[0].(*test.TestIterationDetailsModel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTestIteration indicates an expected call of GetTestIteration.
func (mr *MockTestClientMockRecorder) GetTestIteration(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTestIteration", reflect.TypeOf((*MockTestClient)(nil).GetTestIteration), arg0, arg1)
}

// GetTestIterations mocks base method.
func (m *MockTestClient) GetTestIterations(arg0 context.Context, arg1 test.GetTestIterationsArgs) (*[]test.TestIterationDetailsModel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTestIterations", arg0, arg1)
	ret0, _ := ret[0].(*[]test.TestIterationDetailsModel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTestIterations indicates an expected call of GetTestIterations.
func (mr *MockTestClientMockRecorder) GetTestIterations(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTestIterations", reflect.TypeOf((*MockTestClient)(nil).GetTestIterations), arg0, arg1)
}

// GetTestResultAttachmentContent mocks base method.
func (m *MockTestClient) GetTestResultAttachmentContent(arg0 context.Context, arg1 test.GetTestResultAttachmentContentArgs) (io.ReadCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTestResultAttachmentContent", arg0, arg1)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTestResultAttachmentContent indicates an expected call of GetTestResultAttachmentContent.
func (mr *MockTestClientMockRecorder) GetTestResultAttachmentContent(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTestResultAttachmentContent", reflect.TypeOf((*MockTestClient)(nil).GetTestResultAttachmentContent), arg0, arg1)
}

// GetTestResultAttachmentZip mocks base method.
func (m *MockTestClient) GetTestResultAttachmentZip(arg0 context.Context, arg1 test.GetTestResultAttachmentZipArgs) (io.ReadCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTestResultAttachmentZip", arg0, arg1)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTestResultAttachmentZip indicates an expected call of GetTestResultAttachmentZip.
func (mr *MockTestClientMockRecorder) GetTestResultAttachmentZip(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTestResultAttachmentZip", reflect.TypeOf((*MockTestClient)(nil).GetTestResultAttachmentZip), arg0, arg1)
}

// GetTestResultAttachments mocks base method.
func (m *MockTestClient) GetTestResultAttachments(arg0 context.Context, arg1 test.GetTestResultAttachmentsArgs) (*[]test.TestAttachment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTestResultAttachments", arg0, arg1)
	ret0, _ := ret[0].(*[]test.TestAttachment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTestResultAttachments indicates an expected call of GetTestResultAttachments.
func (mr *MockTestClientMockRecorder) GetTestResultAttachments(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTestResultAttachments", reflect.TypeOf((*MockTestClient)(nil).GetTestResultAttachments), arg0, arg1)
}

// GetTestResultById mocks base method.
func (m *MockTestClient) GetTestResultById(arg0 context.Context, arg1 test.GetTestResultByIdArgs) (*test.TestCaseResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTestResultById", arg0, arg1)
	ret0, _ := ret[0].(*test.TestCaseResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTestResultById indicates an expected call of GetTestResultById.
func (mr *MockTestClientMockRecorder) GetTestResultById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTestResultById", reflect.TypeOf((*MockTestClient)(nil).GetTestResultById), arg0, arg1)
}

// GetTestResults mocks base method.
func (m *MockTestClient) GetTestResults(arg0 context.Context, arg1 test.GetTestResultsArgs) (*[]test.TestCaseResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTestResults", arg0, arg1)
	ret0, _ := ret[0].(*[]test.TestCaseResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTestResults indicates an expected call of GetTestResults.
func (mr *MockTestClientMockRecorder) GetTestResults(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTestResults", reflect.TypeOf((*MockTestClient)(nil).GetTestResults), arg0, arg1)
}

// GetTestRunAttachmentContent mocks base method.
func (m *MockTestClient) GetTestRunAttachmentContent(arg0 context.Context, arg1 test.GetTestRunAttachmentContentArgs) (io.ReadCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTestRunAttachmentContent", arg0, arg1)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTestRunAttachmentContent indicates an expected call of GetTestRunAttachmentContent.
func (mr *MockTestClientMockRecorder) GetTestRunAttachmentContent(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTestRunAttachmentContent", reflect.TypeOf((*MockTestClient)(nil).GetTestRunAttachmentContent), arg0, arg1)
}

// GetTestRunAttachmentZip mocks base method.
func (m *MockTestClient) GetTestRunAttachmentZip(arg0 context.Context, arg1 test.GetTestRunAttachmentZipArgs) (io.ReadCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTestRunAttachmentZip", arg0, arg1)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTestRunAttachmentZip indicates an expected call of GetTestRunAttachmentZip.
func (mr *MockTestClientMockRecorder) GetTestRunAttachmentZip(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTestRunAttachmentZip", reflect.TypeOf((*MockTestClient)(nil).GetTestRunAttachmentZip), arg0, arg1)
}

// GetTestRunAttachments mocks base method.
func (m *MockTestClient) GetTestRunAttachments(arg0 context.Context, arg1 test.GetTestRunAttachmentsArgs) (*[]test.TestAttachment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTestRunAttachments", arg0, arg1)
	ret0, _ := ret[0].(*[]test.TestAttachment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTestRunAttachments indicates an expected call of GetTestRunAttachments.
func (mr *MockTestClientMockRecorder) GetTestRunAttachments(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTestRunAttachments", reflect.TypeOf((*MockTestClient)(nil).GetTestRunAttachments), arg0, arg1)
}

// GetTestRunById mocks base method.
func (m *MockTestClient) GetTestRunById(arg0 context.Context, arg1 test.GetTestRunByIdArgs) (*test.TestRun, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTestRunById", arg0, arg1)
	ret0, _ := ret[0].(*test.TestRun)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTestRunById indicates an expected call of GetTestRunById.
func (mr *MockTestClientMockRecorder) GetTestRunById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTestRunById", reflect.TypeOf((*MockTestClient)(nil).GetTestRunById), arg0, arg1)
}

// GetTestRunCodeCoverage mocks base method.
func (m *MockTestClient) GetTestRunCodeCoverage(arg0 context.Context, arg1 test.GetTestRunCodeCoverageArgs) (*[]test.TestRunCoverage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTestRunCodeCoverage", arg0, arg1)
	ret0, _ := ret[0].(*[]test.TestRunCoverage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTestRunCodeCoverage indicates an expected call of GetTestRunCodeCoverage.
func (mr *MockTestClientMockRecorder) GetTestRunCodeCoverage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTestRunCodeCoverage", reflect.TypeOf((*MockTestClient)(nil).GetTestRunCodeCoverage), arg0, arg1)
}

// GetTestRunStatistics mocks base method.
func (m *MockTestClient) GetTestRunStatistics(arg0 context.Context, arg1 test.GetTestRunStatisticsArgs) (*test.TestRunStatistic, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTestRunStatistics", arg0, arg1)
	ret0, _ := ret[0].(*test.TestRunStatistic)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTestRunStatistics indicates an expected call of GetTestRunStatistics.
func (mr *MockTestClientMockRecorder) GetTestRunStatistics(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTestRunStatistics", reflect.TypeOf((*MockTestClient)(nil).GetTestRunStatistics), arg0, arg1)
}

// GetTestRuns mocks base method.
func (m *MockTestClient) GetTestRuns(arg0 context.Context, arg1 test.GetTestRunsArgs) (*[]test.TestRun, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTestRuns", arg0, arg1)
	ret0, _ := ret[0].(*[]test.TestRun)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTestRuns indicates an expected call of GetTestRuns.
func (mr *MockTestClientMockRecorder) GetTestRuns(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTestRuns", reflect.TypeOf((*MockTestClient)(nil).GetTestRuns), arg0, arg1)
}

// GetTestSessions mocks base method.
func (m *MockTestClient) GetTestSessions(arg0 context.Context, arg1 test.GetTestSessionsArgs) (*[]test.TestSession, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTestSessions", arg0, arg1)
	ret0, _ := ret[0].(*[]test.TestSession)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTestSessions indicates an expected call of GetTestSessions.
func (mr *MockTestClientMockRecorder) GetTestSessions(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTestSessions", reflect.TypeOf((*MockTestClient)(nil).GetTestSessions), arg0, arg1)
}

// GetTestSubResultAttachmentContent mocks base method.
func (m *MockTestClient) GetTestSubResultAttachmentContent(arg0 context.Context, arg1 test.GetTestSubResultAttachmentContentArgs) (io.ReadCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTestSubResultAttachmentContent", arg0, arg1)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTestSubResultAttachmentContent indicates an expected call of GetTestSubResultAttachmentContent.
func (mr *MockTestClientMockRecorder) GetTestSubResultAttachmentContent(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTestSubResultAttachmentContent", reflect.TypeOf((*MockTestClient)(nil).GetTestSubResultAttachmentContent), arg0, arg1)
}

// GetTestSubResultAttachmentZip mocks base method.
func (m *MockTestClient) GetTestSubResultAttachmentZip(arg0 context.Context, arg1 test.GetTestSubResultAttachmentZipArgs) (io.ReadCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTestSubResultAttachmentZip", arg0, arg1)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTestSubResultAttachmentZip indicates an expected call of GetTestSubResultAttachmentZip.
func (mr *MockTestClientMockRecorder) GetTestSubResultAttachmentZip(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTestSubResultAttachmentZip", reflect.TypeOf((*MockTestClient)(nil).GetTestSubResultAttachmentZip), arg0, arg1)
}

// GetTestSubResultAttachments mocks base method.
func (m *MockTestClient) GetTestSubResultAttachments(arg0 context.Context, arg1 test.GetTestSubResultAttachmentsArgs) (*[]test.TestAttachment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTestSubResultAttachments", arg0, arg1)
	ret0, _ := ret[0].(*[]test.TestAttachment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTestSubResultAttachments indicates an expected call of GetTestSubResultAttachments.
func (mr *MockTestClientMockRecorder) GetTestSubResultAttachments(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTestSubResultAttachments", reflect.TypeOf((*MockTestClient)(nil).GetTestSubResultAttachments), arg0, arg1)
}

// QueryTestHistory mocks base method.
func (m *MockTestClient) QueryTestHistory(arg0 context.Context, arg1 test.QueryTestHistoryArgs) (*test.TestHistoryQuery, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryTestHistory", arg0, arg1)
	ret0, _ := ret[0].(*test.TestHistoryQuery)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryTestHistory indicates an expected call of QueryTestHistory.
func (mr *MockTestClientMockRecorder) QueryTestHistory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryTestHistory", reflect.TypeOf((*MockTestClient)(nil).QueryTestHistory), arg0, arg1)
}

// QueryTestRuns mocks base method.
func (m *MockTestClient) QueryTestRuns(arg0 context.Context, arg1 test.QueryTestRunsArgs) (*test.QueryTestRunsResponseValue, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryTestRuns", arg0, arg1)
	ret0, _ := ret[0].(*test.QueryTestRunsResponseValue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryTestRuns indicates an expected call of QueryTestRuns.
func (mr *MockTestClientMockRecorder) QueryTestRuns(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryTestRuns", reflect.TypeOf((*MockTestClient)(nil).QueryTestRuns), arg0, arg1)
}

// RemoveTestCasesFromSuiteUrl mocks base method.
func (m *MockTestClient) RemoveTestCasesFromSuiteUrl(arg0 context.Context, arg1 test.RemoveTestCasesFromSuiteUrlArgs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveTestCasesFromSuiteUrl", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveTestCasesFromSuiteUrl indicates an expected call of RemoveTestCasesFromSuiteUrl.
func (mr *MockTestClientMockRecorder) RemoveTestCasesFromSuiteUrl(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveTestCasesFromSuiteUrl", reflect.TypeOf((*MockTestClient)(nil).RemoveTestCasesFromSuiteUrl), arg0, arg1)
}

// UpdateResultRetentionSettings mocks base method.
func (m *MockTestClient) UpdateResultRetentionSettings(arg0 context.Context, arg1 test.UpdateResultRetentionSettingsArgs) (*test.ResultRetentionSettings, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateResultRetentionSettings", arg0, arg1)
	ret0, _ := ret[0].(*test.ResultRetentionSettings)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateResultRetentionSettings indicates an expected call of UpdateResultRetentionSettings.
func (mr *MockTestClientMockRecorder) UpdateResultRetentionSettings(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateResultRetentionSettings", reflect.TypeOf((*MockTestClient)(nil).UpdateResultRetentionSettings), arg0, arg1)
}

// UpdateSuiteTestCases mocks base method.
func (m *MockTestClient) UpdateSuiteTestCases(arg0 context.Context, arg1 test.UpdateSuiteTestCasesArgs) (*[]test.SuiteTestCase, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSuiteTestCases", arg0, arg1)
	ret0, _ := ret[0].(*[]test.SuiteTestCase)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateSuiteTestCases indicates an expected call of UpdateSuiteTestCases.
func (mr *MockTestClientMockRecorder) UpdateSuiteTestCases(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSuiteTestCases", reflect.TypeOf((*MockTestClient)(nil).UpdateSuiteTestCases), arg0, arg1)
}

// UpdateTestPoints mocks base method.
func (m *MockTestClient) UpdateTestPoints(arg0 context.Context, arg1 test.UpdateTestPointsArgs) (*[]test.TestPoint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTestPoints", arg0, arg1)
	ret0, _ := ret[0].(*[]test.TestPoint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTestPoints indicates an expected call of UpdateTestPoints.
func (mr *MockTestClientMockRecorder) UpdateTestPoints(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTestPoints", reflect.TypeOf((*MockTestClient)(nil).UpdateTestPoints), arg0, arg1)
}

// UpdateTestResults mocks base method.
func (m *MockTestClient) UpdateTestResults(arg0 context.Context, arg1 test.UpdateTestResultsArgs) (*[]test.TestCaseResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTestResults", arg0, arg1)
	ret0, _ := ret[0].(*[]test.TestCaseResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTestResults indicates an expected call of UpdateTestResults.
func (mr *MockTestClientMockRecorder) UpdateTestResults(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTestResults", reflect.TypeOf((*MockTestClient)(nil).UpdateTestResults), arg0, arg1)
}

// UpdateTestRun mocks base method.
func (m *MockTestClient) UpdateTestRun(arg0 context.Context, arg1 test.UpdateTestRunArgs) (*test.TestRun, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTestRun", arg0, arg1)
	ret0, _ := ret[0].(*test.TestRun)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTestRun indicates an expected call of UpdateTestRun.
func (mr *MockTestClientMockRecorder) UpdateTestRun(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTestRun", reflect.TypeOf((*MockTestClient)(nil).UpdateTestRun), arg0, arg1)
}

// UpdateTestSession mocks base method.
func (m *MockTestClient) UpdateTestSession(arg0 context.Context, arg1 test.UpdateTestSessionArgs) (*test.TestSession, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTestSession", arg0, arg1)
	ret0, _ := ret[0].(*test.TestSession)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTestSession indicates an expected call of UpdateTestSession.
func (mr *MockTestClientMockRecorder) UpdateTestSession(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTestSession", reflect.TypeOf((*MockTestClient)(nil).UpdateTestSession), arg0, arg1)
}
