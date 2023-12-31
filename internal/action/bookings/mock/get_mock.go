// Code generated by MockGen. DO NOT EDIT.
// Source: get.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/nikolaevv/airtraffic/internal/model"
)

// MockGetAdaptor is a mock of GetAdaptor interface.
type MockGetAdaptor struct {
	ctrl     *gomock.Controller
	recorder *MockGetAdaptorMockRecorder
}

// MockGetAdaptorMockRecorder is the mock recorder for MockGetAdaptor.
type MockGetAdaptorMockRecorder struct {
	mock *MockGetAdaptor
}

// NewMockGetAdaptor creates a new mock instance.
func NewMockGetAdaptor(ctrl *gomock.Controller) *MockGetAdaptor {
	mock := &MockGetAdaptor{ctrl: ctrl}
	mock.recorder = &MockGetAdaptorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGetAdaptor) EXPECT() *MockGetAdaptorMockRecorder {
	return m.recorder
}

// GetBooking mocks base method.
func (m *MockGetAdaptor) GetBooking(ctx context.Context, id int) (model.Booking, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBooking", ctx, id)
	ret0, _ := ret[0].(model.Booking)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBooking indicates an expected call of GetBooking.
func (mr *MockGetAdaptorMockRecorder) GetBooking(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBooking", reflect.TypeOf((*MockGetAdaptor)(nil).GetBooking), ctx, id)
}
