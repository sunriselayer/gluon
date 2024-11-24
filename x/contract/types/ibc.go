package types

import (
	"encoding/json"
	"fmt"
	"strings"

	transfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
	"github.com/gogo/protobuf/jsonpb"
)

func (m *ContractMetadata) Validate() error {

	return nil
}

type ContractAcknowledgement struct {
	IncomingAck []byte `json:"ibc_ack"`
}

func (a ContractAcknowledgement) Acknowledgement() ([]byte, error) {
	bz, err := json.Marshal(a)
	if err != nil {
		return nil, err
	}

	return bz, nil
}

func GetDenomForThisChain(port, channel, counterpartyPort, counterpartyChannel, denom string) string {
	counterpartyPrefix := transfertypes.GetDenomPrefix(counterpartyPort, counterpartyChannel)
	if strings.HasPrefix(denom, counterpartyPrefix) {
		// unwind denom
		unwoundDenom := denom[len(counterpartyPrefix):]
		denomTrace := transfertypes.ParseDenomTrace(unwoundDenom)
		if denomTrace.Path == "" {
			// denom is now unwound back to native denom
			return unwoundDenom
		}
		// denom is still IBC denom
		return denomTrace.IBCDenom()
	}
	// append port and channel from this chain to denom
	prefixedDenom := transfertypes.GetDenomPrefix(port, channel) + denom
	return transfertypes.ParseDenomTrace(prefixedDenom).IBCDenom()
}

func DecodeContractMetadata(memo string) (*PacketMetadata, error) {
	d := make(map[string]interface{})
	err := json.Unmarshal([]byte(memo), &d)
	if err != nil {
		return nil, err
	}
	if d["contract"] == nil {
		return nil, fmt.Errorf("no contract field in memo")
	}

	m := &PacketMetadata{}
	err = jsonpb.Unmarshal(strings.NewReader(memo), m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
