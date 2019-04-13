package nftbank

const (
	// ModuleName is the name of the module
	ModuleName = "nftbank"

	// StoreKey is the default store key for NFT bank
	StoreKey = ModuleName

	// QuerierRoute is the querier route for the NFT bank store.
	QuerierRoute = StoreKey
)

var (
	nftKeyPrefix = []byte{0x00}
)

//  returns the store key of the given module
func GetNFTKey(ID uint64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, ID)
	return append(nftKeyPrefix, b)
}
