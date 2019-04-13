package nfts

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/bank"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Keeper maintains the link to data storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	coinKeeper bank.Keeper

	storeKey sdk.StoreKey // Unexposed key to access store from sdk.Context

	cdc *codec.Codec // The wire codec for binary encoding/decoding.
}

// NewKeeper creates new instances of the nft Keeper
func NewKeeper(coinKeeper bank.Keeper, storeKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{
		coinKeeper: coinKeeper,
		storeKey:   storeKey,
		cdc:        cdc,
	}
}

// GetNFT gets the entire NFT metadata struct for an ID
func (k Keeper) GetNFT(ctx sdk.Context, denom string, ID uint64,
) (nft NFT, err error) {
	collection, err := k.GetCollection(ctx, denom)
	if err != nil {
		return
	}
	nft, err = collection.GetNFT(ID)
	if err != nil {
		return
	}
	return
}

// SetNFT sets the entire NFT metadata struct for an ID
func (k Keeper) SetNFT(ctx sdk.Context, denom string, nft NFT) (err error) {
	collection, err := k.GetCollection(ctx, denom)
	if err != nil {
		return err
	}
	collection.NFTs[nft.ID] = nft
	k.SetCollection(ctx, denom, collection)
	return
}

// GetCollections returns all the NFTs collections
func (k Keeper) GetCollections(ctx sdk.Context) (collections []Collection) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, collectionKeyPrefix)
	defer iterator.Close()

	var collection Collection
	for ; iterator.Valid(); iterator.Next() {
		err := k.cdc.UnmarshalBinaryLengthPrefixed(iterator.Value(), &collection)
		if err != nil {
			panic(err)
		}
		collections = append(collections, collection)
	}
	return
}

// GetCollection returns a collection of NFTs
func (k Keeper) GetCollection(ctx sdk.Context, denom string,
) (collection Collection, err error) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(GetCollectionKey(denom))
	if b == nil {
		err = fmt.Errorf("collection of %s doesn't exist", denom)
		return
	}
	k.cdc.MustUnmarshalBinaryLengthPrefixed(b, collection)
	return
}

// SetCollection sets the entire collection of a single denom
func (k Keeper) SetCollection(ctx sdk.Context, denom string, collection Collection) {
	store := ctx.KVStore(k.storeKey)
	collectionKey := GetCollectionKey(denom)
	store.Set(collectionKey, k.cdc.MustMarshalBinaryBare(collection))
}
