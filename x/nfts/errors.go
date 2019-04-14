package nfts

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// CodeType definition
type CodeType = sdk.CodeType

// NFT error code
const (
	DefaultCodespace sdk.CodespaceType = ModuleName

	CodeUnknownCollection CodeType = 651
	CodeInvalidNFT CodeType = 652
	CodeUnknownNFT CodeType = 653
)

// ErrUnknownCollection is an error
func ErrUnknownCollection(codespace sdk.CodespaceType, msg string) sdk.Error {
	if msg != "" {
		return sdk.NewError(codespace, CodeUnknownCollection, msg)
	}
	return sdk.NewError(codespace, CodeUnknownCollection, "unknown NFT collection")
}

// ErrInvalidNFT is an error
func ErrInvalidNFT(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidNFT, "invalid NFT")
}

// ErrUnknownNFT is an error
func ErrUnknownNFT(codespace sdk.CodespaceType, msg string) sdk.Error {
	if msg != "" {
		return sdk.NewError(codespace, CodeUnknownNFT, msg)
	}
	return sdk.NewError(codespace, CodeUnknownNFT, "unknown NFT")
}