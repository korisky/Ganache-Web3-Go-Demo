package event

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

// https://dave-appleton.medium.com/overcoming-ethclients-filter-restrictions-81e232a8eccd
// https://ethereum.stackexchange.com/questions/136144/how-to-read-erc1155s-safetransferfrom-transactions-in-golang
// https://ethereum.stackexchange.com/questions/136144/how-to-read-erc1155s-safetransferfrom-transactions-in-golang
type LogTransferErc1155Single struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
}

// SubscribeErc1155SingleTransferEvent help subscribe transfer event
func SubscribeErc1155SingleTransferEvent(ctx context.Context, client *ethclient.Client) {

	// construct the event signature (quit different from Java)
	eventName := "TransferSingle"
	logTransferSingleSig := []byte(eventName + "{address,address,address,uint256,uint256}")
	logTransferSingleSigHash := crypto.Keccak256(logTransferSingleSig)

	fmt.Printf("event signature:%s", common.BytesToHash(logTransferSingleSigHash))

	// add topics into query
	filterQuery := ethereum.FilterQuery{
		FromBlock: big.NewInt(7979219),
		ToBlock:   big.NewInt(7979219),
		Topics:    [][]common.Hash{{common.BytesToHash(logTransferSingleSigHash)}},
	}

	// calling
	logs, err := client.FilterLogs(ctx, filterQuery)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(logs)

}
