// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/isutare412/bloated/api/pkg/core/port (interfaces: TransactionManager,IPRepository,TodoRepository)

// Package portmock is a generated GoMock package.
package portmock

import (
	context "context"
	sql "database/sql"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	ent "github.com/isutare412/bloated/api/pkg/core/ent"
	transaction "github.com/isutare412/bloated/api/pkg/core/transaction"
)

// MockTransactionManager is a mock of TransactionManager interface.
type MockTransactionManager struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionManagerMockRecorder
}

// MockTransactionManagerMockRecorder is the mock recorder for MockTransactionManager.
type MockTransactionManagerMockRecorder struct {
	mock *MockTransactionManager
}

// NewMockTransactionManager creates a new mock instance.
func NewMockTransactionManager(ctrl *gomock.Controller) *MockTransactionManager {
	mock := &MockTransactionManager{ctrl: ctrl}
	mock.recorder = &MockTransactionManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransactionManager) EXPECT() *MockTransactionManagerMockRecorder {
	return m.recorder
}

// WithTx mocks base method.
func (m *MockTransactionManager) WithTx(arg0 context.Context) (transaction.Context, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithTx", arg0)
	ret0, _ := ret[0].(transaction.Context)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WithTx indicates an expected call of WithTx.
func (mr *MockTransactionManagerMockRecorder) WithTx(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithTx", reflect.TypeOf((*MockTransactionManager)(nil).WithTx), arg0)
}

// WithTxOption mocks base method.
func (m *MockTransactionManager) WithTxOption(arg0 context.Context, arg1 *sql.TxOptions) (transaction.Context, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithTxOption", arg0, arg1)
	ret0, _ := ret[0].(transaction.Context)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WithTxOption indicates an expected call of WithTxOption.
func (mr *MockTransactionManagerMockRecorder) WithTxOption(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithTxOption", reflect.TypeOf((*MockTransactionManager)(nil).WithTxOption), arg0, arg1)
}

// MockIPRepository is a mock of IPRepository interface.
type MockIPRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIPRepositoryMockRecorder
}

// MockIPRepositoryMockRecorder is the mock recorder for MockIPRepository.
type MockIPRepositoryMockRecorder struct {
	mock *MockIPRepository
}

// NewMockIPRepository creates a new mock instance.
func NewMockIPRepository(ctrl *gomock.Controller) *MockIPRepository {
	mock := &MockIPRepository{ctrl: ctrl}
	mock.recorder = &MockIPRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIPRepository) EXPECT() *MockIPRepositoryMockRecorder {
	return m.recorder
}

// CreateAll mocks base method.
func (m *MockIPRepository) CreateAll(arg0 context.Context, arg1 []*ent.BannedIP) ([]*ent.BannedIP, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAll", arg0, arg1)
	ret0, _ := ret[0].([]*ent.BannedIP)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAll indicates an expected call of CreateAll.
func (mr *MockIPRepositoryMockRecorder) CreateAll(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAll", reflect.TypeOf((*MockIPRepository)(nil).CreateAll), arg0, arg1)
}

// FindAll mocks base method.
func (m *MockIPRepository) FindAll(arg0 context.Context) ([]*ent.BannedIP, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll", arg0)
	ret0, _ := ret[0].([]*ent.BannedIP)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockIPRepositoryMockRecorder) FindAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockIPRepository)(nil).FindAll), arg0)
}

// MockTodoRepository is a mock of TodoRepository interface.
type MockTodoRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTodoRepositoryMockRecorder
}

// MockTodoRepositoryMockRecorder is the mock recorder for MockTodoRepository.
type MockTodoRepositoryMockRecorder struct {
	mock *MockTodoRepository
}

// NewMockTodoRepository creates a new mock instance.
func NewMockTodoRepository(ctrl *gomock.Controller) *MockTodoRepository {
	mock := &MockTodoRepository{ctrl: ctrl}
	mock.recorder = &MockTodoRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTodoRepository) EXPECT() *MockTodoRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockTodoRepository) Create(arg0 context.Context, arg1 *ent.Todo) (*ent.Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*ent.Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockTodoRepositoryMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTodoRepository)(nil).Create), arg0, arg1)
}

// DeleteByID mocks base method.
func (m *MockTodoRepository) DeleteByID(arg0 context.Context, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByID", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByID indicates an expected call of DeleteByID.
func (mr *MockTodoRepositoryMockRecorder) DeleteByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockTodoRepository)(nil).DeleteByID), arg0, arg1)
}

// FindByUserID mocks base method.
func (m *MockTodoRepository) FindByUserID(arg0 context.Context, arg1 string) ([]*ent.Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByUserID", arg0, arg1)
	ret0, _ := ret[0].([]*ent.Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByUserID indicates an expected call of FindByUserID.
func (mr *MockTodoRepositoryMockRecorder) FindByUserID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByUserID", reflect.TypeOf((*MockTodoRepository)(nil).FindByUserID), arg0, arg1)
}
