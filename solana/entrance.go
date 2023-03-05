package main

import (
	"fmt"
	"github.com/portto/solana-go-sdk/client"
	"github.com/portto/solana-go-sdk/rpc"
	"net/http"
	"strconv"
	"web3Demo/solana/balance"
	"web3Demo/solana/httpProxy"
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

// main are following stuff from https://yihau.gitbook.io/solana-go/tour/create-token-account/associated-token-account
func main() {
	tokenAddress := balance.TryFindAssociatedTokenAddress(
		address,
		"Gh9ZwEmdLJ8DscKNTkTqPbNwLNNBjuSzaG9Vp2KGtKJr")

	tokenBalance, u := balance.TryGetTokenBalance(cli, tokenAddress)
	tokenBalanceStr := strconv.FormatUint(tokenBalance, 10)
	fmt.Println("Token Balance in Raw: ", tokenBalanceStr)
	fmt.Println("Token Decimals: ", int(u))
}
