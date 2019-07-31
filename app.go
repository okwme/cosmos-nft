package app

import (
	"encoding/json"
	"fmt"
	"os"

	abci "github.com/tendermint/tendermint/abci/types"
	cmn "github.com/tendermint/tendermint/libs/common"
	dbm "github.com/tendermint/tendermint/libs/db"
	"github.com/tendermint/tendermint/libs/log"
	tmtypes "github.com/tendermint/tendermint/types"

	bam "github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	distr "github.com/cosmos/cosmos-sdk/x/distribution"
	"github.com/cosmos/cosmos-sdk/x/genaccounts"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	"github.com/cosmos/cosmos-sdk/x/mint"
	"github.com/cosmos/cosmos-sdk/x/nft"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/cosmos/cosmos-sdk/x/supply"
)

const appName = "cosmic"

var (
	// DefaultCLIHome is the default home directories for the application CLI
	DefaultCLIHome = os.ExpandEnv("$HOME/.nftcli")

	// DefaultNodeHome sets the folder where the applcation data and configuration will be stored
	DefaultNodeHome = os.ExpandEnv("$HOME/.nftd")

	// ModuleBasics is the module manager in charge of setting up basic module elements
	ModuleBasics = module.NewBasicManager(
		genaccounts.AppModuleBasic{},
		genutil.AppModuleBasic{},
		auth.AppModuleBasic{},
		bank.AppModuleBasic{},
		params.AppModuleBasic{},
		staking.AppModuleBasic{},
		distr.AppModuleBasic{},
		slashing.AppModuleBasic{},
		supply.AppModuleBasic{},
		nft.AppModuleBasic{},
	)
)

// MakeCodec generates the necessary codecs for Amino
func MakeCodec() *codec.Codec {
	var cdc = codec.New()
	ModuleBasics.RegisterCodec(cdc)
	sdk.RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)
	return cdc
}

type cosmicApp struct {
	*bam.BaseApp
	cdc *codec.Codec

	// Keys to access the substores
	keyMain     *sdk.KVStoreKey
	keyAccount  *sdk.KVStoreKey
	keySupply   *sdk.KVStoreKey
	keyStaking  *sdk.KVStoreKey
	tkeyStaking *sdk.TransientStoreKey
	keyDistr    *sdk.KVStoreKey
	tkeyDistr   *sdk.TransientStoreKey
	keyNFT      *sdk.KVStoreKey
	keyParams   *sdk.KVStoreKey
	tkeyParams  *sdk.TransientStoreKey
	keySlashing *sdk.KVStoreKey

	// Keepers
	accountKeeper  auth.AccountKeeper
	bankKeeper     bank.Keeper
	supplyKeeper   supply.Keeper
	stakingKeeper  staking.Keeper
	slashingKeeper slashing.Keeper
	distrKeeper    distr.Keeper
	paramsKeeper   params.Keeper
	nftKeeper      nft.Keeper

	// Module Manager
	mm *module.Manager
}

