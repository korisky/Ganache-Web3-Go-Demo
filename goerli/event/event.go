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
	logTransferSingleSig := []byte(eventName + "(address,address,address,uint256,uint256)")
	logTransferSingleSigHash := crypto.Keccak256(logTransferSingleSig)

	fmt.Printf("event signature:%s", common.BytesToHash(logTransferSingleSigHash))

	// add topics into query
	filterQuery := ethereum.FilterQuery{
		FromBlock: big.NewInt(7962780),
		ToBlock:   big.NewInt(7962780),
		Topics:    [][]common.Hash{{common.BytesToHash(logTransferSingleSigHash)}},
		//Topics: [][]common.Hash{{common.HexToHash("0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62")}},
	}

	// calling
	logs, err := client.FilterLogs(ctx, filterQuery)
	if err != nil {
		log.Fatal(err)
	}

	for _, vlogs := range logs {
		fmt.Printf("Log Block Number: %d\n", vlogs.BlockNumber)
		fmt.Printf("Log Index: %d\n", vlogs.Index)
		fmt.Println()

		sig := vlogs.Topics[0]
		fmt.Printf("Event Signature: %s\n", sig)
		fmt.Printf("Contract Address: %s\n", vlogs.Address.Hex())

		var logTransferErc1155Single LogTransferErc1155Single
		logTransferErc1155Single.From = common.HexToAddress(vlogs.Topics[2].Hex())
		logTransferErc1155Single.To = common.HexToAddress(vlogs.Topics[3].Hex())
	}
}
