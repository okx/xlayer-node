// Code generated by mockery v2.32.0. DO NOT EDIT.

package sequencer

import (
	context "context"
	big "math/big"

	common "github.com/ethereum/go-ethereum/common"

	mock "github.com/stretchr/testify/mock"

	pgx "github.com/jackc/pgx/v4"

	pool "github.com/0xPolygonHermez/zkevm-node/pool"

	state "github.com/0xPolygonHermez/zkevm-node/state"

	time "time"

	types "github.com/ethereum/go-ethereum/core/types"
)

// DbManagerMock is an autogenerated mock type for the dbManagerInterface type
type DbManagerMock struct {
	mock.Mock
}

// BeginStateTransaction provides a mock function with given fields: ctx
func (_m *DbManagerMock) BeginStateTransaction(ctx context.Context) (pgx.Tx, error) {
	ret := _m.Called(ctx)

	var r0 pgx.Tx
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (pgx.Tx, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) pgx.Tx); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(pgx.Tx)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CloseBatch provides a mock function with given fields: ctx, params
func (_m *DbManagerMock) CloseBatch(ctx context.Context, params ClosingBatchParameters) error {
	ret := _m.Called(ctx, params)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, ClosingBatchParameters) error); ok {
		r0 = rf(ctx, params)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CountReorgs provides a mock function with given fields: ctx, dbTx
func (_m *DbManagerMock) CountReorgs(ctx context.Context, dbTx pgx.Tx) (uint64, error) {
	ret := _m.Called(ctx, dbTx)

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) (uint64, error)); ok {
		return rf(ctx, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) uint64); ok {
		r0 = rf(ctx, dbTx)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, pgx.Tx) error); ok {
		r1 = rf(ctx, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateFirstBatch provides a mock function with given fields: ctx, l2coinbase
func (_m *DbManagerMock) CreateFirstBatch(ctx context.Context, l2coinbase common.Address) state.ProcessingContext {
	ret := _m.Called(ctx, l2coinbase)

	var r0 state.ProcessingContext
	if rf, ok := ret.Get(0).(func(context.Context, common.Address) state.ProcessingContext); ok {
		r0 = rf(ctx, l2coinbase)
	} else {
		r0 = ret.Get(0).(state.ProcessingContext)
	}

	return r0
}

// DeleteBatchByNumber provides a mock function with given fields: ctx, batchNumber, dbTx
func (_m *DbManagerMock) DeleteBatchByNumber(ctx context.Context, batchNumber uint64, dbTx pgx.Tx) error {
	ret := _m.Called(ctx, batchNumber, dbTx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) error); ok {
		r0 = rf(ctx, batchNumber, dbTx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteTransactionFromPool provides a mock function with given fields: ctx, txHash
func (_m *DbManagerMock) DeleteTransactionFromPool(ctx context.Context, txHash common.Hash) error {
	ret := _m.Called(ctx, txHash)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) error); ok {
		r0 = rf(ctx, txHash)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FlushMerkleTree provides a mock function with given fields: ctx
func (_m *DbManagerMock) FlushMerkleTree(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetBalanceByStateRoot provides a mock function with given fields: ctx, l2coinbase, root
func (_m *DbManagerMock) GetBalanceByStateRoot(ctx context.Context, address common.Address, root common.Hash) (*big.Int, error) {
	ret := _m.Called(ctx, address, root)

	var r0 *big.Int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, common.Hash) (*big.Int, error)); ok {
		return rf(ctx, address, root)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, common.Hash) *big.Int); ok {
		r0 = rf(ctx, address, root)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Address, common.Hash) error); ok {
		r1 = rf(ctx, address, root)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBatchByNumber provides a mock function with given fields: ctx, batchNumber, dbTx
func (_m *DbManagerMock) GetBatchByNumber(ctx context.Context, batchNumber uint64, dbTx pgx.Tx) (*state.Batch, error) {
	ret := _m.Called(ctx, batchNumber, dbTx)

	var r0 *state.Batch
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) (*state.Batch, error)); ok {
		return rf(ctx, batchNumber, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) *state.Batch); ok {
		r0 = rf(ctx, batchNumber, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*state.Batch)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, pgx.Tx) error); ok {
		r1 = rf(ctx, batchNumber, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDefaultMinGasPriceAllowed provides a mock function with given fields:
func (_m *DbManagerMock) GetDefaultMinGasPriceAllowed() uint64 {
	ret := _m.Called()

	var r0 uint64
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	return r0
}

// GetForcedBatch provides a mock function with given fields: ctx, forcedBatchNumber, dbTx
func (_m *DbManagerMock) GetForcedBatch(ctx context.Context, forcedBatchNumber uint64, dbTx pgx.Tx) (*state.ForcedBatch, error) {
	ret := _m.Called(ctx, forcedBatchNumber, dbTx)

	var r0 *state.ForcedBatch
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) (*state.ForcedBatch, error)); ok {
		return rf(ctx, forcedBatchNumber, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) *state.ForcedBatch); ok {
		r0 = rf(ctx, forcedBatchNumber, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*state.ForcedBatch)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, pgx.Tx) error); ok {
		r1 = rf(ctx, forcedBatchNumber, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetForcedBatchesSince provides a mock function with given fields: ctx, forcedBatchNumber, maxBlockNumber, dbTx
func (_m *DbManagerMock) GetForcedBatchesSince(ctx context.Context, forcedBatchNumber uint64, maxBlockNumber uint64, dbTx pgx.Tx) ([]*state.ForcedBatch, error) {
	ret := _m.Called(ctx, forcedBatchNumber, maxBlockNumber, dbTx)

	var r0 []*state.ForcedBatch
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, uint64, pgx.Tx) ([]*state.ForcedBatch, error)); ok {
		return rf(ctx, forcedBatchNumber, maxBlockNumber, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, uint64, pgx.Tx) []*state.ForcedBatch); ok {
		r0 = rf(ctx, forcedBatchNumber, maxBlockNumber, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*state.ForcedBatch)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, uint64, pgx.Tx) error); ok {
		r1 = rf(ctx, forcedBatchNumber, maxBlockNumber, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetForkIDByBatchNumber provides a mock function with given fields: batchNumber
func (_m *DbManagerMock) GetForkIDByBatchNumber(batchNumber uint64) uint64 {
	ret := _m.Called(batchNumber)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(uint64) uint64); ok {
		r0 = rf(batchNumber)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	return r0
}

// GetGasPrices provides a mock function with given fields: ctx
func (_m *DbManagerMock) GetGasPrices(ctx context.Context) (pool.GasPrices, error) {
	ret := _m.Called(ctx)

	var r0 pool.GasPrices
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (pool.GasPrices, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) pool.GasPrices); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(pool.GasPrices)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetL1AndL2GasPrice provides a mock function with given fields:
func (_m *DbManagerMock) GetL1AndL2GasPrice() (uint64, uint64) {
	ret := _m.Called()

	var r0 uint64
	var r1 uint64
	if rf, ok := ret.Get(0).(func() (uint64, uint64)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func() uint64); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(uint64)
	}

	return r0, r1
}

// GetLastBatch provides a mock function with given fields: ctx
func (_m *DbManagerMock) GetLastBatch(ctx context.Context) (*state.Batch, error) {
	ret := _m.Called(ctx)

	var r0 *state.Batch
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*state.Batch, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *state.Batch); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*state.Batch)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLastBatchNumber provides a mock function with given fields: ctx
func (_m *DbManagerMock) GetLastBatchNumber(ctx context.Context) (uint64, error) {
	ret := _m.Called(ctx)

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (uint64, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) uint64); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLastBlock provides a mock function with given fields: ctx, dbTx
func (_m *DbManagerMock) GetLastBlock(ctx context.Context, dbTx pgx.Tx) (*state.Block, error) {
	ret := _m.Called(ctx, dbTx)

	var r0 *state.Block
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) (*state.Block, error)); ok {
		return rf(ctx, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) *state.Block); ok {
		r0 = rf(ctx, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*state.Block)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, pgx.Tx) error); ok {
		r1 = rf(ctx, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLastClosedBatch provides a mock function with given fields: ctx
func (_m *DbManagerMock) GetLastClosedBatch(ctx context.Context) (*state.Batch, error) {
	ret := _m.Called(ctx)

	var r0 *state.Batch
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*state.Batch, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *state.Batch); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*state.Batch)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLastL2BlockHeader provides a mock function with given fields: ctx, dbTx
func (_m *DbManagerMock) GetLastL2BlockHeader(ctx context.Context, dbTx pgx.Tx) (*types.Header, error) {
	ret := _m.Called(ctx, dbTx)

	var r0 *types.Header
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) (*types.Header, error)); ok {
		return rf(ctx, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) *types.Header); ok {
		r0 = rf(ctx, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Header)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, pgx.Tx) error); ok {
		r1 = rf(ctx, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLastNBatches provides a mock function with given fields: ctx, numBatches
func (_m *DbManagerMock) GetLastNBatches(ctx context.Context, numBatches uint) ([]*state.Batch, error) {
	ret := _m.Called(ctx, numBatches)

	var r0 []*state.Batch
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint) ([]*state.Batch, error)); ok {
		return rf(ctx, numBatches)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint) []*state.Batch); ok {
		r0 = rf(ctx, numBatches)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*state.Batch)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(ctx, numBatches)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLastTrustedForcedBatchNumber provides a mock function with given fields: ctx, dbTx
