package cosmos

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"testing"
)

// Test_getLatestBlockHeight now got 403 forbidden
func Test_getLatestBlockHeight(t *testing.T) {

	rpcUrl := "fx-grpc.functionx.io:9090"
	//rpcUrl := "192.168.20.222:26657"

	//caCert, err := os.ReadFile("/Users/pundix2022/ssl/cert.pem")
	//if err != nil {
	//	fmt.Println("Error loading certificate:", err)
	//	return
	//}
	//caCertPool := x509.NewCertPool()
	//caCertPool.AppendCertsFromPEM(caCert)
	//
	//config := &tls.Config{
	//	RootCAs:            caCertPool,
	//	InsecureSkipVerify: true,
	//}
	//conn, err := grpc.Dial(rpcUrl, grpc.WithTransportCredentials(credentials.NewTLS(config)))

	pool, _ := x509.SystemCertPool()
	config := &tls.Config{
		RootCAs: pool,
		//InsecureSkipVerify: true,
	}
	creds := credentials.NewTLS(config)
	conn, err := grpc.Dial(rpcUrl, grpc.WithTransportCredentials(creds))

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

// Test_Another is another test
func Test_Another(t *testing.T) {
	nodeURI := "https://fx-grpc.functionx.io:9090"
	client, err := client.NewClientFromNode(nodeURI)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Get the latest block
	var block_height int64
	block_height = 2341234234
	block, err := client.Block(context.Background(), &block_height)
	if err != nil {
		log.Fatalf("Failed to get block: %v", err)
	}

	// Get the block height
	height := block.Block.Height
	fmt.Printf("The latest block height is %d\n", height)
}
