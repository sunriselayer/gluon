package keeper

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"gluon/x/contract/types"
)

// Use this in IBCMiddleware OnRecvPacket
// Module address must receive funds
func (k Keeper) SettleLazyContract(
	ctx sdk.Context,
	lazyContractId uint64,
	user sdk.AccAddress,
	coin sdk.Coin,
) error {
	lazyContract, found := k.GetLazyContract(ctx.Context(), lazyContractId)
	if !found {
		return errorsmod.Wrapf(sdkerrors.ErrNotFound, "LazyContract ID: %d", lazyContractId)
	}

	if lazyContract.Buyer == user.String() {
		if coin.Denom != lazyContract.DenomQuote {
			return types.ErrDenomMismatch
		}

		lazyContract.AmountPendingBuyer = lazyContract.AmountPendingBuyer.Sub(coin.Amount)
		lazyContract.AmountEscrowBuyer = lazyContract.AmountEscrowBuyer.Add(coin.Amount)
	} else if lazyContract.Seller == user.String() {
		if coin.Denom != lazyContract.DenomBase {
			return types.ErrDenomMismatch
		}

		lazyContract.AmountPendingSeller = lazyContract.AmountPendingSeller.Sub(coin.Amount)
		lazyContract.AmountEscrowSeller = lazyContract.AmountEscrowSeller.Add(coin.Amount)
	} else {
		return types.ErrNotBuyerNeitherSeller
	}
	// A case of isBuyer && isSeller can be skipped because it is guaranteed in Order.CrossValidate

	// Finalize the contract if there is no pending amount
	// Allow negative pending amount due to over payment
	if !lazyContract.AmountPendingBuyer.IsPositive() && !lazyContract.AmountPendingSeller.IsPositive() {
		buyerAddress, err := sdk.AccAddressFromBech32(lazyContract.Buyer)
		if err != nil {
			return err
		}
		sellerAddress, err := sdk.AccAddressFromBech32(lazyContract.Seller)
		if err != nil {
			return err
		}

		err = k.bankKeeper.SendCoinsFromModuleToAccount(
			ctx,
			types.ModuleName,
			buyerAddress,
			sdk.NewCoins(
				sdk.NewCoin(lazyContract.DenomBase, lazyContract.AmountEscrowSeller),
			),
		)
		if err != nil {
			return err
		}

		err = k.bankKeeper.SendCoinsFromModuleToAccount(
			ctx,
			types.ModuleName,
			sellerAddress,
			sdk.NewCoins(
				sdk.NewCoin(lazyContract.DenomQuote, lazyContract.AmountEscrowBuyer),
			),
		)
		if err != nil {
			return err
		}

		k.RemoveLazyContract(ctx.Context(), lazyContractId)
	}

	return nil
}
