package accounts

import (
	"encoding/hex"
	"fmt"
	"github.com/portto/solana-go-sdk/types"
	"log"
)

// TryCreateAccount is for creating new Solano account
func TryCreateAccount() {
	// create
	newAccount := types.NewAccount()
	// printing
	fmt.Printf("\nNew Account's PubKey: %v", newAccount.PublicKey.ToBase58())
	fmt.Printf("\nNew Account's PriKey: %v", hex.EncodeToString(newAccount.PrivateKey))
}

// TryRecoverAccount is for recover account info by input priKey in base 58
func TryRecoverAccount(base58priKey string) (types.Account, error) {
	// recover by priKey
	account, err := types.AccountFromBase58(base58priKey)
	if err != nil {
		log.Fatalf("Error, %v/n", err)
	}
	fmt.Printf("\nRecover PubKey: %v\n", account.PublicKey.ToBase58())
	return account, err
}
