package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"quickNode/address"
	"quickNode/transaction"
)

var (
	ctx = context.Background()
	url = "https://goerli.infura.io/v3/3f0482cf4c3545dbabaeab75f414e467"
	// connect to own web3 node url
	client, _ = ethclient.DialContext(ctx, url)
)

func main() {

	// get header number
	header, _ := client.HeaderByNumber(ctx, nil)
	fmt.Println("Latest Block Number:" + header.Number.String() + "\n")
	fmt.Println("Latest Txn Hash:" + header.TxHash.String() + "\n")

	// call query balance
	address.QueryBalance(ctx, client, "0xe10eE98bB84B2073B88353e3AB4433916205DF40")

	// call transaction
	transaction.QueryTransactions(ctx, client)
}
