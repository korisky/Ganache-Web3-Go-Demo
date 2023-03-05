package main

import (
	"github.com/portto/solana-go-sdk/client"
)

func main() {
	// own address account
	address := "AnayTW335MabjhtXTJeBit5jdLhNeUVBVPXeRKCid79D"

	// custom connection would be:
	// cli := client.NewClient(rpc.DevnetRPCEndpoint, rpc.WithHTTPClient(customClietn))
	cli := client.NewClient("https://solana-devnet.g.alchemy.com/v2/On35d8LdFc1QGYD-wCporecGj359qian")

	// calling
	TryGetBalance(cli, address)
}
