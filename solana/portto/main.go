package main

import (
	"github.com/portto/solana-go-sdk/client"
	"github.com/portto/solana-go-sdk/rpc"
	"web3Demo/solana/httpProxy"
)

var (
	// own accountAddress account
	accountAddress   = "AnayTW335MabjhtXTJeBit5jdLhNeUVBVPXeRKCid79D"
	tokenMintAddress = "Gh9ZwEmdLJ8DscKNTkTqPbNwLNNBjuSzaG9Vp2KGtKJr"
	ownEndpoint      = "https://solana-devnet.g.alchemy.com/v2/On35d8LdFc1QGYD-wCporecGj359qian"
	// custom connection would be:
	cli = client.New(rpc.WithEndpoint(ownEndpoint), rpc.WithHTTPClient(httpProxy.GetHttpClient()))
)

// main are following stuff from https://portto.github.io/solana-go-sdk/tour/token-transfer.html
func main() {
	//balance.TryGetBalance(cli, accountAddress)
	//balance.TryGetAssociatedTokenAddressBalance(cli, accountAddress, tokenMintAddress)
}
