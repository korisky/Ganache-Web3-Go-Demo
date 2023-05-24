package cosmos

import (
	"context"
	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	"github.com/cosmos/cosmos-sdk/codec"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"testing"
)

func Test_getLatestBlockHeight(t *testing.T) {

	//rpcUrl := "grpc.osmosis.zone:9090" // osmosis 获取成功
	rpcUrl := "grpc.cosmos.interbloc.org:443" // cosmos-hub
	//rpcUrl := "fx-grpc.functionx.io:9090" // fxcore 获取失败
	//rpcUrl := "192.168.20.222:26657" // fxcore本地 获取失败

	conn, err := grpc.Dial(rpcUrl,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.ForceCodec(codec.NewProtoCodec(nil).GRPCCodec())))
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}

	defer conn.Close()

	tmClient := tmservice.NewServiceClient(conn)
	request := tmservice.GetLatestBlockRequest{}
	res, err := tmClient.GetLatestBlock(context.Background(), &request)
	if err != nil {
		log.Fatalf("Failed to get the latest block: %v", err)
	}

	log.Println("Latest block height:", res.Block.Header.Height)
}
