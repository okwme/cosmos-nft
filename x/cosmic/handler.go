package cosmic

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/nft/keeper"
)

// NewHandler routes the messages to the handlers
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case MsgEditNFTMetadata:
			return HandleMsgEditNFTMetadata(ctx, msg, k)
		case MsgMintNFT:
			return HandleMsgMintNFT(ctx, msg, k)
		case MsgBurnNFT:
			return HandleMsgBurnNFT(ctx, msg, k)
		default:
			errMsg := fmt.Sprintf("unrecognized cosmic-nft message type: %T", msg)
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}
