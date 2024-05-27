// Code generated by MockGen. DO NOT EDIT.
// Source: payment.go

// Package payment is a generated GoMock package.
package payment

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockState is a mock of State interface.
type MockState struct {
	ctrl     *gomock.Controller
	recorder *MockStateMockRecorder
}

// MockStateMockRecorder is the mock recorder for MockState.
type MockStateMockRecorder struct {
	mock *MockState
}

// NewMockState creates a new mock instance.
func NewMockState(ctrl *gomock.Controller) *MockState {
	mock := &MockState{ctrl: ctrl}
	mock.recorder = &MockStateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockState) EXPECT() *MockStateMockRecorder {
	return m.recorder
}

// PaymentIdB2B mocks base method.
func (m *MockState) PaymentIdB2B() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PaymentIdB2B")
	ret0, _ := ret[0].(int)
	return ret0
}

// PaymentIdB2B indicates an expected call of PaymentIdB2B.
func (mr *MockStateMockRecorder) PaymentIdB2B() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PaymentIdB2B", reflect.TypeOf((*MockState)(nil).PaymentIdB2B))
}

// PaymentIdB2C mocks base method.
func (m *MockState) PaymentIdB2C() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PaymentIdB2C")
	ret0, _ := ret[0].(int)
	return ret0
}

// PaymentIdB2C indicates an expected call of PaymentIdB2C.
func (mr *MockStateMockRecorder) PaymentIdB2C() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PaymentIdB2C", reflect.TypeOf((*MockState)(nil).PaymentIdB2C))
}

// SetPaymentIdB2B mocks base method.
func (m *MockState) SetPaymentIdB2B(id int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetPaymentIdB2B", id)
}

// SetPaymentIdB2B indicates an expected call of SetPaymentIdB2B.
func (mr *MockStateMockRecorder) SetPaymentIdB2B(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPaymentIdB2B", reflect.TypeOf((*MockState)(nil).SetPaymentIdB2B), id)
}

// SetPaymentIdB2C mocks base method.
func (m *MockState) SetPaymentIdB2C(id int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetPaymentIdB2C", id)
}

// SetPaymentIdB2C indicates an expected call of SetPaymentIdB2C.
func (mr *MockStateMockRecorder) SetPaymentIdB2C(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPaymentIdB2C", reflect.TypeOf((*MockState)(nil).SetPaymentIdB2C), id)
}
