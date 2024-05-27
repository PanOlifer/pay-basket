// Code generated by MockGen. DO NOT EDIT.
// Source: preorder.go

// Package resolver is a generated GoMock package.
package resolver

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	catalog_types "go.citilink.cloud/catalog_types"
	preorder "go.citilink.cloud/order/internal/preorder"
)

// MockPreOrderList is a mock of PreOrderList interface.
type MockPreOrderList struct {
	ctrl     *gomock.Controller
	recorder *MockPreOrderListMockRecorder
}

// MockPreOrderListMockRecorder is the mock recorder for MockPreOrderList.
type MockPreOrderListMockRecorder struct {
	mock *MockPreOrderList
}

// NewMockPreOrderList creates a new mock instance.
func NewMockPreOrderList(ctrl *gomock.Controller) *MockPreOrderList {
	mock := &MockPreOrderList{ctrl: ctrl}
	mock.recorder = &MockPreOrderListMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPreOrderList) EXPECT() *MockPreOrderListMockRecorder {
	return m.recorder
}

// FilterByProductIds mocks base method.
func (m *MockPreOrderList) FilterByProductIds(ctx context.Context, productIds ...catalog_types.ProductId) (preorder.PreOrders, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range productIds {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FilterByProductIds", varargs...)
	ret0, _ := ret[0].(preorder.PreOrders)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FilterByProductIds indicates an expected call of FilterByProductIds.
func (mr *MockPreOrderListMockRecorder) FilterByProductIds(ctx interface{}, productIds ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, productIds...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilterByProductIds", reflect.TypeOf((*MockPreOrderList)(nil).FilterByProductIds), varargs...)
}
