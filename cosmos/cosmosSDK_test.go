package cosmos

import (
	"context"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"net/http"
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

	rpcUrl := "grpc.osmosis.zone:9090"
	//rpcUrl := "fx-grpc.functionx.io:9090"

	//config := &tls.Config{
	//	InsecureSkipVerify: false,
	//}

	//conn, err := grpc.Dial(rpcUrl, grpc.WithTransportCredentials(credentials.NewTLS(config)))
	conn, err := grpc.Dial(rpcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
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
