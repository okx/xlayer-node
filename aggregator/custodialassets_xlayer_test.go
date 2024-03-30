package aggregator

import (
	"context"
	"math/big"
	"testing"
	"time"

	agglayerTypes "github.com/0xPolygon/agglayer/rpc/types"
	"github.com/0xPolygon/agglayer/tx"
	zktypes "github.com/0xPolygonHermez/zkevm-node/config/types"
	"github.com/0xPolygonHermez/zkevm-node/hex"
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
)

const (
	domain = "http://asset-onchain.base-defi.svc.test.local:7001"
	//	seqAddr                = "1a13bddcc02d363366e04d4aa588d3c125b0ff6f"
	//	aggAddr                = "66e39a1e507af777e8c385e2d91559e20e306303"
	seqAddr                = "66e39a1e507af777e8c385e2d91559e20e306303"
	aggAddr                = "1a13bddcc02d363366e04d4aa588d3c125b0ff6f"
	contractAddr           = "837bf712c91949da16e0201045ecabc669eaf4ba"
	contractAddrAgg        = "837bf712c91949da16e0201045ecabc669eaf4ba"
	l1ChainID       uint64 = 11155111
	AccessKey              = "74w82q40cz"
	SecretKey              = "3v2c07o12j9760ag"
	// domain       = "http://127.0.0.1:7001"
	// seqAddr      = "f39Fd6e51aad88F6F4ce6aB8827279cffFb92266"
	// aggAddr      = "70997970C51812dc3A010C7d01b50e0d17dc79C8"
	// contractAddr = "812cB73e48841a6736bB94c65c56341817cE6304"
)

func TestClientPostSignRequestAndWaitResultTxSigner(t *testing.T) {
	newStateRoot := common.BigToHash(big.NewInt(100))
	t.Log(newStateRoot.Hex())
	newLocalExitRoot := common.BigToHash(big.NewInt(1000))
	t.Log(newLocalExitRoot.Hex())
	proof := agglayerTypes.ArgBytes([]byte("sampleProof"))

	t.Log(proof.Hex())

	tnx := tx.Tx{
		LastVerifiedBatch: agglayerTypes.ArgUint64(36),
		NewVerifiedBatch:  *agglayerTypes.ArgUint64Ptr(72),
		ZKP: tx.ZKP{
			NewStateRoot:     newStateRoot,
			NewLocalExitRoot: newLocalExitRoot,
			Proof:            proof,
		},
		RollupID: 2,
	}
	t.Log(tnx.Hash())

	agg := &Aggregator{
		cfg: Config{
			CustodialAssets: CustodialAssetsConfig{
				Enable:            true,
				URL:               domain,
				Symbol:            2882,
				SequencerAddr:     common.HexToAddress(seqAddr),
				AggregatorAddr:    common.HexToAddress(aggAddr),
				WaitResultTimeout: zktypes.NewDuration(4 * time.Minute),
				OperateTypeSeq:    5,
				OperateTypeAgg:    6,
				ProjectSymbol:     3011,
				OperateSymbol:     2,
				SysFrom:           3,
				UserID:            0,
				OperateAmount:     0,
				RequestSignURI:    "/priapi/v1/assetonchain/ecology/ecologyOperate",
				QuerySignURI:      "/priapi/v1/assetonchain/ecology/querySignDataByOrderNo",
				AccessKey:         AccessKey,
				SecretKey:         SecretKey,
			},
		},
	}
	ctx := context.WithValue(context.Background(), traceID, uuid.New().String())

	myTx, err := agg.signTx(ctx, tnx)
	if err == nil {
		t.Log(myTx.Tx)
		t.Log(myTx.Signature)
	} else {
		t.Log(err)
	}
}

func TestTxHash(t *testing.T) {
	newStateRoot := common.BigToHash(big.NewInt(10000000))
	t.Log(newStateRoot.Hex())
	newLocalExitRoot := common.BigToHash(big.NewInt(500000000))
	t.Log(newLocalExitRoot.Hex())
	proof := agglayerTypes.ArgBytes([]byte("sampleProof"))

	t.Log(proof.Hex())

	t.Log(agglayerTypes.ArgUint64(20000000).Hex())
	t.Log(hex.EncodeToString([]byte(agglayerTypes.ArgUint64(20000000).Hex())))

	tnx := tx.Tx{
		LastVerifiedBatch: agglayerTypes.ArgUint64(20000000),
		NewVerifiedBatch:  agglayerTypes.ArgUint64(300000000),
		ZKP: tx.ZKP{
			NewStateRoot:     newStateRoot,
			NewLocalExitRoot: newLocalExitRoot,
			Proof:            proof,
		},
		RollupID: 2,
	}
	t.Log(tnx.Hash())
}

func TestApprove(t *testing.T) {
	agg := &Aggregator{
		cfg: Config{
			CustodialAssets: CustodialAssetsConfig{
				Enable:            true,
				URL:               domain,
				Symbol:            2882,
				SequencerAddr:     common.HexToAddress(seqAddr),
				AggregatorAddr:    common.HexToAddress(aggAddr),
				WaitResultTimeout: zktypes.NewDuration(4 * time.Minute),
				OperateTypeSeq:    5,
				OperateTypeAgg:    7,
				ProjectSymbol:     3011,
				OperateSymbol:     2,
				SysFrom:           3,
				UserID:            0,
				OperateAmount:     0,
				RequestSignURI:    "/priapi/v1/assetonchain/ecology/ecologyOperate",
				QuerySignURI:      "/priapi/v1/assetonchain/ecology/querySignDataByOrderNo",
				AccessKey:         AccessKey,
				SecretKey:         SecretKey,
			},
		},
	}
	approveToAddress := "0x43Db1155C06548666E2928f4970694CA21B1835a"
	approveAmount := "10000000000000000000000000"
	contractAddress := "0x6a7c3f4b0651d6da389ad1d11d962ea458cdca70"

	gasLimit := 80000
	nonce := 708
	gasPrice := "0.0000001"

	payload, err := agg.approve(approveToAddress, approveAmount, contractAddress, uint64(gasLimit), uint64(nonce), gasPrice)
	if err != nil {
		t.Log(hex.EncodeToString(payload))
	} else {
		t.Log(err)
	}
}
