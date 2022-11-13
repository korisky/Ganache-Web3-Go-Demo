package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"quickNode/goerli/account"
	"quickNode/goerli/transaction"
)

var (
	ctx = context.Background()
	url = "https://goerli.infura.io/v3/3f0482cf4c3545dbabaeab75f414e467"
	// connect to own web3 node url
	client, _ = ethclient.DialContext(ctx, url)
)

// from https://goethereumbook.org/client/
func main() {

	// get header number
	header, _ := client.HeaderByNumber(ctx, nil)
	fmt.Println("Latest Block Number:" + header.Number.String() + "\n")
	fmt.Println("Latest Txn Hash:" + header.TxHash.String() + "\n")

	// call query balance
	account.QueryBalance(ctx, client, "0xe10eE98bB84B2073B88353e3AB4433916205DF40", nil)

	// call transaction
	transaction.QueryTransactions(ctx, client)

	// judge contract or address
	account.JudgeAddressOrSmartContract(ctx, client, "0xe10eE98bB84B2073B88353e3AB4433916205DF40")
	account.JudgeAddressOrSmartContract(ctx, client, "0xff02b7d59975E76F67B63b20b813a9Ec0f6AbD60")
}
