package main

import (
	"github.com/portto/solana-go-sdk/client"
	"github.com/portto/solana-go-sdk/rpc"
	"web3Demo/solana/httpProxy"
	"web3Demo/solana/portto/balance"
)

var (
	// own accountAddress account
	accountAddress   = "AnayTW335MabjhtXTJeBit5jdLhNeUVBVPXeRKCid79D"
	tokenMintAddress = "Gh9ZwEmdLJ8DscKNTkTqPbNwLNNBjuSzaG9Vp2KGtKJr"
	ownEndpoint      = "https://solana-devnet.g.alchemy.com/v2/On35d8LdFc1QGYD-wCporecGj359qian"
	// custom connection would be:
	cli = client.New(rpc.WithEndpoint(ownEndpoint), rpc.WithHTTPClient(httpProxy.GetHttpClient()))
)

// main are following stuff from https://yihau.gitbook.io/solana-go/tour/create-token-account/associated-token-account
func main() {
	balance.TryGetBalance(cli, accountAddress)
	balance.TryGetAssociatedTokenAddressBalance(cli, accountAddress, tokenMintAddress)
}