func (_m *DbManagerMock) GetLastTrustedForcedBatchNumber(ctx context.Context, dbTx pgx.Tx) (uint64, error) {
	ret := _m.Called(ctx, dbTx)

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) (uint64, error)); ok {
		return rf(ctx, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) uint64); ok {
		r0 = rf(ctx, dbTx)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, pgx.Tx) error); ok {
		r1 = rf(ctx, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLatestGer provides a mock function with given fields: ctx, maxBlockNumber
func (_m *DbManagerMock) GetLatestGer(ctx context.Context, maxBlockNumber uint64) (state.GlobalExitRoot, time.Time, error) {
	ret := _m.Called(ctx, maxBlockNumber)

	var r0 state.GlobalExitRoot
	var r1 time.Time
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) (state.GlobalExitRoot, time.Time, error)); ok {
		return rf(ctx, maxBlockNumber)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64) state.GlobalExitRoot); ok {
		r0 = rf(ctx, maxBlockNumber)
	} else {
		r0 = ret.Get(0).(state.GlobalExitRoot)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64) time.Time); ok {
		r1 = rf(ctx, maxBlockNumber)
	} else {
		r1 = ret.Get(1).(time.Time)
	}

	if rf, ok := ret.Get(2).(func(context.Context, uint64) error); ok {
		r2 = rf(ctx, maxBlockNumber)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetLatestVirtualBatchTimestamp provides a mock function with given fields: ctx, dbTx
func (_m *DbManagerMock) GetLatestVirtualBatchTimestamp(ctx context.Context, dbTx pgx.Tx) (time.Time, error) {
	ret := _m.Called(ctx, dbTx)

	var r0 time.Time
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) (time.Time, error)); ok {
		return rf(ctx, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) time.Time); ok {
		r0 = rf(ctx, dbTx)
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	if rf, ok := ret.Get(1).(func(context.Context, pgx.Tx) error); ok {
		r1 = rf(ctx, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStorageAt provides a mock function with given fields: ctx, l2coinbase, position, root
func (_m *DbManagerMock) GetStorageAt(ctx context.Context, address common.Address, position *big.Int, root common.Hash) (*big.Int, error) {
	ret := _m.Called(ctx, address, position, root)

	var r0 *big.Int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, *big.Int, common.Hash) (*big.Int, error)); ok {
		return rf(ctx, address, position, root)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, *big.Int, common.Hash) *big.Int); ok {
		r0 = rf(ctx, address, position, root)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Address, *big.Int, common.Hash) error); ok {
		r1 = rf(ctx, address, position, root)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStoredFlushID provides a mock function with given fields: ctx
func (_m *DbManagerMock) GetStoredFlushID(ctx context.Context) (uint64, string, error) {
	ret := _m.Called(ctx)

	var r0 uint64
	var r1 string
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context) (uint64, string, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) uint64); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context) string); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(context.Context) error); ok {
		r2 = rf(ctx)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetTransactionsByBatchNumber provides a mock function with given fields: ctx, batchNumber
func (_m *DbManagerMock) GetTransactionsByBatchNumber(ctx context.Context, batchNumber uint64) ([]types.Transaction, []uint8, error) {
	ret := _m.Called(ctx, batchNumber)

	var r0 []types.Transaction
	var r1 []uint8
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) ([]types.Transaction, []uint8, error)); ok {
		return rf(ctx, batchNumber)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64) []types.Transaction); ok {
		r0 = rf(ctx, batchNumber)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64) []uint8); ok {
		r1 = rf(ctx, batchNumber)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]uint8)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, uint64) error); ok {
		r2 = rf(ctx, batchNumber)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetWIPBatch provides a mock function with given fields: ctx
func (_m *DbManagerMock) GetWIPBatch(ctx context.Context) (*WipBatch, error) {
	ret := _m.Called(ctx)

	var r0 *WipBatch
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*WipBatch, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *WipBatch); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*WipBatch)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsBatchClosed provides a mock function with given fields: ctx, batchNum
func (_m *DbManagerMock) IsBatchClosed(ctx context.Context, batchNum uint64) (bool, error) {
	ret := _m.Called(ctx, batchNum)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) (bool, error)); ok {
		return rf(ctx, batchNum)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64) bool); ok {
		r0 = rf(ctx, batchNum)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, batchNum)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OpenBatch provides a mock function with given fields: ctx, processingContext, dbTx
