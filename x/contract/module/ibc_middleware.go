package contract

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	transfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
	channeltypes "github.com/cosmos/ibc-go/v8/modules/core/04-channel/types"
	porttypes "github.com/cosmos/ibc-go/v8/modules/core/05-port/types"
	exported "github.com/cosmos/ibc-go/v8/modules/core/exported"

	keeper "gluon/x/contract/keeper"
	types "gluon/x/contract/types"

	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type IBCMiddleware struct {
	porttypes.IBCModule
	keeper *keeper.Keeper
}

// NewIBCMiddleware creates a new IBCMiddleware given the keeper and underlying application.
func NewIBCMiddleware(
	app porttypes.IBCModule,
	k *keeper.Keeper,
) IBCMiddleware {
	return IBCMiddleware{
		IBCModule: app,
		keeper:    k,
	}
}

// receiveFunds receives funds from the packet into the override receiver
// address and returns an error if the funds cannot be received.
func (im IBCMiddleware) receiveFunds(
	ctx sdk.Context,
	packet channeltypes.Packet,
	data transfertypes.FungibleTokenPacketData,
	overrideReceiver string,
	relayer sdk.AccAddress,
) (exported.Acknowledgement, error) {
	overrideData := transfertypes.FungibleTokenPacketData{
		Denom:    data.Denom,
		Amount:   data.Amount,
		Sender:   data.Sender,
		Receiver: overrideReceiver, // override receiver
		// Memo explicitly zeroed
	}
	overrideDataBz := transfertypes.ModuleCdc.MustMarshalJSON(&overrideData)
	overridePacket := channeltypes.Packet{
		Sequence:           packet.Sequence,
		SourcePort:         packet.SourcePort,
		SourceChannel:      packet.SourceChannel,
		DestinationPort:    packet.DestinationPort,
		DestinationChannel: packet.DestinationChannel,
		Data:               overrideDataBz, // override data
		TimeoutHeight:      packet.TimeoutHeight,
		TimeoutTimestamp:   packet.TimeoutTimestamp,
	}

	ack := im.IBCModule.OnRecvPacket(ctx, overridePacket, relayer)

	if ack == nil {
		return ack, fmt.Errorf("ack is nil")
	}

	if !ack.Success() {
		return ack, fmt.Errorf("ack error: %s", string(ack.Acknowledgement()))
	}

	return ack, nil
}

// OnRecvPacket checks the memo field on this packet and if the metadata inside's root key indicates this packet
// should be handled by the swap middleware it attempts to perform a swap. If the swap is successful
// the underlying application's OnRecvPacket callback is invoked, an ack error is returned otherwise.
func (im IBCMiddleware) OnRecvPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	relayer sdk.AccAddress,
) exported.Acknowledgement {
	var data transfertypes.FungibleTokenPacketData
	if err := transfertypes.ModuleCdc.UnmarshalJSON(packet.GetData(), &data); err != nil {
		// If this happens either a) a user has crafted an invalid packet, b) a
		// software developer has connected the middleware to a stack that does
		// not have a transfer module, or c) the transfer module has been modified
		// to accept other Packets. The best thing we can do here is pass the packet
		// on down the stack.
		return im.IBCModule.OnRecvPacket(ctx, packet, relayer)
	}

	m, err := types.DecodeContractMetadata(data.Memo)
	if err != nil {
		return im.IBCModule.OnRecvPacket(ctx, packet, relayer)
	}
	metadata := *m.Contract

	if err := metadata.Validate(); err != nil {
		return channeltypes.NewErrorAcknowledgement(err)
	}

	// Receive the incoming fund
	settlerModuleAddress := im.keeper.AccountKeeper.GetModuleAddress(types.ModuleName)
	incomingAck, err := im.receiveFunds(ctx, packet, data, settlerModuleAddress.String(), relayer)
	if err != nil {
		return channeltypes.NewErrorAcknowledgement(err)
	}

	denom := types.GetDenomForThisChain(
		packet.DestinationPort,
		packet.DestinationChannel,
		packet.SourcePort,
		packet.SourceChannel,
		data.Denom,
	)

	amount, ok := sdkmath.NewIntFromString(data.Amount)
	if !ok {
		return channeltypes.NewErrorAcknowledgement(errorsmod.Wrap(sdkerrors.ErrInvalidCoins, "invalid coin amount"))
	}

	user, err := sdk.AccAddressFromBech32(data.Receiver)
	if err != nil {
		return channeltypes.NewErrorAcknowledgement(err)
	}
	err = im.keeper.SettleLazyContract(
		ctx,
		metadata.LazyContractId,
		user,
		sdk.NewCoin(denom, amount),
	)
	if err != nil {
		return channeltypes.NewErrorAcknowledgement(err)
	}

	fullAck := types.ContractAcknowledgement{
		IncomingAck: incomingAck.Acknowledgement(),
	}

	bz, err := fullAck.Acknowledgement()
	if err != nil {
		return channeltypes.NewErrorAcknowledgement(err)
	}

	return channeltypes.NewResultAcknowledgement(bz)
}

// OnAcknowledgementPacket implements the IBCModule interface.
func (im IBCMiddleware) OnAcknowledgementPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	acknowledgement []byte,
	relayer sdk.AccAddress,
) error {
	return im.IBCModule.OnAcknowledgementPacket(ctx, packet, acknowledgement, relayer)
}

// OnTimeoutPacket implements the IBCModule interface.
func (im IBCMiddleware) OnTimeoutPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	relayer sdk.AccAddress,
) error {
	return im.IBCModule.OnTimeoutPacket(ctx, packet, relayer)
}
