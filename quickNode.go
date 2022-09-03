package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
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

}
