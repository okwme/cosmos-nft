package nftmarket

// import (
// 	"fmt"

// 	sdk "github.com/cosmos/cosmos-sdk/types"
// 	"github.com/cosmos/cosmos-sdk/x/nft/keeper"
// 	"github.com/cosmos/cosmos-sdk/x/nft/types"
// )

// // NewHandler routes the messages to the handlers
// func NewHandler(k keeper.Keeper) sdk.Handler {
// 	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
// 		switch msg := msg.(type) {
// 		case MsgBuyNFT:
// 			return handleMsgBuyNFT(ctx, msg, k)
// 		default:
// 			errMsg := fmt.Sprintf("unrecognized cosmic-nft message type: %T", msg)
// 			return sdk.ErrUnknownRequest(errMsg).Result()
// 		}
// 	}
// }

// func handleMsgBuyNFT(ctx sdk.Context, msg types.MsgBuyNFT, k keeper.Keeper,
// ) sdk.Result {

// 	nft, err := k.GetNFT(ctx, msg.Denom, msg.ID)
// 	if err != nil {
// 		return err.Result()
// 	}

// 	owner, found := k.GetOwner(ctx, nft.Owner)
// 	if !found {
// 		panic(fmt.Sprintf("%s should have an ownership relation with NFT %d", nft.Owner, msg.ID))
// 	}
// 	// owner[msg.Denom]

// 	_, err = k.bk.SubtractCoins(msg.Sender, msg.Amount)
// 	if err != nil {
// 		return err.Result()
// 	}
// 	_, err = k.bk.AddCoins(nft.Owner, msg.Amount)
// 	if err != nil {
// 		return err.Result()
// 	}

// 	nft.Owner = msg.Sender

// 	// TODO: add to new owners ownership

// 	err = k.SetNFT(ctx, nft)
// 	if err != nil {
// 		return err.Result()
// 	}

// 	resTags := sdk.NewTags(
// 		tags.Category, tags.TxCategory,
// 		tags.Sender, msg.Sender.String(),
// 		tags.Owner, msg.Owner.String(),
// 		tags.Denom, msg.Denom.String(),
// 		tags.NFTID, msg.ID,
// 	)
// 	return sdk.Result{
// 		Tags: resTags,
// 	}
// }
