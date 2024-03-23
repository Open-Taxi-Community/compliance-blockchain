package types

const (
	// ModuleName defines the module name
	ModuleName = "compliance"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	ApprovalKey      = "Post/value/"
	ApprovalCountKey = "Post/count/"

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_compliance"
)

var (
	ParamsKey = []byte("p_compliance")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
