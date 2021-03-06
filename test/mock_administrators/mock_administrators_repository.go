// Code generated by MockGen. DO NOT EDIT.
// Source: usecase/repository/administrators_repository.go

// Package mock_repository is a generated GoMock package.
package mock_administrators

import (
	domain "app/domain"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockAdministratorsRepository is a mock of AdministratorsRepository interface.
type MockAdministratorsRepository struct {
	ctrl     *gomock.Controller
	recorder *MockAdministratorsRepositoryMockRecorder
}

// MockAdministratorsRepositoryMockRecorder is the mock recorder for MockAdministratorsRepository.
type MockAdministratorsRepositoryMockRecorder struct {
	mock *MockAdministratorsRepository
}

// NewMockAdministratorsRepository creates a new mock instance.
func NewMockAdministratorsRepository(ctrl *gomock.Controller) *MockAdministratorsRepository {
	mock := &MockAdministratorsRepository{ctrl: ctrl}
	mock.recorder = &MockAdministratorsRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAdministratorsRepository) EXPECT() *MockAdministratorsRepositoryMockRecorder {
	return m.recorder
}

// FindByEmail mocks base method.
func (m *MockAdministratorsRepository) FindByEmail(email string) (domain.Administrator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByEmail", email)
	ret0, _ := ret[0].(domain.Administrator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByEmail indicates an expected call of FindByEmail.
func (mr *MockAdministratorsRepositoryMockRecorder) FindByEmail(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByEmail", reflect.TypeOf((*MockAdministratorsRepository)(nil).FindByEmail), email)
}
