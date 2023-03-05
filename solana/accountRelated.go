package main

import (
	"context"
	"fmt"
	"github.com/portto/solana-go-sdk/client"
	"log"
)

func TryGetBalance(cli *client.Client, acc string) {
	// try request
	balance, err := cli.GetBalance(context.Background(), acc)
	if err != nil {
		log.Fatalf("Got error, %v\n", err)
	}
	// resp
	fmt.Printf("Address account balance would be: %v\n", balance)
}