// NewCosmicApp is a constructor function for nftApp
func NewCosmicApp(logger log.Logger, db dbm.DB) *cosmicApp {

	// First define the top level codec that will be shared by the different modules
	cdc := MakeCodec()

	// BaseApp handles interactions with Tendermint through the ABCI protocol
	bApp := bam.NewBaseApp(appName, logger, db, auth.DefaultTxDecoder(cdc))

	// Here you initialize your application with the store keys it requires
	var app = &cosmicApp{
		BaseApp: bApp,
		cdc:     cdc,

		keyMain:     sdk.NewKVStoreKey(bam.MainStoreKey),
		keyAccount:  sdk.NewKVStoreKey(auth.StoreKey),
		keySupply:   sdk.NewKVStoreKey(supply.StoreKey),
		keyStaking:  sdk.NewKVStoreKey(staking.StoreKey),
		tkeyStaking: sdk.NewTransientStoreKey(staking.TStoreKey),
		keyDistr:    sdk.NewKVStoreKey(distr.StoreKey),
		tkeyDistr:   sdk.NewTransientStoreKey(distr.TStoreKey),
		keyNFT:      sdk.NewKVStoreKey(nft.StoreKey),
		keyParams:   sdk.NewKVStoreKey(params.StoreKey),
		tkeyParams:  sdk.NewTransientStoreKey(params.TStoreKey),
		keySlashing: sdk.NewKVStoreKey(slashing.StoreKey),
	}

	// The ParamsKeeper handles parameter storage for the application
	app.paramsKeeper = params.NewKeeper(app.cdc, app.keyParams, app.tkeyParams, params.DefaultCodespace)
	// Set specific supspaces
	authSubspace := app.paramsKeeper.Subspace(auth.DefaultParamspace)
	bankSupspace := app.paramsKeeper.Subspace(bank.DefaultParamspace)
	stakingSubspace := app.paramsKeeper.Subspace(staking.DefaultParamspace)
	distrSubspace := app.paramsKeeper.Subspace(distr.DefaultParamspace)
	slashingSubspace := app.paramsKeeper.Subspace(slashing.DefaultParamspace)

	// The AccountKeeper handles address -> account lookups
	app.accountKeeper = auth.NewAccountKeeper(
		app.cdc,
		app.keyAccount,
		authSubspace,
		auth.ProtoBaseAccount,
	)

	// The BankKeeper allows you perform sdk.Coins interactions
	app.bankKeeper = bank.NewBaseKeeper(
		app.accountKeeper,
		bankSupspace,
		bank.DefaultCodespace,
	)

	maccPerms := map[string][]string{
		auth.FeeCollectorName:     nil,
		distr.ModuleName:          nil,
		nft.ModuleName:            nil,
		mint.ModuleName:           {supply.Minter},
		staking.BondedPoolName:    {supply.Burner, supply.Staking},
		staking.NotBondedPoolName: {supply.Burner, supply.Staking},
	}

	app.supplyKeeper = supply.NewKeeper(app.cdc, app.keySupply, app.accountKeeper,
		app.bankKeeper, supply.DefaultCodespace, maccPerms)

	// The staking keeper
	stakingKeeper := staking.NewKeeper(
		app.cdc,
		app.keyStaking,
		app.tkeyStaking,
		app.supplyKeeper,
		stakingSubspace,
		staking.DefaultCodespace,
	)

	app.distrKeeper = distr.NewKeeper(
		app.cdc,
		app.keyDistr,
		distrSubspace,
		&stakingKeeper,
		app.supplyKeeper,
		distr.DefaultCodespace,
		auth.FeeCollectorName,
	)

	app.slashingKeeper = slashing.NewKeeper(
		app.cdc,
		app.keySlashing,
		&stakingKeeper,
		slashingSubspace,
		slashing.DefaultCodespace,
	)

	// register the staking hooks
	// NOTE: stakingKeeper above is passed by reference, so that it will contain these hooks
	app.stakingKeeper = *stakingKeeper.SetHooks(
		staking.NewMultiStakingHooks(
			app.distrKeeper.Hooks(),
			app.slashingKeeper.Hooks()),
	)

	// The NFTKeeper is the Keeper from the module NFTs
	// It handles interactions with the nftstore
	app.nftKeeper = nft.NewKeeper(
		app.cdc,
		app.keyNFT,
	)

	nftModule := nft.NewAppModule(app.nftKeeper)
	overriddenNFTModule := NewOverrideNFTModule(nftModule, app.nftKeeper)

	app.mm = module.NewManager(
		genaccounts.NewAppModule(app.accountKeeper),
		genutil.NewAppModule(app.accountKeeper, app.stakingKeeper, app.BaseApp.DeliverTx),
		auth.NewAppModule(app.accountKeeper),
		bank.NewAppModule(app.bankKeeper, app.accountKeeper),
		overriddenNFTModule,
		distr.NewAppModule(app.distrKeeper, app.supplyKeeper),
		slashing.NewAppModule(app.slashingKeeper, app.stakingKeeper),
		staking.NewAppModule(app.stakingKeeper, app.distrKeeper, app.accountKeeper, app.supplyKeeper),
	)

	app.mm.SetOrderBeginBlockers(distr.ModuleName, slashing.ModuleName)
	app.mm.SetOrderEndBlockers(staking.ModuleName)

	// Sets the order of Genesis - Order matters, genutil is to always come last
	app.mm.SetOrderInitGenesis(
		genaccounts.ModuleName,
		distr.ModuleName,
		staking.ModuleName,
		auth.ModuleName,
		bank.ModuleName,
		slashing.ModuleName,
		nft.ModuleName,
		genutil.ModuleName,
	)

	// register all module routes and module queriers
	app.mm.RegisterRoutes(app.Router(), app.QueryRouter())

	// The initChainer handles translating the genesis.json file into initial state for the network
	app.SetInitChainer(app.InitChainer)
	app.SetBeginBlocker(app.BeginBlocker)
	app.SetEndBlocker(app.EndBlocker)

	// The AnteHandler handles signature verification and transaction pre-processing
	app.SetAnteHandler(
		auth.NewAnteHandler(
			app.accountKeeper,
			app.supplyKeeper,
			auth.DefaultSigVerificationGasConsumer,
		),
	)
	fmt.Println("keyMain", app.keyMain.Name())
	fmt.Println("keyAccount", app.keyAccount.Name())
	fmt.Println("keySupply", app.keySupply.Name())
	fmt.Println("keyStaking", app.keyStaking.Name())
	fmt.Println("tkeyStaking", app.tkeyStaking.Name())
	fmt.Println("keyDistr", app.keyDistr.Name())
	fmt.Println("keySlashing", app.keySlashing.Name())
	fmt.Println("keyNFT", app.keyNFT.Name())
	fmt.Println("keyParams", app.keyParams.Name())
	fmt.Println("tkeyParams", app.tkeyParams.Name())
	app.MountStores(
		app.keyMain,
		app.keyAccount,
		app.keySupply,
		app.keyStaking,
		app.tkeyStaking,
		app.keyDistr,
		app.tkeyDistr,
		app.keySlashing,
		app.keyNFT,
		app.keyParams,
		app.tkeyParams,
	)

	err := app.LoadLatestVersion(app.keyMain)
	if err != nil {
		cmn.Exit(err.Error())
	}

	return app
}

