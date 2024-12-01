package types

const (
	// ModuleName defines the module name
	ModuleName = "order"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_order"

    
)

var (
	ParamsKey = []byte("p_order")
)



func KeyPrefix(p string) []byte {
    return []byte(p)
}
