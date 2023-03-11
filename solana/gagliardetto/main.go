package main

import (
	"context"
	"fmt"
	"web3Demo/solana/httpProxy"
)

func main() {
	client := httpProxy.NewRPC("https://solana-devnet.g.alchemy.com/v2/On35d8LdFc1QGYD-wCporecGj359qian")
	resp, _ := client.GetVersion(context.TODO())
	fmt.Println("Solana-Core version: " + resp.SolanaCore)
}