// GenesisState represents chain state at the start of the chain. Any initial state (account balances) are stored here.
type GenesisState map[string]json.RawMessage

// NewDefaultGenesisState creates a new default genesis file
func NewDefaultGenesisState() GenesisState {
	return ModuleBasics.DefaultGenesis()
}

func (app *cosmicApp) InitChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	var genesisState GenesisState

	err := app.cdc.UnmarshalJSON(req.AppStateBytes, &genesisState)
	if err != nil {
		panic(err)
	}

	return app.mm.InitGenesis(ctx, genesisState)
}

func (app *cosmicApp) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {
	return app.mm.BeginBlock(ctx, req)
}
func (app *cosmicApp) EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {
	return app.mm.EndBlock(ctx, req)
}
func (app *cosmicApp) LoadHeight(height int64) error {
	return app.LoadVersion(height, app.keyMain)
}

//_________________________________________________________
func (app *cosmicApp) ExportAppStateAndValidators(forZeroHeight bool, jailWhiteList []string,
) (appState json.RawMessage, validators []tmtypes.GenesisValidator, err error) {

	// as if they could withdraw from the start of the next block
	ctx := app.NewContext(true, abci.Header{Height: app.LastBlockHeight()})

	genState := app.mm.ExportGenesis(ctx)
	appState, err = codec.MarshalJSONIndent(app.cdc, genState)
	if err != nil {
		return nil, nil, err
	}

	validators = staking.WriteValidators(ctx, app.stakingKeeper)

	return appState, validators, nil
}
