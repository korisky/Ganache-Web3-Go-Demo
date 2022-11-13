package account

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math"
	"math/big"
)

// QueryBalance is for querying balance remaining in the input account
func QueryBalance(ctx context.Context, client *ethclient.Client, hexAddress string, blockNum *big.Int) {
	address := common.HexToAddress(hexAddress)
	balance, err := client.BalanceAt(ctx, address, blockNum)
	if err != nil {
		log.Fatal(err)
	} else {
		// from wei to Eth
		fBalance := new(big.Float)
		fBalance.SetString(balance.String())
		ethValue := new(big.Float).Quo(fBalance, big.NewFloat(math.Pow10(18)))
		fmt.Println("balance in Eth would be : " + ethValue.String() + "\n")
	}
}

func JudgeAddressOrSmartContract(ctx context.Context, client *ethclient.Client, hexAddress string) {
	address := common.HexToAddress(hexAddress)
	bytes, err := client.CodeAt(context.Background(), address, nil) // nil is the latest block
	if err != nil {
		log.Fatal(err)
	}

	isContract := len(bytes) > 0
	fmt.Printf("is contract: %v\n", isContract)
}
