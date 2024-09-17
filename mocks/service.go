// Package mock_service is a generated GoMock package.
package mock_service

import (
	entity "Bankirka/internal/entity"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockBankInt is a mock of BankInt interface.
type MockBankInt struct {
	ctrl     *gomock.Controller
	recorder *MockBankIntMockRecorder
}

// MockBankIntMockRecorder is the mock recorder for MockBankInt.
type MockBankIntMockRecorder struct {
	mock *MockBankInt
}

// NewMockBankInt creates a new mock instance.
func NewMockBankInt(ctrl *gomock.Controller) *MockBankInt {
	mock := &MockBankInt{ctrl: ctrl}
	mock.recorder = &MockBankIntMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBankInt) EXPECT() *MockBankIntMockRecorder {
	return m.recorder
}

// ChangeBalance mocks base method.
func (m *MockBankInt) ChangeBalance(arg0 int, arg1 entity.Difference) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeBalance", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangeBalance indicates an expected call of ChangeBalance.
func (mr *MockBankIntMockRecorder) ChangeBalance(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeBalance", reflect.TypeOf((*MockBankInt)(nil).ChangeBalance), arg0, arg1)
}

// CreatePerson mocks base method.
func (m *MockBankInt) CreatePerson(arg0 int, arg1 entity.Balance) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePerson", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreatePerson indicates an expected call of CreatePerson.
func (mr *MockBankIntMockRecorder) CreatePerson(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePerson", reflect.TypeOf((*MockBankInt)(nil).CreatePerson), arg0, arg1)
}

// ShowBalance mocks base method.
func (m *MockBankInt) ShowBalance(arg0 int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShowBalance", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ShowBalance indicates an expected call of ShowBalance.
func (mr *MockBankIntMockRecorder) ShowBalance(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShowBalance", reflect.TypeOf((*MockBankInt)(nil).ShowBalance), arg0)
}
