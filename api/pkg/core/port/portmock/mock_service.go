// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/isutare412/bloated/api/pkg/core/port (interfaces: TodoService)

// Package portmock is a generated GoMock package.
package portmock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	ent "github.com/isutare412/bloated/api/pkg/core/ent"
)

// MockTodoService is a mock of TodoService interface.
type MockTodoService struct {
	ctrl     *gomock.Controller
	recorder *MockTodoServiceMockRecorder
}

// MockTodoServiceMockRecorder is the mock recorder for MockTodoService.
type MockTodoServiceMockRecorder struct {
	mock *MockTodoService
}

// NewMockTodoService creates a new mock instance.
func NewMockTodoService(ctrl *gomock.Controller) *MockTodoService {
	mock := &MockTodoService{ctrl: ctrl}
	mock.recorder = &MockTodoServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTodoService) EXPECT() *MockTodoServiceMockRecorder {
	return m.recorder
}

// AddTodo mocks base method.
func (m *MockTodoService) AddTodo(arg0 context.Context, arg1 *ent.Todo) (*ent.Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTodo", arg0, arg1)
	ret0, _ := ret[0].(*ent.Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddTodo indicates an expected call of AddTodo.
func (mr *MockTodoServiceMockRecorder) AddTodo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTodo", reflect.TypeOf((*MockTodoService)(nil).AddTodo), arg0, arg1)
}

// DeleteTodo mocks base method.
func (m *MockTodoService) DeleteTodo(arg0 context.Context, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTodo", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTodo indicates an expected call of DeleteTodo.
func (mr *MockTodoServiceMockRecorder) DeleteTodo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTodo", reflect.TypeOf((*MockTodoService)(nil).DeleteTodo), arg0, arg1)
}

// TodosOfUser mocks base method.
func (m *MockTodoService) TodosOfUser(arg0 context.Context, arg1 string) ([]*ent.Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TodosOfUser", arg0, arg1)
	ret0, _ := ret[0].([]*ent.Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TodosOfUser indicates an expected call of TodosOfUser.
func (mr *MockTodoServiceMockRecorder) TodosOfUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TodosOfUser", reflect.TypeOf((*MockTodoService)(nil).TodosOfUser), arg0, arg1)
}