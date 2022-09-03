package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

var (
	ctx = context.Background()
	url = "https://kovan.infura.io/v3/3f0482cf4c3545dbabaeab75f414e467"
	// connect to own web3 node url
	client, _ = ethclient.DialContext(ctx, url)
)

func main() {

	// get header number
	header, _ := client.HeaderByNumber(ctx, nil)
	fmt.Println(header.Number.String())
	fmt.Println(header.TxHash.String())

	address := common.HexToAddress("0xe10eE98bB84B2073B88353e3AB4433916205DF40")
	balance, err := client.BalanceAt(ctx, address, nil)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(balance.String())
	}

}
