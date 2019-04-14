package nfts

import (
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/codec"
)

// query endpoints supported by the nft Querier
const (
	// QueryDenoms      = "denoms"
	// QueryTotalSupply = "totalSupply"
	// QueryIDs         = "ids"
	QueryBalanceOf = "balanceOf"
	QueryOwnerOf   = "ownerOf"
	QueryMetadata  = "metadata"
)

// NewQuerier is the module level router for state queries
func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		// case QueryDenoms:
		// 	return queryDenoms(ctx, req, keeper)
		// case QueryTotalSupply:
		// 	return queryTotalSupply(ctx, path[1:], req, keeper)
		// case QueryIDs:
		// 	return queryIDs(ctx, path[1:], req, keeper)

		case QueryBalanceOf:
			return queryBalanceOf(ctx, path[1:], req, keeper)
		case QueryOwnerOf:
			return queryOwnerOf(ctx, path[1:], req, keeper)
		case QueryMetadata:
			return queryMetadata(ctx, path[1:], req, keeper)
		default:
			return nil, sdk.ErrUnknownRequest("unknown nft query endpoint")
		}
	}
}

// nolint: unparam
func queryBalanceOf(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) (res []byte, err sdk.Error) {
	denom := strings.TrimSpace(path[0])
	owner := strings.TrimSpace(path[1])

	collection, found := keeper.GetCollection(ctx, denom)

	if !found {
		return []byte{}, ErrUnknownCollection
	}
	balance := 0
	for k, v := range collection.NFTs {
		fmt.Println("k:", k, "v:", v)
		if v.Owner.String() == owner {
			balance++
		}
	}

	bz, err2 := codec.MarshalJSONIndent(keeper.cdc, QueryResResolve{balance})
	if err2 != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}

// // Query Result Payload for a resolve query
// type QueryResResolve struct {
// 	Value string `json:"value"`
// }

// // implement fmt.Stringer
// func (r QueryResResolve) String() string {
// 	return r.Value
// }

// // nolint: unparam
// func queryWhois(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) (res []byte, err sdk.Error) {
// 	name := path[0]

// 	whois := keeper.GetWhois(ctx, name)

// 	bz, err2 := codec.MarshalJSONIndent(keeper.cdc, whois)
// 	if err2 != nil {
// 		panic("could not marshal result to JSON")
// 	}

// 	return bz, nil
// }

// // implement fmt.Stringer
// func (w Whois) String() string {
// 	return strings.TrimSpace(fmt.Sprintf(`Owner: %s
// Value: %s
// Price: %s`, w.Owner, w.Value, w.Price))
// }

// func queryNames(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) (res []byte, err sdk.Error) {
// 	var namesList QueryResNames

// 	iterator := keeper.GetNamesIterator(ctx)

// 	for ; iterator.Valid(); iterator.Next() {
// 		name := string(iterator.Key())
// 		namesList = append(namesList, name)
// 	}

// 	bz, err2 := codec.MarshalJSONIndent(keeper.cdc, namesList)
// 	if err2 != nil {
// 		panic("could not marshal result to JSON")
// 	}

// 	return bz, nil
// }

// // Query Result Payload for a names query
// type QueryResNames []string

// // implement fmt.Stringer
// func (n QueryResNames) String() string {
// 	return strings.Join(n[:], "\n")
// }
