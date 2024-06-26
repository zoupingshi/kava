package incentive

import (
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/kava-labs/kava/x/incentive/keeper"
	"github.com/kava-labs/kava/x/incentive/types"
)

// BeginBlocker runs at the start of every block
func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)

	params := k.GetParams(ctx)

	for _, rp := range params.USDXMintingRewardPeriods {
		k.AccumulateUSDXMintingRewards(ctx, rp)
	}
	for _, rp := range params.HardSupplyRewardPeriods {
		k.AccumulateHardSupplyRewards(ctx, rp)
	}
	for _, rp := range params.HardBorrowRewardPeriods {
		k.AccumulateHardBorrowRewards(ctx, rp)
	}
	for _, rp := range params.DelegatorRewardPeriods {
		k.AccumulateDelegatorRewards(ctx, rp)
	}
	for _, rp := range params.SwapRewardPeriods {
		k.AccumulateSwapRewards(ctx, rp)
	}
	for _, rp := range params.SavingsRewardPeriods {
		k.AccumulateSavingsRewards(ctx, rp)
	}
	for _, rp := range params.EarnRewardPeriods {
		if err := k.AccumulateEarnRewards(ctx, rp); err != nil {
			panic(fmt.Sprintf("failed to accumulate earn rewards: %s", err))
		}
	}
}
