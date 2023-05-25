package cosmos

import (
	"crypto/tls"
	"crypto/x509"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	// RpcUrl ->
	RpcUrl = "https://rpc.osmotest5.osmosis.zone"
	// GRpcUrl ->
	GRpcUrl = "fx-grpc.functionx.io:9090"
	pool, _ = x509.SystemCertPool()
	config  = &tls.Config{
		RootCAs: pool,
	}
	Conn, _ = grpc.Dial(RpcUrl, grpc.WithTransportCredentials(credentials.NewTLS(config)))
)
