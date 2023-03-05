package main

import (
	"context"
	"fmt"
	"github.com/portto/solana-go-sdk/client"
	"log"
)

var SOLANA_ENDPOINT = "https://solana-devnet.g.alchemy.com/v2/On35d8LdFc1QGYD-wCporecGj359qian"

func main() {
	// custom connection would be:
	// cli := client.NewClient(rpc.DevnetRPCEndpoint, rpc.WithHTTPClient(customClietn))
	cli := client.NewClient(SOLANA_ENDPOINT)

	resp, err := cli.GetVersion(context.TODO())
	if err != nil {
		log.Fatalf("Failed to get version info, err: %v", err)
	}

	fmt.Printf("The version is: %v\n", resp.SolanaCore)
}
