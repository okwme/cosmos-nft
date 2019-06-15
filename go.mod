module github.com/okwme/cosmos-nft

go 1.12

require (
	github.com/cosmos/cosmos-sdk v0.28.2-0.20190517070908-8ff9b25facc5
	github.com/cosmos/gaia v0.0.0
	github.com/spf13/cobra v0.0.3
	github.com/spf13/viper v1.3.2
	github.com/tendermint/go-amino v0.15.0
	github.com/tendermint/tendermint v0.31.5
	google.golang.org/genproto v0.0.0-20190327125643-d831d65fe17d // indirect
)

replace golang.org/x/crypto => github.com/tendermint/crypto v0.0.0-20180820045704-3764759f34a5

replace github.com/cosmos/cosmos-sdk => /root/GitHub.com/cosmos/cosmos-sdk
