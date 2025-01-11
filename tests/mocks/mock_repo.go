// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/twergi/calculator/internal/app/usecases/calculator (interfaces: Repository)
//
// Generated by this command:
//
//	mockgen -destination ./mock_repo.go -package mocks -mock_names Repository=MockRepository github.com/twergi/calculator/internal/app/usecases/calculator Repository
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// GetLastResult mocks base method.
func (m *MockRepository) GetLastResult(arg0 context.Context) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLastResult", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLastResult indicates an expected call of GetLastResult.
func (mr *MockRepositoryMockRecorder) GetLastResult(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLastResult", reflect.TypeOf((*MockRepository)(nil).GetLastResult), arg0)
}

// SaveResult mocks base method.
func (m *MockRepository) SaveResult(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveResult", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveResult indicates an expected call of SaveResult.
func (mr *MockRepositoryMockRecorder) SaveResult(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveResult", reflect.TypeOf((*MockRepository)(nil).SaveResult), arg0, arg1)
}
