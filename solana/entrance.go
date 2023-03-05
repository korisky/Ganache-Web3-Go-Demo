package main

import (
	"github.com/portto/solana-go-sdk/client"
	"quickNode/solana/accounts"
	"quickNode/solana/balance"
)

var (
	// own address account
	address = "AnayTW335MabjhtXTJeBit5jdLhNeUVBVPXeRKCid79D"

	// custom connection would be:
	// cli := client.NewClient(rpc.DevnetRPCEndpoint, rpc.WithHTTPClient(customClietn))
	cli = client.NewClient("https://solana-devnet.g.alchemy.com/v2/On35d8LdFc1QGYD-wCporecGj359qian")
)

func main() {
	// accounts related
	account, err := accounts.TryRecoverAccount("")
	if err == nil {
		balance.TryRequestAirdrop(cli, account.PublicKey.ToBase58())
	}
}
