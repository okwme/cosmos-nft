package nfts

import sdk "github.com/cosmos/cosmos-sdk/types"

// GenesisState is the bank state that must be provided at genesis.
type GenesisState struct {
	Collections []Collection `json:"Collections"`
}

// NewGenesisState creates a new genesis state.
func NewGenesisState(collections []Collection) GenesisState {
	return GenesisState{
		Collections: collections,
	}
}

// DefaultGenesisState returns a default genesis state
func DefaultGenesisState() GenesisState {
	return NewGenesisState([]Collection{})
}

// InitGenesis sets distribution information for genesis.
func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) {
	for _, collection := range data.Collections {
		keeper.SetCollection(ctx, collection.Denom, collection)
	}
}

// ExportGenesis returns a GenesisState for a given context and keeper.
func ExportGenesis(ctx sdk.Context, keeper Keeper) GenesisState {
	return NewGenesisState(keeper.GetCollections(ctx))
}

// ValidateGenesis performs basic validation of nfts genesis data returning an
// error for any failed validation criteria.
func ValidateGenesis(data GenesisState) error {
	for _, collection := range data.Collections {
		err := collection.ValidateBasic()
		if err != nil {
			return err
		}
	}
	return nil
}
