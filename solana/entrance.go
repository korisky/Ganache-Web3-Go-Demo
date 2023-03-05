package main

import (
	"github.com/portto/solana-go-sdk/client"
	"github.com/portto/solana-go-sdk/rpc"
	"net/http"
	"quickNode/solana/httpProxy"
)

var (
	// own address account
	address     = "AnayTW335MabjhtXTJeBit5jdLhNeUVBVPXeRKCid79D"
	ownEndpoint = "https://solana-devnet.g.alchemy.com/v2/On35d8LdFc1QGYD-wCporecGj359qian"
	// custom connection would be:
	// cli := client.New(rpc.withEndpoint(rpc.DevnetRPCEndpoint), rpc.WithHTTPClient(customClietn))
	httpClient = &http.Client{
		Transport: httpProxy.LoggingRoundTripper{Proxied: http.DefaultTransport},
	}
	cli = client.New(rpc.WithEndpoint(ownEndpoint), rpc.WithHTTPClient(httpClient))
)

func main() {

}
