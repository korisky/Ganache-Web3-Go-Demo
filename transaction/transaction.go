package transaction

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	ctx = context.Background()
	url = "https://goerli.infura.io/v3/3f0482cf4c3545dbabaeab75f414e467"
	// connect to own web3 node url
	client, _ = ethclient.DialContext(ctx, url)
)

func QueryTransactions() {
	block, _ := client.BlockByNumber(ctx, nil)
	for _, txn := range block.Transactions() {
		fmt.Println("Single Txn")
		fmt.Println(txn.Value().String())
		fmt.Println(txn.Gas())
		fmt.Println(txn.GasPrice().Uint64())
		fmt.Println(txn.Nonce())
		fmt.Println(txn.To().Hex())
		fmt.Println()
	}
}
