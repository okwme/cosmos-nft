package client

import (
	"github.com/cosmos/cosmos-sdk/client"
	cosmiccmd "github.com/okwme/cosmos-nft/x/cosmic/client"
	"github.com/spf13/cobra"

	amino "github.com/tendermint/go-amino"
)

// ModuleClient exports all client functionality from this module
type ModuleClient struct {
	storeKey string
	cdc      *amino.Codec
}

// NewModuleClient creates a new module client
func NewModuleClient(storeKey string, cdc *amino.Codec) ModuleClient {
	return ModuleClient{storeKey, cdc}
}

// GetTxCmd returns the transaction commands for this module
func (mc ModuleClient) GetTxCmd() *cobra.Command {
	nftTxCmd := &cobra.Command{
		Use:   "nft-cosmic",
		Short: "Cosmic Non-Fungible Token transactions subcommands",
	}

	nftTxCmd.AddCommand(client.PostCommands(
		cosmiccmd.GetCmdEditNFTMetadata(mc.cdc),
		cosmiccmd.GetCmdMintNFT(mc.cdc),
		cosmiccmd.GetCmdBurnNFT(mc.cdc),
	)...)

	return nftTxCmd
}
