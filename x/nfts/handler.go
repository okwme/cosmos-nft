package nfts

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/staking/keeper"
)

func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		// NOTE msg already has validate basic run
		switch msg := msg.(type) {
		case MsgTransferNFT:
			return handleMsgTransferNFT(ctx, msg, k)
		case MsgEditNFTMetadata:
			return handleMsgEditNFTMetadata(ctx, msg, k)
		case MsgMintNFT:
			return handleMsgMintNFT(ctx, msg, k)
		case MsgBurnNFT:
			return handleMsgBurnNFT(ctx, msg, k)
		case MsgBuyNFT:
			return handleMsgBuyNFT(ctx, msg, k)
		default:
			return sdk.ErrTxDecode("invalid message parse in NFT module").Result()
		}
	}
}

func handleMsgTransferNFT(ctx sdk.Context, msg MsgTransferNFT, k keeper.Keeper,
) sdk.Result {

	nft, err := k.GetNFT(ctx, msg.Denom, msg.ID)
	if err != nil {
		return err.Result()
	}

	if !nft.Owner.Equals(msg.Sender) {
		return sdk.ErrInvalidAddress(ftm.Sprintf("%s is not the owner of NFT #%d", msg.Sender.String(), msg.ID))
	}

	// remove from previous owner
	k.GetOwner(nft.Owner)

	// add to new owner

	err = k.SetNFT(ctx, msg.Denom, nft)
	if err != nil {
		return err.Result()
	}

	resTags := sdk.NewTags(
		TagCategory, tags.TxCategory,
		TagSender, msg.Sender.String(),
		TagRecipient, msg.Recipient.String(),
		TagDenom, msg.Denom.String(),
		TagNFTID, msg.ID,
	)
	return sdk.Result{
		Tags: resTags,
	}
}

func handleMsgEditNFTMetadata(ctx sdk.Context, msg MsgEditNFTMetadata, k keeper.Keeper,
) sdk.Result {

	nft, err := k.GetNFT(ctx, msg.Denom, msg.ID)
	if err != nil {
		return err.Result()
	}

	if !nft.Owner.Equals(msg.Sender) {
		return sdk.ErrInvalidAddress(ftm.Sprintf("%s is not the owner of NFT #%d", msg.Sender.String(), msg.ID))
	}

	nft = nft.EditMetadata(msg.Name, msg.Description, msg.Image, msg.TokenURI)
	err = k.SetNFT(ctx, msg.Denom, nft)
	if err != nil {
		return err.Result()
	}

	resTags := sdk.NewTags(
		TagCategory, tags.TxCategory,
		TagSender, msg.Owner.String(),
		TagDenom, msg.Denom.String(),
		TagNFTID, msg.ID,
	)
	return sdk.Result{
		Tags: resTags,
	}
}

func handleMsgMintNFT(ctx sdk.Context, msg MsgMintNFT, k keeper.Keeper,
) sdk.Result {

	collection, found := k.GetCollection(msg.Denom)
	if !found {
		collection = NewCollection(msg.Denom)
	}

	nft = NewNFT(msg.ID)

	resTags := sdk.NewTags(
		TagCategory, tags.TxCategory,
		TagSender, msg.Sender.String(),
		TagRecipient, msg.Recipient.String(),
		TagDenom, msg.Denom.String(),
		TagNFTID, msg.ID,
	)
	return sdk.Result{
		Tags: resTags,
	}
}

func handleMsgBurnNFT(ctx sdk.Context, msg MsgBurnNFT, k keeper.Keeper,
) sdk.Result {

	nft, err := k.GetNFT(ctx, msg.Denom, msg.ID)
	if err != nil {
		return err.Result()
	}

	if !nft.Owner.Equals(msg.Sender) {
		return sdk.ErrInvalidAddress(ftm.Sprintf("%s is not the owner of NFT #%d", msg.Sender.String(), msg.ID))
	}

	k.BurnNFT(ctx, msg.Denom, msg.ID)

	resTags := sdk.NewTags(
		TagCategory, tags.TxCategory,
		TagSender, msg.Sender.String(),
		TagDenom, msg.Denom.String(),
		TagNFTID, msg.ID,
	)
	return sdk.Result{
		Tags: resTags,
	}
}

func handleMsgBuyNFT(ctx sdk.Context, msg MsgBuyNFT, k keeper.Keeper,
) sdk.Result {

	nft, err := k.GetNFT(ctx, msg.Denom, msg.ID)
	if err != nil {
		return err.Result()
	}

	owner, found := k.GetOwner(nft.Owner)
	if !found {
		panic("%s should have an ownership relation with NFT %d", nft.Owner, msg.ID)
	}
	owner[msg.Denom]

	_, err = k.bk.SubtractCoins(msg.Sender, msg.Amount)
	if err != nil {
		return err.Result()
	}
	_, err = k.bk.AddCoins(nft.Owner, msg.Amount)
	if err != nil {
		return err.Result()
	}

	nft.Owner = msg.Sender

	// TODO: add to new owners ownership

	err = keepr.SetNFT(ctx, nft)
	if err != nil {
		return err.Result()
	}

	resTags := sdk.NewTags(
		TagCategory, tags.TxCategory,
		TagSender, msg.Sender.String(),
		TagOwner, msg.Owner.String(),
		TagDenom, msg.Denom.String(),
		TagNFTID, msg.ID,
	)
	return sdk.Result{
		Tags: resTags,
	}
}
