package event

import (
	"context"
	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"testing"
)

func Test_EthGetLog(t *testing.T) {

	//url := "https://chain-gateway.functionx.io/v1/polygon-mainnet/"
	url := "https://polygon-mumbai.g.alchemy.com/v2/"
	client, _ := ethclient.DialContext(context.Background(), url)

	logTransferSingleSig := []byte("Transfer(address,adadress,uint256)")

	filterQuery := ethereum.FilterQuery{
		FromBlock: big.NewInt(45056699),
		ToBlock:   big.NewInt(45056699 + 50),
		Topics:    [][]common.Hash{{common.BytesToHash(crypto.Keccak256(logTransferSingleSig))}},
	}

	// calling
	logs, err := client.FilterLogs(context.Background(), filterQuery)
	if err != nil {
		log.Fatal(err)
	}

	spew.Dump(logs)
}
