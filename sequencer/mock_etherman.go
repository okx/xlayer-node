// Code generated by mockery v2.39.0. DO NOT EDIT.

package sequencer

import (
	context "context"
	big "math/big"

	common "github.com/ethereum/go-ethereum/common"

	coretypes "github.com/ethereum/go-ethereum/core/types"

	mock "github.com/stretchr/testify/mock"

	types "github.com/0xPolygonHermez/zkevm-node/etherman/types"
)

// EthermanMock is an autogenerated mock type for the etherman type
type EthermanMock struct {
	mock.Mock
}

// BuildSequenceBatchesTxData provides a mock function with given fields: sender, sequences, l2CoinBase
func (_m *EthermanMock) BuildSequenceBatchesTxData(sender common.Address, sequences []types.Sequence, l2CoinBase common.Address, committeeSignaturesAndAddrs []byte) (*common.Address, []byte, error) {
	ret := _m.Called(sender, sequences, l2CoinBase)

	if len(ret) == 0 {
		panic("no return value specified for BuildSequenceBatchesTxData")
	}

	var r0 *common.Address
	var r1 []byte
	var r2 error
	if rf, ok := ret.Get(0).(func(common.Address, []types.Sequence, common.Address) (*common.Address, []byte, error)); ok {
		return rf(sender, sequences, l2CoinBase)
	}
	if rf, ok := ret.Get(0).(func(common.Address, []types.Sequence, common.Address) *common.Address); ok {
		r0 = rf(sender, sequences, l2CoinBase)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*common.Address)
		}
	}

	if rf, ok := ret.Get(1).(func(common.Address, []types.Sequence, common.Address) []byte); ok {
		r1 = rf(sender, sequences, l2CoinBase)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]byte)
		}
	}

	if rf, ok := ret.Get(2).(func(common.Address, []types.Sequence, common.Address) error); ok {
		r2 = rf(sender, sequences, l2CoinBase)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// EstimateGasSequenceBatches provides a mock function with given fields: sender, sequences, l2CoinBase
func (_m *EthermanMock) EstimateGasSequenceBatches(sender common.Address, sequences []types.Sequence, l2CoinBase common.Address, committeeSignaturesAndAddrs []byte) (*coretypes.Transaction, error) {
	ret := _m.Called(sender, sequences, l2CoinBase)

	if len(ret) == 0 {
		panic("no return value specified for EstimateGasSequenceBatches")
	}

	var r0 *coretypes.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(common.Address, []types.Sequence, common.Address) (*coretypes.Transaction, error)); ok {
		return rf(sender, sequences, l2CoinBase)
	}
	if rf, ok := ret.Get(0).(func(common.Address, []types.Sequence, common.Address) *coretypes.Transaction); ok {
		r0 = rf(sender, sequences, l2CoinBase)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*coretypes.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(common.Address, []types.Sequence, common.Address) error); ok {
		r1 = rf(sender, sequences, l2CoinBase)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLatestBatchNumber provides a mock function with given fields:
func (_m *EthermanMock) GetLatestBatchNumber() (uint64, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetLatestBatchNumber")
	}

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func() (uint64, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLatestBlockNumber provides a mock function with given fields: ctx
func (_m *EthermanMock) GetLatestBlockNumber(ctx context.Context) (uint64, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetLatestBlockNumber")
	}

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

// GetLatestBlockTimestamp provides a mock function with given fields: ctx
func (_m *EthermanMock) GetLatestBlockTimestamp(ctx context.Context) (uint64, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetLatestBlockTimestamp")
	}

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

// GetSendSequenceFee provides a mock function with given fields: numBatches
func (_m *EthermanMock) GetSendSequenceFee(numBatches uint64) (*big.Int, error) {
	ret := _m.Called(numBatches)

	if len(ret) == 0 {
		panic("no return value specified for GetSendSequenceFee")
	}

	var r0 *big.Int
	var r1 error
	if rf, ok := ret.Get(0).(func(uint64) (*big.Int, error)); ok {
		return rf(numBatches)
	}
	if rf, ok := ret.Get(0).(func(uint64) *big.Int); ok {
		r0 = rf(numBatches)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(numBatches)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TrustedSequencer provides a mock function with given fields:
func (_m *EthermanMock) TrustedSequencer() (common.Address, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for TrustedSequencer")
	}

	var r0 common.Address
	var r1 error
	if rf, ok := ret.Get(0).(func() (common.Address, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() common.Address); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Address)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewEthermanMock creates a new instance of EthermanMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEthermanMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *EthermanMock {
	mock := &EthermanMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
