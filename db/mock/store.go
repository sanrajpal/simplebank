// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/techschool/simplebank/db/sqlc (interfaces: Store)
//
// Generated by this command:
//
//	mockgen -destination db/mock/store.go github.com/techschool/simplebank/db/sqlc Store
//

// Package mock_db is a generated GoMock package.
package mock_db

import (
	context "context"
	reflect "reflect"

	sqlc "github.com/techschool/simplebank/db/sqlc"
	gomock "go.uber.org/mock/gomock"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
	isgomock struct{}
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// AddAccountBalance mocks base method.
func (m *MockStore) AddAccountBalance(ctx context.Context, arg sqlc.AddAccountBalanceParams) (sqlc.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAccountBalance", ctx, arg)
	ret0, _ := ret[0].(sqlc.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddAccountBalance indicates an expected call of AddAccountBalance.
func (mr *MockStoreMockRecorder) AddAccountBalance(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAccountBalance", reflect.TypeOf((*MockStore)(nil).AddAccountBalance), ctx, arg)
}

// CreateAccount mocks base method.
func (m *MockStore) CreateAccount(ctx context.Context, arg sqlc.CreateAccountParams) (sqlc.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccount", ctx, arg)
	ret0, _ := ret[0].(sqlc.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAccount indicates an expected call of CreateAccount.
func (mr *MockStoreMockRecorder) CreateAccount(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccount", reflect.TypeOf((*MockStore)(nil).CreateAccount), ctx, arg)
}

// CreateEntry mocks base method.
func (m *MockStore) CreateEntry(ctx context.Context, arg sqlc.CreateEntryParams) (sqlc.Entry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEntry", ctx, arg)
	ret0, _ := ret[0].(sqlc.Entry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateEntry indicates an expected call of CreateEntry.
func (mr *MockStoreMockRecorder) CreateEntry(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEntry", reflect.TypeOf((*MockStore)(nil).CreateEntry), ctx, arg)
}

// CreateTransfer mocks base method.
func (m *MockStore) CreateTransfer(ctx context.Context, arg sqlc.CreateTransferParams) (sqlc.Transfer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTransfer", ctx, arg)
	ret0, _ := ret[0].(sqlc.Transfer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTransfer indicates an expected call of CreateTransfer.
func (mr *MockStoreMockRecorder) CreateTransfer(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTransfer", reflect.TypeOf((*MockStore)(nil).CreateTransfer), ctx, arg)
}

// DeleteAccount mocks base method.
func (m *MockStore) DeleteAccount(ctx context.Context, id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAccount", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAccount indicates an expected call of DeleteAccount.
func (mr *MockStoreMockRecorder) DeleteAccount(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAccount", reflect.TypeOf((*MockStore)(nil).DeleteAccount), ctx, id)
}

// GetAccount mocks base method.
func (m *MockStore) GetAccount(ctx context.Context, id int64) (sqlc.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccount", ctx, id)
	ret0, _ := ret[0].(sqlc.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccount indicates an expected call of GetAccount.
func (mr *MockStoreMockRecorder) GetAccount(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*MockStore)(nil).GetAccount), ctx, id)
}

// GetAccountForUpdate mocks base method.
func (m *MockStore) GetAccountForUpdate(ctx context.Context, id int64) (sqlc.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountForUpdate", ctx, id)
	ret0, _ := ret[0].(sqlc.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountForUpdate indicates an expected call of GetAccountForUpdate.
func (mr *MockStoreMockRecorder) GetAccountForUpdate(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountForUpdate", reflect.TypeOf((*MockStore)(nil).GetAccountForUpdate), ctx, id)
}

// GetEntry mocks base method.
func (m *MockStore) GetEntry(ctx context.Context, id int64) (sqlc.Entry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEntry", ctx, id)
	ret0, _ := ret[0].(sqlc.Entry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEntry indicates an expected call of GetEntry.
func (mr *MockStoreMockRecorder) GetEntry(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEntry", reflect.TypeOf((*MockStore)(nil).GetEntry), ctx, id)
}

// GetTransfer mocks base method.
func (m *MockStore) GetTransfer(ctx context.Context, id int64) (sqlc.Transfer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransfer", ctx, id)
	ret0, _ := ret[0].(sqlc.Transfer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransfer indicates an expected call of GetTransfer.
func (mr *MockStoreMockRecorder) GetTransfer(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransfer", reflect.TypeOf((*MockStore)(nil).GetTransfer), ctx, id)
}

// ListAccount mocks base method.
func (m *MockStore) ListAccount(ctx context.Context, arg sqlc.ListAccountParams) ([]sqlc.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAccount", ctx, arg)
	ret0, _ := ret[0].([]sqlc.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAccount indicates an expected call of ListAccount.
func (mr *MockStoreMockRecorder) ListAccount(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAccount", reflect.TypeOf((*MockStore)(nil).ListAccount), ctx, arg)
}

// ListEntry mocks base method.
func (m *MockStore) ListEntry(ctx context.Context, arg sqlc.ListEntryParams) ([]sqlc.Entry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListEntry", ctx, arg)
	ret0, _ := ret[0].([]sqlc.Entry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListEntry indicates an expected call of ListEntry.
func (mr *MockStoreMockRecorder) ListEntry(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEntry", reflect.TypeOf((*MockStore)(nil).ListEntry), ctx, arg)
}

// ListTransfers mocks base method.
func (m *MockStore) ListTransfers(ctx context.Context, arg sqlc.ListTransfersParams) ([]sqlc.Transfer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTransfers", ctx, arg)
	ret0, _ := ret[0].([]sqlc.Transfer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTransfers indicates an expected call of ListTransfers.
func (mr *MockStoreMockRecorder) ListTransfers(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTransfers", reflect.TypeOf((*MockStore)(nil).ListTransfers), ctx, arg)
}

// TransferTx mocks base method.
func (m *MockStore) TransferTx(ctx context.Context, arg sqlc.TransferTxParams) (sqlc.TransferTxResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransferTx", ctx, arg)
	ret0, _ := ret[0].(sqlc.TransferTxResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TransferTx indicates an expected call of TransferTx.
func (mr *MockStoreMockRecorder) TransferTx(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransferTx", reflect.TypeOf((*MockStore)(nil).TransferTx), ctx, arg)
}

// UpdateAccount mocks base method.
func (m *MockStore) UpdateAccount(ctx context.Context, arg sqlc.UpdateAccountParams) (sqlc.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAccount", ctx, arg)
	ret0, _ := ret[0].(sqlc.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAccount indicates an expected call of UpdateAccount.
func (mr *MockStoreMockRecorder) UpdateAccount(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAccount", reflect.TypeOf((*MockStore)(nil).UpdateAccount), ctx, arg)
}
