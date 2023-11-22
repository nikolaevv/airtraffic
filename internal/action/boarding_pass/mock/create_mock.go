// Code generated by MockGen. DO NOT EDIT.
// Source: create.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/nikolaevv/airtraffic/internal/model"
)

// MockCreateAdaptor is a mock of CreateAdaptor interface.
type MockCreateAdaptor struct {
	ctrl     *gomock.Controller
	recorder *MockCreateAdaptorMockRecorder
}

// MockCreateAdaptorMockRecorder is the mock recorder for MockCreateAdaptor.
type MockCreateAdaptorMockRecorder struct {
	mock *MockCreateAdaptor
}

// NewMockCreateAdaptor creates a new mock instance.
func NewMockCreateAdaptor(ctrl *gomock.Controller) *MockCreateAdaptor {
	mock := &MockCreateAdaptor{ctrl: ctrl}
	mock.recorder = &MockCreateAdaptorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCreateAdaptor) EXPECT() *MockCreateAdaptorMockRecorder {
	return m.recorder
}

// CreateBoardingPass mocks base method.
func (m *MockCreateAdaptor) CreateBoardingPass(ctx context.Context, flightID, seatID int) (model.BoardingPass, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBoardingPass", ctx, flightID, seatID)
	ret0, _ := ret[0].(model.BoardingPass)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBoardingPass indicates an expected call of CreateBoardingPass.
func (mr *MockCreateAdaptorMockRecorder) CreateBoardingPass(ctx, flightID, seatID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBoardingPass", reflect.TypeOf((*MockCreateAdaptor)(nil).CreateBoardingPass), ctx, flightID, seatID)
}
