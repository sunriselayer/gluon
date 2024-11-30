package types

const (
	// ModuleName defines the module name
	ModuleName = "perp"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_perp"
)

var (
	ParamsKey = []byte("p_perp")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
