package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/distribution/types"
)

// get outstanding rewards
func (k Keeper) GetValidatorOutstandingRewardsCoins(ctx context.Context, val sdk.ValAddress) (sdk.DecCoins, error) {
	rewards, err := k.GetValidatorOutstandingRewards(ctx, val)
	if err != nil {
		return nil, err
	}

	return rewards.Rewards, nil
}

// GetFeePoolLiquidityProviderCoins gets the coins in the fee pool that is for liquidity providers.
func (k Keeper) GetFeePoolLiquidityProviderCoins(ctx context.Context) (sdk.DecCoins, error) {
	feePool, err := k.FeePool.Get(ctx)
	if err != nil {
		return nil, err
	}
	return feePool.LiquidityProviderPool, nil
}

// GetDistributionAccount returns the distribution ModuleAccount
func (k Keeper) GetDistributionAccount(ctx context.Context) sdk.ModuleAccountI {
	return k.authKeeper.GetModuleAccount(ctx, types.ModuleName)
}
