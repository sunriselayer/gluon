package keeper

import (
	"context"
	"gluon/x/contract/types"
	"time"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func transferAmount(params types.Params, amount sdkmath.Int, lazyContract bool) sdkmath.Int {
	if lazyContract {
		return params.LazyContractMarginRatio.MulInt(amount).RoundInt()
	}

	return amount
}

func (k Keeper) transferRecipientAddress(defaultTarget sdk.AccAddress, atLeastOneLazyContract bool) sdk.AccAddress {
	if atLeastOneLazyContract {
		return k.AccountKeeper.GetModuleAddress(types.ModuleName)
	}

	return defaultTarget
}

func (k Keeper) Contract(
	ctx sdk.Context,
	buy types.Order,
	sell types.Order,
	amount sdkmath.Int,
	price sdkmath.LegacyDec,
) error {
	denomBase := buy.DenomBase
	denomQuote := buy.DenomQuote

	addressBuy, err := sdk.AccAddressFromBech32(buy.Address)
	if err != nil {
		return err
	}
	addressSell, err := sdk.AccAddressFromBech32(sell.Address)
	if err != nil {
		return err
	}

	params := k.GetParams(ctx)
	amountBase := amount
	amountQuote := price.MulInt(amount).RoundInt()
	amountBaseTransfer := transferAmount(params, amountBase, sell.LazyContract)
	amountQuoteTransfer := transferAmount(params, amountQuote, buy.LazyContract)

	atLeastOneLazyContract := buy.LazyContract || sell.LazyContract

	if amountQuoteTransfer.IsPositive() {
		recipient := k.transferRecipientAddress(addressSell, atLeastOneLazyContract)
		err = k.bankKeeper.SendCoins(
			ctx,
			addressBuy,
			recipient,
			sdk.NewCoins(
				sdk.NewCoin(denomQuote, amountQuoteTransfer),
			),
		)
		if err != nil {
			return err
		}
	}

	if amountBaseTransfer.IsPositive() {
		recipient := k.transferRecipientAddress(addressBuy, atLeastOneLazyContract)
		err = k.bankKeeper.SendCoins(
			ctx,
			addressSell,
			recipient,
			sdk.NewCoins(
				sdk.NewCoin(denomBase, amountBaseTransfer),
			),
		)
		if err != nil {
			return err
		}
	}

	err = k.AddContractedAmount(ctx, buy.Expiry, buy.Id, amount)
	if err != nil {
		return err
	}
	err = k.AddContractedAmount(ctx, sell.Expiry, sell.Id, amount)
	if err != nil {
		return err
	}

	if atLeastOneLazyContract {
		expiry := ctx.BlockTime().Add(params.LazyContractPeriod)

		// Check order expiry and contract expiry to prevent from order info garbage collected before the settlement
		if buy.Expiry.Before(expiry) || sell.Expiry.Before(expiry) {
			return types.ErrOrderExpireBeforeLazyContract
		}

		lazyContractId := k.AppendLazyContract(ctx, types.LazyContract{
			Buyer:               buy.Address,
			Seller:              sell.Address,
			DenomBase:           denomBase,
			DenomQuote:          denomQuote,
			AmountEscrowBuyer:   amountQuoteTransfer,
			AmountEscrowSeller:  amountBaseTransfer,
			AmountPendingBuyer:  amountQuote.Sub(amountQuoteTransfer),
			AmountPendingSeller: amountBase.Sub(amountBaseTransfer),
			Expiry:              expiry,
		})
		k.SetSortedLazyContract(ctx, types.SortedLazyContract{
			Expiry: uint64(expiry.UnixMilli()),
			Id:     lazyContractId,
		})
	}

	return nil
}

func (k Keeper) AddContractedAmount(ctx context.Context, expiry time.Time, id string, amount sdkmath.Int) error {
	sortedOrder, found := k.GetSortedOrder(ctx, expiry, id)
	if !found {
		return types.ErrOrderNotFound
	}

	sortedOrder.ContractedAmount = sortedOrder.ContractedAmount.Add(amount)
	k.SetSortedOrder(ctx, sortedOrder)

	return nil
}
