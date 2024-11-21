package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ModuleName defines the module name
	ModuleName = "contract"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_contract"

	// Version defines the current version the IBC module supports
	Version = "contract-1"

	// PortID is the default port id that module binds to
	PortID = "contract"
)

var (
	ParamsKey = []byte("p_contract")
)

var (
	// PortKey defines the key to store the port ID in store
	PortKey = KeyPrefix("contract-port-")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	LazyContractKey         = "LazyContract/value/"
	LazyContractCountKey    = "LazyContract/count/"
	SortedLazyContractIdKey = "SortedLazyContract/id/"
)

func SortedLazyContractIdPrefix(expiry time.Time) []byte {
	return append([]byte(SortedLazyContractIdKey), sdk.Uint64ToBigEndian(uint64(expiry.UnixMilli()))...)
}
