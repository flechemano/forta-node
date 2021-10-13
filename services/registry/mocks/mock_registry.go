// Code generated by MockGen. DO NOT EDIT.
// Source: services/registry/registry.go

// Package mock_registry is a generated GoMock package.
package mock_registry

import (
	context "context"
	big "math/big"
	reflect "reflect"

	ethereum "github.com/ethereum/go-ethereum"
	types "github.com/ethereum/go-ethereum/core/types"
	domain "github.com/forta-protocol/forta-node/domain"
	regtypes "github.com/forta-protocol/forta-node/services/registry/regtypes"
	gomock "github.com/golang/mock/gomock"
)

// MockIPFSClient is a mock of IPFSClient interface.
type MockIPFSClient struct {
	ctrl     *gomock.Controller
	recorder *MockIPFSClientMockRecorder
}

// MockIPFSClientMockRecorder is the mock recorder for MockIPFSClient.
type MockIPFSClientMockRecorder struct {
	mock *MockIPFSClient
}

// NewMockIPFSClient creates a new mock instance.
func NewMockIPFSClient(ctrl *gomock.Controller) *MockIPFSClient {
	mock := &MockIPFSClient{ctrl: ctrl}
	mock.recorder = &MockIPFSClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIPFSClient) EXPECT() *MockIPFSClientMockRecorder {
	return m.recorder
}

// GetAgentFile mocks base method.
func (m *MockIPFSClient) GetAgentFile(cid string) (*regtypes.AgentFile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAgentFile", cid)
	ret0, _ := ret[0].(*regtypes.AgentFile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAgentFile indicates an expected call of GetAgentFile.
func (mr *MockIPFSClientMockRecorder) GetAgentFile(cid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAgentFile", reflect.TypeOf((*MockIPFSClient)(nil).GetAgentFile), cid)
}

// MockEthClient is a mock of EthClient interface.
type MockEthClient struct {
	ctrl     *gomock.Controller
	recorder *MockEthClientMockRecorder
}

// MockEthClientMockRecorder is the mock recorder for MockEthClient.
type MockEthClientMockRecorder struct {
	mock *MockEthClient
}

// NewMockEthClient creates a new mock instance.
func NewMockEthClient(ctrl *gomock.Controller) *MockEthClient {
	mock := &MockEthClient{ctrl: ctrl}
	mock.recorder = &MockEthClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEthClient) EXPECT() *MockEthClientMockRecorder {
	return m.recorder
}

// BlockByHash mocks base method.
func (m *MockEthClient) BlockByHash(ctx context.Context, hash string) (*domain.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockByHash", ctx, hash)
	ret0, _ := ret[0].(*domain.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BlockByHash indicates an expected call of BlockByHash.
func (mr *MockEthClientMockRecorder) BlockByHash(ctx, hash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockByHash", reflect.TypeOf((*MockEthClient)(nil).BlockByHash), ctx, hash)
}

// BlockByNumber mocks base method.
func (m *MockEthClient) BlockByNumber(ctx context.Context, number *big.Int) (*domain.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockByNumber", ctx, number)
	ret0, _ := ret[0].(*domain.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BlockByNumber indicates an expected call of BlockByNumber.
func (mr *MockEthClientMockRecorder) BlockByNumber(ctx, number interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockByNumber", reflect.TypeOf((*MockEthClient)(nil).BlockByNumber), ctx, number)
}

// BlockNumber mocks base method.
func (m *MockEthClient) BlockNumber(ctx context.Context) (*big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockNumber", ctx)
	ret0, _ := ret[0].(*big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BlockNumber indicates an expected call of BlockNumber.
func (mr *MockEthClientMockRecorder) BlockNumber(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockNumber", reflect.TypeOf((*MockEthClient)(nil).BlockNumber), ctx)
}

// ChainID mocks base method.
func (m *MockEthClient) ChainID(ctx context.Context) (*big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChainID", ctx)
	ret0, _ := ret[0].(*big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChainID indicates an expected call of ChainID.
func (mr *MockEthClientMockRecorder) ChainID(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChainID", reflect.TypeOf((*MockEthClient)(nil).ChainID), ctx)
}

// Close mocks base method.
func (m *MockEthClient) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockEthClientMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockEthClient)(nil).Close))
}

// GetLogs mocks base method.
func (m *MockEthClient) GetLogs(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLogs", ctx, q)
	ret0, _ := ret[0].([]types.Log)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLogs indicates an expected call of GetLogs.
func (mr *MockEthClientMockRecorder) GetLogs(ctx, q interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLogs", reflect.TypeOf((*MockEthClient)(nil).GetLogs), ctx, q)
}

// TraceBlock mocks base method.
func (m *MockEthClient) TraceBlock(ctx context.Context, number *big.Int) ([]domain.Trace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TraceBlock", ctx, number)
	ret0, _ := ret[0].([]domain.Trace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TraceBlock indicates an expected call of TraceBlock.
func (mr *MockEthClientMockRecorder) TraceBlock(ctx, number interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TraceBlock", reflect.TypeOf((*MockEthClient)(nil).TraceBlock), ctx, number)
}

// TransactionReceipt mocks base method.
func (m *MockEthClient) TransactionReceipt(ctx context.Context, txHash string) (*domain.TransactionReceipt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransactionReceipt", ctx, txHash)
	ret0, _ := ret[0].(*domain.TransactionReceipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TransactionReceipt indicates an expected call of TransactionReceipt.
func (mr *MockEthClientMockRecorder) TransactionReceipt(ctx, txHash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransactionReceipt", reflect.TypeOf((*MockEthClient)(nil).TransactionReceipt), ctx, txHash)
}
