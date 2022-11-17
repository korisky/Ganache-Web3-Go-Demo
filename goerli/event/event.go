package event

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

type LogTransferErc1155Single struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
}

// SubscribeErc1155SingleTransferEvent help subscribe transfer event
func SubscribeErc1155SingleTransferEvent(ctx context.Context, client *ethclient.Client) {
	filterQuery := ethereum.FilterQuery{
		Topics: nil,
	}
}
