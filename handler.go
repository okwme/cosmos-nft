package app

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/nft"
)

// OverrideNFTModule overrides the NFT module for custom handlers
type OverrideNFTModule struct {
	nft.AppModule
	k nft.Keeper
}

// NewHandler module handler for the OerrideNFTModule
func (am OverrideNFTModule) NewHandler() sdk.Handler {
	return CustomNFTHandler(am.k)
}

// NewOverrideNFTModule generates a new NFT Module
func NewOverrideNFTModule(appModule nft.AppModule, keeper nft.Keeper) OverrideNFTModule {
	return OverrideNFTModule{
		AppModule: appModule,
		k:         keeper,
	}
}

// CustomNFTHandler routes the messages to the handlers
func CustomNFTHandler(k nft.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case nft.MsgTransferNFT:
			return nft.HandleMsgTransferNFT(ctx, msg, k)
		case nft.MsgEditNFTMetadata:
			return nft.HandleMsgEditNFTMetadata(ctx, msg, k)
		case nft.MsgMintNFT:
			return HandleMsgMintNFTCustom(ctx, msg, k)
		case nft.MsgBurnNFT:
			return nft.HandleMsgBurnNFT(ctx, msg, k)
		default:
			errMsg := fmt.Sprintf("unrecognized nft message type: %T", msg)
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

// HandleMsgMintNFTCustom handles MsgMintNFT
func HandleMsgMintNFTCustom(ctx sdk.Context, msg nft.MsgMintNFT, k nft.Keeper,
) sdk.Result {

	isTwilight := checkTwilight(ctx)

	if isTwilight {
		return nft.HandleMsgMintNFT(ctx, msg, k)
	}

	errMsg := fmt.Sprintf("Can't mint astral bodies outside of twilight!")
	return sdk.ErrUnknownRequest(errMsg).Result()
}

func checkTwilight(ctx sdk.Context) bool {
	header := ctx.BlockHeader()
	time := header.Time
	fmt.Println("time", time)
	return true
}
