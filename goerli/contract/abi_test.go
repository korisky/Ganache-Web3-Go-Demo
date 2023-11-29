package contract

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
	"log"
	"math/big"
	"strings"
	"testing"
)

const nodeUrl = ""
const contract = "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"

func Test_ContractABICalling(t *testing.T) {

	// node requester init
	client, err := ethclient.Dial(nodeUrl)
	if err != nil {
		log.Fatalln(err)
	}

	// contract & ABI preparation
	contractAddr := common.HexToAddress(contract)
	targetABI := `[{"constant":true,"inputs":[],"name":"totalSupply","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"}]`
	parsedABI, err := abi.JSON(strings.NewReader(targetABI))
	if err != nil {
		log.Fatalln(err)
	}
	data, err := parsedABI.Pack("totalSupply")
	if err != nil {
		log.Fatalln(err)
	}

	// execute chain query
	msg := ethereum.CallMsg{
		To:   &contractAddr,
		Data: data,
	}
	result, err := client.CallContract(context.Background(), msg, nil)
	if err != nil {
		log.Fatalln(err)
	}

	// decode result
	abiOutput, err := parsedABI.Unpack("totalSupply", result)
	if err != nil {
		log.Fatalln(err)
	}
	// here use NewFromBigInt(xxx, -18), means divide 10^18, and remain accuracy
	rawTotalSupply := abiOutput[0].(*big.Int)
	totalSupply := decimal.NewFromBigInt(rawTotalSupply, -18)
	fmt.Printf("Concrete total supply: %s", totalSupply.String())
}
