// Code generated by MockGen. DO NOT EDIT.
// Source: .\pkg\thesaurus\thesaurus.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockThesaurus is a mock of Thesaurus interface.
type MockThesaurus struct {
	ctrl     *gomock.Controller
	recorder *MockThesaurusMockRecorder
}

// MockThesaurusMockRecorder is the mock recorder for MockThesaurus.
type MockThesaurusMockRecorder struct {
	mock *MockThesaurus
}

// NewMockThesaurus creates a new mock instance.
func NewMockThesaurus(ctrl *gomock.Controller) *MockThesaurus {
	mock := &MockThesaurus{ctrl: ctrl}
	mock.recorder = &MockThesaurusMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockThesaurus) EXPECT() *MockThesaurusMockRecorder {
	return m.recorder
}

// GetSynonyms mocks base method.
func (m *MockThesaurus) GetSynonyms(arg0 string) []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSynonyms", arg0)
	ret0, _ := ret[0].([]string)
	return ret0
}

// GetSynonyms indicates an expected call of GetSynonyms.
func (mr *MockThesaurusMockRecorder) GetSynonyms(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSynonyms", reflect.TypeOf((*MockThesaurus)(nil).GetSynonyms), arg0)
}

// MockOnlineThesaurus is a mock of OnlineThesaurus interface.
type MockOnlineThesaurus struct {
	ctrl     *gomock.Controller
	recorder *MockOnlineThesaurusMockRecorder
}

// MockOnlineThesaurusMockRecorder is the mock recorder for MockOnlineThesaurus.
type MockOnlineThesaurusMockRecorder struct {
	mock *MockOnlineThesaurus
}

// NewMockOnlineThesaurus creates a new mock instance.
func NewMockOnlineThesaurus(ctrl *gomock.Controller) *MockOnlineThesaurus {
	mock := &MockOnlineThesaurus{ctrl: ctrl}
	mock.recorder = &MockOnlineThesaurusMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOnlineThesaurus) EXPECT() *MockOnlineThesaurusMockRecorder {
	return m.recorder
}

// Query mocks base method.
func (m *MockOnlineThesaurus) Query(arg0 string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Query", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Query indicates an expected call of Query.
func (mr *MockOnlineThesaurusMockRecorder) Query(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockOnlineThesaurus)(nil).Query), arg0)
}

// MockHttpClient is a mock of HttpClient interface.
type MockHttpClient struct {
	ctrl     *gomock.Controller
	recorder *MockHttpClientMockRecorder
}

// MockHttpClientMockRecorder is the mock recorder for MockHttpClient.
type MockHttpClientMockRecorder struct {
	mock *MockHttpClient
}

// NewMockHttpClient creates a new mock instance.
func NewMockHttpClient(ctrl *gomock.Controller) *MockHttpClient {
	mock := &MockHttpClient{ctrl: ctrl}
	mock.recorder = &MockHttpClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHttpClient) EXPECT() *MockHttpClientMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockHttpClient) Get(arg0 string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockHttpClientMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockHttpClient)(nil).Get), arg0)
}
