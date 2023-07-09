package cosmos

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"time"
)

var (
	// RpcUrl ->
	RpcUrl = "https://rpc.osmotest5.osmosis.zone"
	// GRpcUrl ->
	//GRpcUrl = "fx-grpc.functionx.io:9090"
	GRpcUrl = "testnet-fx-grpc.functionx.io:9090"
	//GRpcUrl = "grpc.osmotest5.osmosis.zone:443"
	pool, _ = x509.SystemCertPool()
	config  = &tls.Config{
		RootCAs: pool,
	}
	Conn, _ = grpc.Dial(GRpcUrl,
		grpc.WithTransportCredentials(credentials.NewTLS(config)),
		grpc.WithUnaryInterceptor(grpc_middleware.ChainUnaryClient(GrpcLoggingInterceptor)), // logging interceptor
	)
)

// GrpcLoggingInterceptor is for GRPC logging
func GrpcLoggingInterceptor(ctx context.Context, method string, req, reply interface{},
	cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {

	startTime := time.Now()
	log.Printf("GRPC method: %s", method)
	err := invoker(ctx, method, req, reply, cc, opts...)
	log.Printf("Time taken: %s", time.Since(startTime))
	return err
}
