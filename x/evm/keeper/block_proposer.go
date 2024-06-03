package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
)

// GetCoinbaseAddress returns the block proposer's validator operator address.
func (k Keeper) GetCoinbaseAddress(ctx sdk.Context, proposerAddress sdk.ConsAddress) (common.Address, error) {
	validator, err := k.stakingKeeper.GetValidatorByConsAddr(ctx, GetProposerAddress(ctx, proposerAddress))
	if err != nil {
		return common.Address{}, err
	}

	valOper, err := k.stakingKeeper.ValidatorAddressCodec().StringToBytes(validator.GetOperator())
	if err != nil {
		return common.Address{}, err
	}

	coinbase := common.BytesToAddress(valOper)
	return coinbase, nil
}

// GetProposerAddress returns current block proposer's address when provided proposer address is empty.
func GetProposerAddress(ctx sdk.Context, proposerAddress sdk.ConsAddress) sdk.ConsAddress {
	if len(proposerAddress) == 0 {
		proposerAddress = ctx.BlockHeader().ProposerAddress
	}
	return proposerAddress
}
