package address

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

// QueryBalance is for querying balance remaining in the input address
func QueryBalance(ctx context.Context, client *ethclient.Client, hexAddress string) {
	address := common.HexToAddress(hexAddress)
	balance, err := client.BalanceAt(ctx, address, nil)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("balance is : " + balance.String() + "\n")
	}
}
