package cli

import (
	"github.com/spf13/viper"

	"github.com/cosmos/cosmos-sdk/x/nft"
)

const (
	flagName        = "name"
	flagDescription = "description"
	flagImage       = "image"
	flagTokenURI    = "tokenURI"
)

func parseEditMetadataFlags() (nft.MsgEditNFTMetadata, error) {
	msg := nft.MsgEditNFTMetadata{}

	msg.EditName := viper.IsSet(flagName)
	msg.Name := viper.GetString(flagName)

	msg.EditDescription := viper.IsSet(flagDescription)
	msg.Description := viper.GetString(flagDescription)

	msg.EditImage := viper.IsSet(flagImage)
	msg.Image := viper.GetString(flagImage)

	msg.EditTokenURI := viper.IsSet(flagTokenURI)
	msg.TokenURI := viper.GetString(flagTokenURI)

	return msg, nil
}
