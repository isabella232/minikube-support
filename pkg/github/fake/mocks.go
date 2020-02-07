// Code generated by MockGen. DO NOT EDIT.
// Source: client.go

// Package fake is a generated GoMock package.
package fake

import (
	io "io"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
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

// SetApiToken mocks base method
func (m *MockClient) SetApiToken(apiToken string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetApiToken", apiToken)
}

// SetApiToken indicates an expected call of SetApiToken
func (mr *MockClientMockRecorder) SetApiToken(apiToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetApiToken", reflect.TypeOf((*MockClient)(nil).SetApiToken), apiToken)
}

// GetLatestReleaseTag mocks base method
func (m *MockClient) GetLatestReleaseTag(org, repository string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLatestReleaseTag", org, repository)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLatestReleaseTag indicates an expected call of GetLatestReleaseTag
func (mr *MockClientMockRecorder) GetLatestReleaseTag(org, repository interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLatestReleaseTag", reflect.TypeOf((*MockClient)(nil).GetLatestReleaseTag), org, repository)
}

// DownloadReleaseAsset mocks base method
func (m *MockClient) DownloadReleaseAsset(org, repository, tag, assetName string) (io.ReadCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DownloadReleaseAsset", org, repository, tag, assetName)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DownloadReleaseAsset indicates an expected call of DownloadReleaseAsset
func (mr *MockClientMockRecorder) DownloadReleaseAsset(org, repository, tag, assetName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadReleaseAsset", reflect.TypeOf((*MockClient)(nil).DownloadReleaseAsset), org, repository, tag, assetName)
}
