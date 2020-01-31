// Code generated by MockGen. DO NOT EDIT.
// Source: customer.go

// Package mocks is a generated GoMock package.
package mocks

import (
	entity "github.com/YukihiroMaekawa/go-lesson/internal/domain/entity"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// Mockdber is a mock of dber interface
type Mockdber struct {
	ctrl     *gomock.Controller
	recorder *MockdberMockRecorder
}

// MockdberMockRecorder is the mock recorder for Mockdber
type MockdberMockRecorder struct {
	mock *Mockdber
}

// NewMockdber creates a new mock instance
func NewMockdber(ctrl *gomock.Controller) *Mockdber {
	mock := &Mockdber{ctrl: ctrl}
	mock.recorder = &MockdberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *Mockdber) EXPECT() *MockdberMockRecorder {
	return m.recorder
}

// CreateCustomer mocks base method
func (m *Mockdber) CreateCustomer(request *entity.RegisterCustomerRequest) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCustomer", request)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCustomer indicates an expected call of CreateCustomer
func (mr *MockdberMockRecorder) CreateCustomer(request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCustomer", reflect.TypeOf((*Mockdber)(nil).CreateCustomer), request)
}
