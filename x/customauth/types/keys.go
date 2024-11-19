package types

import "fmt"

const (
	// ModuleName defines the module name
	ModuleName = "customauth"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_customauth"
)

var (
	ParamsKey = []byte("p_customauth")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	PairingKey      = "Pairing/value/"
	PairingCountKey = "Pairing/count/"
)

func PairingKeyPrefix(address string) []byte {
	return KeyPrefix(fmt.Sprintf("%s%s/", PairingKey, address))
}
