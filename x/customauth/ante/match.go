package ante

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
)

type OperatorMsgMatchHandler func(msg sdk.Msg) bool

func IsOperatorMsg(tx sdk.Tx, msgMatchHandler OperatorMsgMatchHandler) bool {
	for _, msg := range tx.GetMsgs() {
		if msgMatchHandler(msg) {
			continue
		}

		return false
	}

	return true
}

func GetPairingId(tx authsigning.Tx) *uint64 {

	return nil
}
