package accounts

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/portto/solana-go-sdk/client"
	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/types"
	"log"
)

// TryCreateAccount is for creating new Solano account
func TryCreateAccount() {
	// create
	newAccount := types.NewAccount()
	// printing
	fmt.Printf("New Account's PubKey: %v\n", newAccount.PublicKey.ToBase58())
	fmt.Printf("New Account's PriKey: %v\n", hex.EncodeToString(newAccount.PrivateKey))
}

// TryRecoverAccount is for recover account info by input priKey in base 58
func TryRecoverAccount(base58priKey string) (types.Account, error) {
	// recover by priKey
	account, err := types.AccountFromBase58(base58priKey)
	if err != nil {
		log.Fatalf("Error, %v/n", err)
	}
	fmt.Printf("Recover PubKey: %v\n", account.PublicKey.ToBase58())
	return account, err
}

func TryGetTokenAccountsByOwner(cli *client.Client, base58Address string) {
	resp, _ := cli.GetTokenAccountsByOwner(context.Background(), base58Address)
	spew.Dump(resp)
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
