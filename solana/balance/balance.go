package balance

import (
	"context"
	"fmt"
	"github.com/portto/solana-go-sdk/client"
	"github.com/portto/solana-go-sdk/common"
	"log"
	"strconv"
)

// TryGetBalance could retry updated SOL remaining balance of an account
func TryGetBalance(cli *client.Client, pub string) {
	// req
	balance, err := cli.GetBalance(context.Background(), pub)
	if err != nil {
		log.Fatalf("Got error, %v\n", err)
	}
	// resp
	fmt.Printf("Address account balance would be: %v\n", balance)
}

func TryGetTokenBalance(cli *client.Client, tokenPub string) (uint64, uint8) {
	balance, u, err := cli.GetTokenAccountBalance(context.Background(), tokenPub)
	if err != nil {
		log.Fatalf("Got error, %v\n", err)
	}
	return balance, u
}

// TryRequestAirdrop is for request airdrop
func TryRequestAirdrop(cli *client.Client, pub string) {
	// req
	txHash, err := cli.RequestAirdrop(context.Background(), pub, 1e9)
	if err != nil {
		log.Fatalf("Got error, %v\n", err)
	}
	// resp
	fmt.Printf("Got txHash: %v for airdrop\n", txHash)
}

// TryFindAssociatedTokenAddress is for find all tokenAccount for an AddressAccount
func TryFindAssociatedTokenAddress(account string, tokenAddress string) string {
	// req
	ata, _, err := common.FindAssociatedTokenAddress(
		common.PublicKeyFromString(account),
		common.PublicKeyFromString(tokenAddress))
	if err != nil {
		log.Fatalf("find ata error, err: %v", err)
	}
	fmt.Println("Account Token Address: ", ata.ToBase58())
	return ata.ToBase58()
}

// TryGetAssociatedTokenAddressBalance try find associated token address & get the balance
func TryGetAssociatedTokenAddressBalance(cli *client.Client, base58pubKey string, tokenAddress string) {
	// 1. find associated-token-address by using pubKey & token-address
	ata := TryFindAssociatedTokenAddress(base58pubKey, tokenAddress)
	// 2. try to get token balance
	tokenBalance, u := TryGetTokenBalance(cli, ata)
	tokenBalanceStr := strconv.FormatUint(tokenBalance, 10)
	fmt.Println("Token Balance in Raw: ", tokenBalanceStr)
	fmt.Println("Token Decimals: ", int(u))
}