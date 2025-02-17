// Code generated by MockGen. DO NOT EDIT.
// Source: C:/Users/Dinesh/go/src/BillingGo/CustomerMS/handler/handler.go
//
// Generated by this command:
//
//	mockgen -source=C:/Users/Dinesh/go/src/BillingGo/CustomerMS/handler/handler.go -destination=C:/Users/Dinesh/go/src/BillingGo/CustomerMS/mocks/handlerMocks.go -package=mocks
//
// Package mocks is a generated GoMock package.
package mocks

import (
	http "net/http"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockBillHandler is a mock of BillHandler interface.
type MockBillHandler struct {
	ctrl     *gomock.Controller
	recorder *MockBillHandlerMockRecorder
}

// MockBillHandlerMockRecorder is the mock recorder for MockBillHandler.
type MockBillHandlerMockRecorder struct {
	mock *MockBillHandler
}

// NewMockBillHandler creates a new mock instance.
func NewMockBillHandler(ctrl *gomock.Controller) *MockBillHandler {
	mock := &MockBillHandler{ctrl: ctrl}
	mock.recorder = &MockBillHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBillHandler) EXPECT() *MockBillHandlerMockRecorder {
	return m.recorder
}

// DELETE mocks base method.
func (m *MockBillHandler) DELETE(response http.ResponseWriter, req *http.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DELETE", response, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// DELETE indicates an expected call of DELETE.
func (mr *MockBillHandlerMockRecorder) DELETE(response, req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DELETE", reflect.TypeOf((*MockBillHandler)(nil).DELETE), response, req)
}

// GET mocks base method.
func (m *MockBillHandler) GET(response http.ResponseWriter, req *http.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GET", response, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// GET indicates an expected call of GET.
func (mr *MockBillHandlerMockRecorder) GET(response, req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GET", reflect.TypeOf((*MockBillHandler)(nil).GET), response, req)
}

// POST mocks base method.
func (m *MockBillHandler) POST(response http.ResponseWriter, req *http.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "POST", response, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// POST indicates an expected call of POST.
func (mr *MockBillHandlerMockRecorder) POST(response, req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "POST", reflect.TypeOf((*MockBillHandler)(nil).POST), response, req)
}

// PUT mocks base method.
func (m *MockBillHandler) PUT(response http.ResponseWriter, req *http.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PUT", response, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// PUT indicates an expected call of PUT.
func (mr *MockBillHandlerMockRecorder) PUT(response, req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PUT", reflect.TypeOf((*MockBillHandler)(nil).PUT), response, req)
}
