package ante

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type OperatorMsgMatchHandler func(msg sdk.Msg) bool

func Match(tx sdk.Tx, msgMatchHandler OperatorMsgMatchHandler) bool {
	for _, msg := range tx.GetMsgs() {
		if msgMatchHandler(msg) {
			continue
		}

		return false
	}

	return true
}
