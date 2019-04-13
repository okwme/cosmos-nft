package app

import (
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/tendermint/tendermint/libs/log"

	bam "github.com/cosmos/cosmos-sdk/baseapp"
	dbm "github.com/tendermint/tendermint/libs/db"
)

const (
	appName = "nfts"
)

type nftsApp struct {
	*bam.BaseApp
}

func NewNftsApp(logger log.Logger, db dbm.DB) *nftsApp {

	cdc := MakeCodec()

	bApp := bam.NewBaseApp(appName, logger, db, auth.DefaultTxDecoder(cdc))

	var app = &nftsApp{
		BaseApp: bApp,
		cdc:     cdc,
	}

	return app
}
