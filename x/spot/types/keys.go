package types

const (
	// ModuleName defines the module name
	ModuleName = "spot"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_spot"

    
)

var (
	ParamsKey = []byte("p_spot")
)



func KeyPrefix(p string) []byte {
    return []byte(p)
}
