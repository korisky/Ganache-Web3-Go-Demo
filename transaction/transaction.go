package transaction

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
)

func QueryTransactions(ctx context.Context, client *ethclient.Client) {
	block, _ := client.BlockByNumber(ctx, nil)
	for _, txn := range block.Transactions() {
		fmt.Println("Single Txn")
		fmt.Println(txn.Value().String())
		fmt.Println(txn.Gas())
		fmt.Println(txn.GasPrice().Uint64())
		fmt.Println(txn.Nonce())
		fmt.Println(txn.To().Hex())
		fmt.Println()
		break
	}
}
