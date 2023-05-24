package cosmos

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"testing"
)

// Test_getLatestBlockHeight now got 403 forbidden
func Test_getLatestBlockHeight(t *testing.T) {

	rpcUrl := "fx-grpc.functionx.io:9090"
	//rpcUrl := "fx-grpc.functionx.io:443"
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
