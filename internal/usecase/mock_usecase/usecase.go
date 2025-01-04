// Code generated by MockGen. DO NOT EDIT.
// Source: usecase.go
//
// Generated by this command:
//
//	mockgen -source=usecase.go -destination=mock_usecase/usecase.go -package=mock_usecase
//

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	model "boilerplate-go/internal/model"
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockIUseCase is a mock of IUseCase interface.
type MockIUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockIUseCaseMockRecorder
}

// MockIUseCaseMockRecorder is the mock recorder for MockIUseCase.
type MockIUseCaseMockRecorder struct {
	mock *MockIUseCase
}

// NewMockIUseCase creates a new mock instance.
func NewMockIUseCase(ctrl *gomock.Controller) *MockIUseCase {
	mock := &MockIUseCase{ctrl: ctrl}
	mock.recorder = &MockIUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUseCase) EXPECT() *MockIUseCaseMockRecorder {
	return m.recorder
}

// CheckFruits mocks base method.
func (m *MockIUseCase) CheckFruits(ctx context.Context, fruits []model.Fruit) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckFruits", ctx, fruits)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckFruits indicates an expected call of CheckFruits.
func (mr *MockIUseCaseMockRecorder) CheckFruits(ctx, fruits any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckFruits", reflect.TypeOf((*MockIUseCase)(nil).CheckFruits), ctx, fruits)
}

// GetFruits mocks base method.
func (m *MockIUseCase) GetFruits(ctx context.Context) ([]model.Fruit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFruits", ctx)
	ret0, _ := ret[0].([]model.Fruit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFruits indicates an expected call of GetFruits.
func (mr *MockIUseCaseMockRecorder) GetFruits(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFruits", reflect.TypeOf((*MockIUseCase)(nil).GetFruits), ctx)
}
