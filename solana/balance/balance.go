package balance

import (
	"context"
	"fmt"
	"github.com/portto/solana-go-sdk/client"
	"log"
)

// TryGetBalance could retry updated SOL remaining balance of an account
func TryGetBalance(cli *client.Client, acc string) {
	// try request
	balance, err := cli.GetBalance(context.Background(), acc)
	if err != nil {
		log.Fatalf("Got error, %v\n", err)
	}
	// resp
	fmt.Printf("\nAddress account balance would be: %v\n", balance)
}
