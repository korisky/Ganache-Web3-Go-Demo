package grpc

import (
	"context"
	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	"log"
	"own.cosmos.demo"
	"testing"
)

// Test_getLatestBlock
func Test_getLatestBlock(t *testing.T) {
	defer cosmos.Conn.Close()

	tmClient := tmservice.NewServiceClient(cosmos.Conn)
	request := tmservice.GetLatestBlockRequest{}
	res, _ := tmClient.GetLatestBlock(context.Background(), &request)
	log.Println("Latest block height:", res.Block.Header.Height)
}
