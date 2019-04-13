package nfts

import (
	"strings"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var cdc = codec.New()

// verify interface at compile time
// var _ sdk.Msg = &MsgTransferNFT{} // TODO: fix

// MsgTransferNFT defines a TransferNFT message
type MsgTransferNFT struct {
	Sender    sdk.AccAddress
	Recipient sdk.AccAddress
	Denom     string
	ID        uint64
}

// NewMsgTransferNFT is a constructor function for MsgSetName
func NewMsgTransferNFT(sender, recipient sdk.AccAddress, denom string, id uint64,
) MsgTransferNFT {
	return MsgTransferNFT{
		Sender:    sender,
		Recipient: recipient,
		Denom:     denom,
		ID:        id,
	}
}

// Route Implements Msg TODO: fix
// func (msg MsgTransferNFT) Route() string { return RouterKey }

// Type Implements Msg
func (msg MsgTransferNFT) Type() string { return "transfer_nft" }

// ValidateBasic Implements Msg.
func (msg MsgTransferNFT) ValidateBasic() sdk.Error {
	if strings.TrimSpace(msg.Denom) == "" {
		return ErrInvalidCollection(DefaultCodespace)
	}
	if msg.Sender.Empty() {
		return sdk.ErrInvalidAddress("invalid sender address")
	}
	if msg.Recipient.Empty() {
		return sdk.ErrInvalidAddress("invalid recipient address")
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgTransferNFT) GetSignBytes() []byte {
	bz := cdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners Implements Msg.
func (msg MsgTransferNFT) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

// MsgBuyNFT defines a MsgBuyNFT message
type MsgBuyNFT struct {
	Sender sdk.AccAddress
	Owner  sdk.AccAddress
	Amount sdk.Coins
	Denom  string
	ID     uint64
}

// NewMsgBuyNFT is a constructor function for MsgBuyNFT
func NewMsgBuyNFT(sender, recipient sdk.AccAddress, denom string, id uint64,
) MsgBuyNFT {
	return MsgBuyNFT{
		Sender:    sender,
		Recipient: recipient,
		Denom:     denom,
		ID:        id,
	}
}

// Route Implements Msg TODO: fix
// func (msg MsgBuyNFT) Route() string { return RouterKey }

// Type Implements Msg
func (msg MsgBuyNFT) Type() string { return "buy_nft" }

// ValidateBasic Implements Msg.
func (msg MsgBuyNFT) ValidateBasic() sdk.Error {
	if strings.TrimSpace(msg.Denom) == "" {
		return ErrInvalidCollection(DefaultCodespace)
	}
	if msg.Sender.Empty() {
		return sdk.ErrInvalidAddress("invalid sender address")
	}
	if msg.Owner.Empty() {
		return sdk.ErrInvalidAddress("invalid owner address")
	}
	if !msg.Amount.IsValid() {
		return sdk.ErrInvalidCoins()
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgBuyNFT) GetSignBytes() []byte {
	bz := cdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners Implements Msg.
func (msg MsgBuyNFT) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

// MsgEditNFTMetadata edits an NFT's metadata
type MsgEditNFTMetadata struct {
	Owner       sdk.AccAddress
	Name        string
	Description string
	Image       string
	TokenURI    string
}

// NewMsgEditNFTMetadata is a constructor function for MsgSetName
func NewMsgEditNFTMetadata(owner sdk.AccAddress, tokenURI, description, image, name string,
) MsgEditNFTMetadata {
	return MsgEditNFTMetadata{
		Owner:       owner,
		Name:        strings.TrimSpace(name),
		Description: strings.TrimSpace(description),
		Image:       strings.TrimSpace(image),
		TokenURI:    strings.TrimSpace(tokenURI),
	}
}

// Route Implements Msg TODO: fix
// func (msg MsgEditNFTMetadata) Route() string { return RouterKey }

// Type Implements Msg
func (msg MsgEditNFTMetadata) Type() string { return "edit_metadata" }

// ValidateBasic Implements Msg.
func (msg MsgEditNFTMetadata) ValidateBasic() sdk.Error {
	if msg.Owner.Empty() {
		return sdk.ErrInvalidAddress("invalid owner address")
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgEditNFTMetadata) GetSignBytes() []byte {
	bz := cdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners Implements Msg.
func (msg MsgEditNFTMetadata) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

// MsgMintNFT defines a MintNFT message
type MsgMintNFT struct {
	Sender      sdk.AccAddress
	Recipient   sdk.AccAddress
	ID          uint64
	Denom       string
	Name        string
	Description string
	Image       string
	TokenURI    string
}

// NewMsgMintNFT is a constructor function for MsgMintNFT
func NewMsgMintNFT(sender, recipient sdk.AccAddress, id uint64, denom string, name string, description string, image string, tokenURI string,
) MsgMintNFT {
	return MsgMintNFT{
		Sender:      sender,
		Recipient:   recipient,
		ID:          id,
		Denom:       denom,
		Name:        name,
		Description: description,
		Image:       image,
		TokenURI:    tokenURI,
	}
}

// Route Implements Msg TODO: fix
// func (msg MsgMintNFT) Route() string { return RouterKey }

// Type Implements Msg
func (msg MsgMintNFT) Type() string { return "mint_nft" }

// ValidateBasic Implements Msg.
func (msg MsgMintNFT) ValidateBasic() sdk.Error {
	if strings.TrimSpace(msg.Denom) == "" {
		return ErrInvalidCollection(DefaultCodespace)
	}
	if msg.Sender.Empty() {
		return sdk.ErrInvalidAddress("invalid sender address")
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgMintNFT) GetSignBytes() []byte {
	bz := cdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners Implements Msg.
func (msg MsgMintNFT) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

// MsgBurnNFT defines a BurnNFT message
type MsgBurnNFT struct {
	Sender sdk.AccAddress
	ID     uint64
	Denom  string
}

// NewMsgBurnNFT is a constructor function for MsgBurnNFT
func NewMsgBurnNFT(sender sdk.AccAddress, id uint64, denom string,
) MsgBurnNFT {
	return MsgBurnNFT{
		Sender: sender,
		ID:     id,
		Denom:  denom,
	}
}

// Route Implements Msg TODO: fix
// func (msg MsgBurnNFT) Route() string { return RouterKey }

// Type Implements Msg
func (msg MsgBurnNFT) Type() string { return "burn_nft" }

// ValidateBasic Implements Msg.
func (msg MsgBurnNFT) ValidateBasic() sdk.Error {
	if strings.TrimSpace(msg.Denom) == "" {
		return ErrInvalidCollection(DefaultCodespace)
	}
	if msg.Sender.Empty() {
		return sdk.ErrInvalidAddress("invalid sender address")
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgBurnNFT) GetSignBytes() []byte {
	bz := cdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners Implements Msg.
func (msg MsgBurnNFT) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}
