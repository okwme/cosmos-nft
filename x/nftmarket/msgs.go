package nftmarket

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

/* --------------------------------------------------------------------------- */
// MsgBuyNFT
/* --------------------------------------------------------------------------- */

// MsgBuyNFT defines a MsgBuyNFT message
type MsgBuyNFT struct {
	Sender sdk.AccAddress
	Amount sdk.Coins
	Denom  string
	ID     uint64
}

// NewMsgBuyNFT is a constructor function for MsgBuyNFT
func NewMsgBuyNFT(sender, owner sdk.AccAddress, denom string, id uint64) MsgBuyNFT {
	return MsgBuyNFT{
		Sender: sender,
		Denom:  strings.TrimSpace(denom),
		ID:     id,
	}
}

// Route Implements Msg
func (msg MsgBuyNFT) Route() string { return RouterKey }

// Type Implements Msg
func (msg MsgBuyNFT) Type() string { return "buy_nft" }

// ValidateBasic Implements Msg.
func (msg MsgBuyNFT) ValidateBasic() sdk.Error {
	if msg.Denom == "" {
		return ErrInvalidCollection(DefaultCodespace)
	}
	if msg.Sender.Empty() {
		return sdk.ErrInvalidAddress("invalid sender address")
	}
	if !msg.Amount.IsValid() {
		return sdk.ErrInvalidCoins("invalid amount provided")
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
