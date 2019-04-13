# Cosmos NFT Module

## Interface

### The following methods from the 721 standard would need to become Msg Types:

* `transferFrom(address _from, address _to, uint256 _tokenId, bytes data)`
  * transfer on behalf of another user maybe out of scope for hackathon
  * This is a transfer function that includes the ability to transfer an asset on behalf of another user. We'd need to keep track of allowances so that a user would be able to give permission to another party to move assets on their behalf. This might be necessary for pegs as well–maybe the permission includes chain-ids as well?
* `approve(address _approved, uint256 _tokenId)`
  * __Out of scope__
  * This would be a necessary component of `transferFrom` to approve transfers on behalf of other parties.
* `setApprovalForAll(address _operator, bool _approved)`
  * __Out of scope__
  * Variation of `approve` to allow multiple approvals at once
* `safeTransferFrom(address _from, address _to, uint256 _tokenId)`
  * __Out of scope__
  * This version of `TransferFrom` checks with the party receiving the NFT to confirm they are prepared to accept the asset or whether they have an action that should be performed upon receiving it. Maybe out of scope for first Cosmo NFT spec but also good to consider this. Maybe we'd have a registry of recipients who can decide whether they accept an NFT sent to them or reject it?
* `safeTransferFrom(address _from, address _to, uint256 _tokenId, bytes data)`
  * __Out of scope__
  * This is a variation of `safeTransferFrom` without a data parameter. This data parameter is used to daisy-chain arbitrary on chain transactions. Something similar might be another Msg type that might want to be executed directly following the first transaction. QUESTION: Is it possible to create arbitrary transaction that bundle multiple MsgTypes? Maybe for 1st draft this implementation should not be reproduced on Cosmos.* `safeTransferFrom(address _from, address _to, uint256 _tokenId, bytes data)`

### Non-Spec but needed for Hackathon
* `mint()`
  * Need some custom logic for minting new tokens
* `burn()`
  * Need some custom logic for minting new tokens
* `setMetadata()`
  * Add / Update metadata for NFTs
  * Can we have variable number of key/value pairs here or do we need to just store it as a blob?

### The following queriers would need to be created from the 721 spec:
* `balanceOf(address _owner)`
  * Number of NFTs the address owns
* `ownerOf(uint256 _tokenId)`
  * Owner address of a specific NFT
* `getApproved(uint256 _tokenId)`
  * __Out of scope__
  * If transferFrom includes the ability to move on behalf of others it would be necessary to check the approvals
* `isApprovedForAll(address _owner, address _operator)`
  * __Out of scope__
  * An extension to the ability to check for approvals

The Metadata spec is actually separated from the required fields of the NFT. This might also be good to make optional with a reference to whether or not they exist. Something like [EIP 165](https://github.com/ethereum/EIPs/blob/master/EIPS/eip-165.md) would be needed for that and should also be considered. In the mean time we could expect the following queriers for metadata:
* `name()`
  * Return a string for the NFT family name
* `symbol()`
  * Return a string for the NFT family symbol
* ~~`tokenURI`~~
  * This should be replaced with an actual metadata data object. Even if all data is stored off chain the tokenURI for that data could be contained within the returned data object.
* `metadata`
  * This would return an object with a minimum of the ERC-721 metadata scheme, plus optional additional info under `tokenURI`. The return format may be a parameter so that if the recipient is prepared to parse JSON they can do so, but may also be able to request it with a different serialization.

```json
{
    "title": "Asset Metadata",
    "type": "object",
    "properties": {
        "name": {
            "type": "string",
            "description": "Identifies the asset to which this NFT represents"
        },
        "description": {
            "type": "string",
            "description": "Describes the asset to which this NFT represents"
        },
        "image": {
            "type": "string",
            "description": "A URI pointing to a resource with mime type image/* representing the asset to which this NFT represents. Consider making any images at a width between 320 and 1080 pixels and aspect ratio between 1.91:1 and 4:5 inclusive."
        },
        "tokenURI": {
            "type": "string",
            "description": "A URI pointing to a resource with more token related metadata that doesn't belong on chain."
        }
    }
}
```

### Types

```go
type NFT interface {
    map[type]type
}
```

```go
type NftMetadata struct {
    Description string `json:"description"`
    Image string `json:"image"`
    TokenURI string  `json:"url"`
}
```

```go
type NFT struct {
    Name string `json:"name"`
    Denom string `json:"denom"`
    Metadata NftMetadata `json:"metadata"`
    Owners map[sdk.AccAddress]uint64 `json:"owners"`
}
```

```go
type NFTs map[denom string][]NFT
```




### Messages




 ### Keeepers

 ```go
 (keeper Keeper) GetNFTsByDenom(ctx sdk.Context, denom string, uint64 ID) (nft NFT, err error) {
    store := ctx.KVStore(keeper.storeKey)
    b := store.Get(GetNFTKey(denom, ID))
    if b == nil {
        err = fmt.Errorf("NFT with ID %s doesn't exist", ID)
        return
    }
    keeper.cdc.MustUnmarshalBinaryLengthPrefixed(b, nft)
 }
```

 ```go
 (keeper Keeper) SetNFT(ctx sdk.Context, nft NFT) {
    store := ctx.KVStore(keeper.storeKey)
    nftKey := GetNFTKey(nft.ID)
    b := keeper.cdc.MustMarshalBinaryLengthPrefixed(nft)
    store.Set(nftKey, b)
 }
```

### Queriers

 * getName()
 * getSymbol()
 * getMetadata()