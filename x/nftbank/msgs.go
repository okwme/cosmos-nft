package nftbank

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	// sdk "github.com/cosmos/cosmos-sdk/types"
)

// MsgTransferNFT defines a TransferNFT message
type MsgTransferNFT struct {
	Sender      sdk.AccAddress
	Recepient   sdk.AccAddress
	Collectibles Collectibles
}

// NewMsgSetNFT is a constructor function for MsgSetName
func MsgTransferNFT(sender, recepient sdk.AccAddress, collectible) MsgSetNFT {
	return MsgSetNFT{
		Denom:       denom,
		ID:          id,
		Owner:       owner,
		Name:        name,
		Description: description,
		Image:       image,
		TokenURI:    tokenURI,
	}
}

// MsgSenNFT defines a SetName message
type MsgTransferNFTFrom struct {
	Sender sdk.AccAddress
	Owner sdk.AccAddress
	Recipient sdk.AccAddress
	NftID uint64
}

// MsgTransferNFTFrom is a constructor function for MsgSetName
func MsgTransferNFTFrom(sender, owner, recipient sdk.AccAddress, id uint64) MsgSetNFT {
	return MsgSetNFT{
		Sender: sender,
		Owner: owner,
		Recipient: recipient,
		NftID: id,
	}
}

