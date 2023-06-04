package grpc

import (
	"context"
	"crypto/sha256"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	"github.com/cosmos/cosmos-sdk/types/tx"
	cosmos_proto "github.com/cosmos/gogoproto/proto"
	"log"
	"own.cosmos.demo"
	"testing"
)

func Test_getNodeInfo(t *testing.T) {
	defer cosmos.Conn.Close()

	client := tmservice.NewServiceClient(cosmos.Conn)
	request := tmservice.GetNodeInfoRequest{}
	info, _ := client.GetNodeInfo(context.Background(), &request)
	log.Printf("Node info: Network: %v, ListenAddr: %v", info.DefaultNodeInfo.Network, info.DefaultNodeInfo.ListenAddr)
	log.Printf("Application info: goVer: %v, appName: %v, cosmosSdkVer: %v",
		info.ApplicationVersion.GoVersion, info.ApplicationVersion.AppName, info.ApplicationVersion.CosmosSdkVersion)
}

// Test_getLatestBlock
func Test_getLatestBlock(t *testing.T) {
	defer cosmos.Conn.Close()

	tmClient := tmservice.NewServiceClient(cosmos.Conn)
	request := tmservice.GetLatestBlockRequest{}
	res, _ := tmClient.GetLatestBlock(context.Background(), &request)
	log.Println("Latest block height:", res.Block.Header.Height)
}

// Test_decodeBlock
func Test_decodeBlock(t *testing.T) {
	defer cosmos.Conn.Close()

	//registry := types.NewInterfaceRegistry()
	//std.RegisterInterfaces(registry)
	//banktypes.RegisterInterfaces(registry)
	//pc := codec.NewProtoCodec(registry)

	tmClient := tmservice.NewServiceClient(cosmos.Conn)
	request := tmservice.GetBlockByHeightRequest{Height: int64(8592209)}
	res, err := tmClient.GetBlockByHeight(context.Background(), &request)
	if err != nil {
		log.Fatalf("Failed to get block by height: %v", err)
	}

	blockHash := sha256.Sum256(res.BlockId.Hash)
	fmt.Printf("Block hash: %x\n", blockHash)

	// traverse & decode transaction
	for _, txBytes := range res.Block.Data.Txs {
		var txObj tx.Tx
		err := cosmos_proto.Unmarshal(txBytes, &txObj)
		if err != nil {
			log.Fatalf("Failed to unmarshal transaction: %v", err)
		}

		// construct txHash
		hash := sha256.Sum256(txBytes)
		fmt.Printf("Transaction hash: %x\n", hash)

	}
}
