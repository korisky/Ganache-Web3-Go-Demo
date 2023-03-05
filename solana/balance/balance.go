package balance

import (
	"context"
	"fmt"
	"github.com/portto/solana-go-sdk/client"
	"log"
)

// TryGetBalance could retry updated SOL remaining balance of an account
func TryGetBalance(cli *client.Client, pub string) {
	// req
	balance, err := cli.GetBalance(context.Background(), pub)
	if err != nil {
		log.Fatalf("Got error, %v\n", err)
	}
	// resp
	fmt.Printf("\nAddress account balance would be: %v\n", balance)
}

// TryRequestAirdrop is for request airdrop
func TryRequestAirdrop(cli *client.Client, pub string) {
	// req
	txHash, err := cli.RequestAirdrop(context.Background(), pub, 1e9)
	if err != nil {
		log.Fatalf("Got error, %v\n", err)
	}
	// resp
	fmt.Printf("Got txHash: %v for airdrop", txHash)
}
