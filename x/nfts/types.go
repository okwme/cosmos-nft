package nfts

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NFT non fungible token definition
type NFT struct {
	ID          uint64         `json:"id"`
	Owner       sdk.AccAddress `json:"owner"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Image       string         `json:"image"`
	TokenURI    string         `json:"url"`
}

// NewNFT creates a new NFT
func NewNFT(ID uint64, owner sdk.AccAddress, tokenURI, description, image, name string,
) NFT {
	return NFT{
		ID:          ID,
		Owner:       owner,
		Name:        strings.TrimSpace(name),
		Description: strings.TrimSpace(description),
		Image:       strings.TrimSpace(image),
		TokenURI:    strings.TrimSpace(tokenURI),
	}
}

// Collection of non fungible tokens
type Collection struct {
	Denom string         `json:"denom"`
	NFTs  map[uint64]NFT `json:"nfts"`
}

// NewCollection creates a new NFT Collection
func NewCollection(denom string) Collection {
	return Collection{
		Denom: strings.TrimSpace(denom),
		NFTs:  make(map[uint64]NFT),
	}
}

// GetNFT gets a NFT
func (collection Collection) GetNFT(ID uint64) (nft NFT, err error) {
	nft, ok := collection.NFTs[ID]
	if !ok {
		return nft, fmt.Errorf("collection %s doesn't contain an NFT with ID %d", collection.Denom, ID)
	}
	return
}

// ValidateBasic validates a Collection
func (collection Collection) ValidateBasic() (err error) {
	if collection.Denom == "" {
		return fmt.Errorf("collection name cannot be blank")
	}
	if len(collection.NFTs) == 0 {
		return fmt.Errorf("collection %s cannot be empty")
	}
	return
}
