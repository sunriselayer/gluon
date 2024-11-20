package ante

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type CustomSigVerifierMap map[sdk.Msg]func(cak CustomAuthKeeper, tx sdk.Tx) error

func (m CustomSigVerifierMap) Match(tx sdk.Tx) bool {
	for _, msg := range tx.GetMsgs() {
		if _, found := m[msg]; found {
			continue
		}

		return false
	}

	return true
}
