package main

import (
	"web3Demo/solana/gagliardetto/balance"
	"web3Demo/solana/httpProxy"
)

var (
	// own accountAddress account
	accountAddress   = "AnayTW335MabjhtXTJeBit5jdLhNeUVBVPXeRKCid79D"
	tokenMintAddress = "Gh9ZwEmdLJ8DscKNTkTqPbNwLNNBjuSzaG9Vp2KGtKJr"
	ownEndpoint      = "https://solana-devnet.g.alchemy.com/v2/On35d8LdFc1QGYD-wCporecGj359qian"
	cli              = httpProxy.NewRPC(ownEndpoint)
)

func main() {
	balance.TryGetTokenAmount(cli, accountAddress, tokenMintAddress)
}
