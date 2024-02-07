package gasprice

import (
	"context"
	"time"

	"github.com/0xPolygonHermez/zkevm-node/etherman"
	"github.com/0xPolygonHermez/zkevm-node/log"
	"github.com/0xPolygonHermez/zkevm-node/state"
)

// L2GasPricer interface for gas price suggester.
type L2GasPricer interface {
	UpdateGasPriceAvg()
}

// Apollo fetch dynamic config from apollo.
type Apollo interface {
	FetchL2GasPricerConfig(config *Config)
}

// NewL2GasPriceSuggester init.
func NewL2GasPriceSuggester(ctx context.Context, cfg Config, pool poolInterface, ethMan *etherman.Client, state *state.State, fetch Apollo) {
	var gpricer L2GasPricer
	switch cfg.Type {
	case LastNBatchesType:
		log.Info("Lastnbatches type selected")
		gpricer = newLastNL2BlocksGasPriceSuggester(ctx, cfg, state, pool)
	case FollowerType:
		log.Info("Follower type selected")
		gpricer = newFollowerGasPriceSuggester(ctx, cfg, pool, ethMan, fetch)
	case DefaultType:
		log.Info("Default type selected")
		gpricer = newDefaultGasPriceSuggester(ctx, cfg, pool, fetch)
	case FixedType:
		log.Info("Fixed type selected")
		gpricer = newFixedGasPriceSuggester(ctx, cfg, state, pool, ethMan, fetch)
	default:
		log.Fatal("unknown l2 gas price suggester type ", cfg.Type, ". Please specify a valid one: 'lastnbatches', 'follower' or 'default'")
	}

	updateTimer := time.NewTimer(cfg.UpdatePeriod.Duration)
	cleanTimer := time.NewTimer(cfg.CleanHistoryPeriod.Duration)
	for {
		select {
		case <-ctx.Done():
			log.Info("Finishing l2 gas price suggester...")
			return
		case <-updateTimer.C:
			gpricer.UpdateGasPriceAvg()
			updateTimer.Reset(cfg.UpdatePeriod.Duration)
		case <-cleanTimer.C:
			cleanGasPriceHistory(pool, cfg.CleanHistoryTimeRetention.Duration)
			cleanTimer.Reset(cfg.CleanHistoryPeriod.Duration)
		}
	}
}

func cleanGasPriceHistory(pool poolInterface, timeRetention time.Duration) {
	ctx := context.Background()
	err := pool.DeleteGasPricesHistoryOlderThan(ctx, time.Now().UTC().Add(-timeRetention))
	if err != nil {
		log.Errorf("failed to delete pool gas price history: %v", err)
	}
}
