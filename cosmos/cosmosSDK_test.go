package cosmos

import (
	"context"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	"google.golang.org/grpc"
	"io"
	"log"
	"net/http"
	"time"

	"testing"
)

func Test_getStatusRpc(t *testing.T) {

	// Osmosis Testnet Rpc
	rpcUrl := "https://rpc.osmotest5.osmosis.zone"

	resp, err := http.Get(rpcUrl + "/status")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(body)
}

func Test_getStatusGRpc(t *testing.T) {

	//rpcUrl := "https://rpc.osmosis.zone:443"
	rpcUrl := "grpc.osmosis.zone:9090"

	conn, err := grpc.Dial(rpcUrl, grpc.WithInsecure(), grpc.WithTimeout(1*time.Second))
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
