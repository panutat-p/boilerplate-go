// Code generated by MockGen. DO NOT EDIT.
// Source: store.go
//
// Generated by this command:
//
//	mockgen -source=store.go -destination=mock_store/store.go -package=mock_store
//

// Package mock_store is a generated GoMock package.
package mock_store

import (
	model "boilerplate-go/internal/model"
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockIStore is a mock of IStore interface.
type MockIStore struct {
	ctrl     *gomock.Controller
	recorder *MockIStoreMockRecorder
	isgomock struct{}
}

// MockIStoreMockRecorder is the mock recorder for MockIStore.
type MockIStoreMockRecorder struct {
	mock *MockIStore
}

// NewMockIStore creates a new mock instance.
func NewMockIStore(ctrl *gomock.Controller) *MockIStore {
	mock := &MockIStore{ctrl: ctrl}
	mock.recorder = &MockIStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIStore) EXPECT() *MockIStoreMockRecorder {
	return m.recorder
}

// ReadFruitFile mocks base method.
func (m *MockIStore) ReadFruitFile(ctx context.Context) ([]model.Fruit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadFruitFile", ctx)
	ret0, _ := ret[0].([]model.Fruit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadFruitFile indicates an expected call of ReadFruitFile.
func (mr *MockIStoreMockRecorder) ReadFruitFile(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadFruitFile", reflect.TypeOf((*MockIStore)(nil).ReadFruitFile), ctx)
}
