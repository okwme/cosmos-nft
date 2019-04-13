package nfts

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// CodeType definition
type CodeType = sdk.CodeType

// NFT error code
const (
	DefaultCodespace sdk.CodespaceType = ModuleName

	CodeInvalidCollection CodeType = 666
)

// ErrInvalidCollection is an error
func ErrInvalidCollection(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidCollection, "Invalid denom provided while dealing with NFT collection")
}