func (_m *DbManagerMock) OpenBatch(ctx context.Context, processingContext state.ProcessingContext, dbTx pgx.Tx) error {
	ret := _m.Called(ctx, processingContext, dbTx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, state.ProcessingContext, pgx.Tx) error); ok {
		r0 = rf(ctx, processingContext, dbTx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ProcessForcedBatch provides a mock function with given fields: ForcedBatchNumber, request
func (_m *DbManagerMock) ProcessForcedBatch(ForcedBatchNumber uint64, request state.ProcessRequest) (*state.ProcessBatchResponse, error) {
	ret := _m.Called(ForcedBatchNumber, request)

	var r0 *state.ProcessBatchResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(uint64, state.ProcessRequest) (*state.ProcessBatchResponse, error)); ok {
		return rf(ForcedBatchNumber, request)
	}
	if rf, ok := ret.Get(0).(func(uint64, state.ProcessRequest) *state.ProcessBatchResponse); ok {
		r0 = rf(ForcedBatchNumber, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*state.ProcessBatchResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(uint64, state.ProcessRequest) error); ok {
		r1 = rf(ForcedBatchNumber, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StoreProcessedTxAndDeleteFromPool provides a mock function with given fields: ctx, tx
func (_m *DbManagerMock) StoreProcessedTxAndDeleteFromPool(ctx context.Context, tx transactionToStore) error {
	ret := _m.Called(ctx, tx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, transactionToStore) error); ok {
		r0 = rf(ctx, tx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateTxStatus provides a mock function with given fields: ctx, hash, newStatus, isWIP, reason
func (_m *DbManagerMock) UpdateTxStatus(ctx context.Context, hash common.Hash, newStatus pool.TxStatus, isWIP bool, reason *string) error {
	ret := _m.Called(ctx, hash, newStatus, isWIP, reason)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash, pool.TxStatus, bool, *string) error); ok {
		r0 = rf(ctx, hash, newStatus, isWIP, reason)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewDbManagerMock creates a new instance of DbManagerMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDbManagerMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *DbManagerMock {
	mock := &DbManagerMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
