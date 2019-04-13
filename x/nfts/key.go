package nfts

const (
	// ModuleName is the name of the module
	ModuleName = "nfts"

	// StoreKey is the default store key for NFT bank
	StoreKey = ModuleName

	// QuerierRoute is the querier route for the NFT bank store.
	QuerierRoute = StoreKey
)

var (
	collectionKeyPrefix = []byte{0x00}
)

// GetCollectionKey gets the key of a collection
func GetCollectionKey(denom string) []byte {
	return append(collectionKeyPrefix, []byte(denom)...)
}
