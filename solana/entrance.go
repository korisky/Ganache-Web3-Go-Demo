package main

import (
	"fmt"
	"github.com/portto/solana-go-sdk/client"
	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/rpc"
	"net/http"
	"web3Demo/solana/httpProxy"
)

var (
	// own address account
	address      = "AnayTW335MabjhtXTJeBit5jdLhNeUVBVPXeRKCid79D"
	tokenAddress = "Gh9ZwEmdLJ8DscKNTkTqPbNwLNNBjuSzaG9Vp2KGtKJr"
	ownEndpoint  = "https://solana-devnet.g.alchemy.com/v2/On35d8LdFc1QGYD-wCporecGj359qian"
	// custom connection would be:
	// cli := client.New(rpc.withEndpoint(rpc.DevnetRPCEndpoint), rpc.WithHTTPClient(customClietn))
	httpClient = &http.Client{
		Transport: httpProxy.LoggingRoundTripper{Proxied: http.DefaultTransport},
	}
	cli = client.New(rpc.WithEndpoint(ownEndpoint), rpc.WithHTTPClient(httpClient))
)

// main are following stuff from https://yihau.gitbook.io/solana-go/tour/create-token-account/associated-token-account
func main() {
	//balance.TryGetAssociatedTokenAddressBalance(cli, address, tokenAddress)
	associatedTokenAddress, _, _ := OwnFindAssociatedTokenAddress(address, tokenAddress)
	fmt.Println("associated token address ", associatedTokenAddress)
}

// OwnFindAssociatedTokenAddress for Java migration use, but found out Java not support UnsignedBytes
func OwnFindAssociatedTokenAddress(accountAddress, tokenAddress string) (common.PublicKey, uint8, error) {
	walletAddress := common.PublicKeyFromString(accountAddress)
	tokenMintAddress := common.PublicKeyFromString(tokenAddress)
	fmt.Println("walletAddress bytes: ", walletAddress.Bytes())
	fmt.Println(len(walletAddress.Bytes()))
	fmt.Println("tokenProgram bytes: ", common.TokenProgramID.Bytes())
	fmt.Println(len(common.TokenProgramID.Bytes()))
	fmt.Println("tokenMintAddress bytes: ", tokenMintAddress.Bytes())
	fmt.Println(len(tokenMintAddress.Bytes()))

	seeds := [][]byte{}
	seeds = append(seeds, walletAddress.Bytes())
	seeds = append(seeds, common.TokenProgramID.Bytes())
	seeds = append(seeds, tokenMintAddress.Bytes())

	fmt.Println("whole seeds bytes: ", seeds)
	fmt.Println(len(seeds))

	return common.FindProgramAddress(seeds, common.SPLAssociatedTokenAccountProgramID)
}
