#!/bin/bash
rm -rf ~/.nft*
nftd init cosmic --chain-id cosmic-chain
nftcli keys add billy
nftcli config indent true
nftcli config output json
nftcli config trust-node true
nftcli config chain-id cosmic-chain
nftd add-genesis-account $(nftcli keys show billy -a) 100000000stake
# nftd gentx --name billy   
echo "1234567890" | nftd gentx --name billy
nftd collect-gentxs                                                 
nftd validate-genesis                                               
nftd start                                                          