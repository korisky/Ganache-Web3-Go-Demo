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

func Test_getNodeInfo(t *testing.T) {
	defer cosmos.Conn.Close()

	client := tmservice.NewServiceClient(cosmos.Conn)
	request := tmservice.GetNodeInfoRequest{}
	info, _ := client.GetNodeInfo(context.Background(), &request)
	log.Printf("Node info: Network: %v, ListenAddr: %v", info.DefaultNodeInfo.Network, info.DefaultNodeInfo.ListenAddr)
	log.Printf("Application info: goVer: %v, appName: %v, cosmosSdkVer: %v",
		info.ApplicationVersion.GoVersion, info.ApplicationVersion.AppName, info.ApplicationVersion.CosmosSdkVersion)
}
