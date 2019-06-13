module github.com/okwme/cosmos-nft

go 1.12

require (
	github.com/cosmos/cosmos-sdk v0.28.2-0.20190429233659-dcb7e04cb2b6
	github.com/mattn/go-isatty v0.0.7 // indirect
	github.com/prometheus/procfs v0.0.0-20190328153300-af7bedc223fb // indirect
	github.com/spf13/afero v1.2.2 // indirect
	github.com/syndtr/goleveldb v1.0.0 // indirect
	github.com/tendermint/tendermint v0.31.5
	golang.org/x/sys v0.0.0-20190329044733-9eb1bfa1ce65 // indirect
	google.golang.org/genproto v0.0.0-20190327125643-d831d65fe17d // indirect
	google.golang.org/grpc v1.19.1 // indirect
)

replace golang.org/x/crypto => github.com/tendermint/crypto v0.0.0-20180820045704-3764759f34a5

replace github.com/cosmos/cosmos-sdk => /home/billyrennekamp/go/src/github.com/cosmos/cosmos-sdk
