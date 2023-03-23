package accounts

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/portto/solana-go-sdk/client"
	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/rpc"
	"github.com/portto/solana-go-sdk/types"
	"log"
	"sort"
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

// TryFindTxnSigByOwnerAddress is for query top limit latest txn by using owner address (might be slow if contains many address)
func TryFindTxnSigByOwnerAddress(cli *client.Client, base58Owner string, limit int) []client.GetTransactionResponse {
	// 1. find all related account
	associatedAccounts, _ := cli.GetTokenAccountsByOwner(context.Background(), base58Owner)

	// 2. find all related accounts signatures, desc
	uniSigSet := make(map[string]rpc.SignatureWithStatus)
	walletAccSigs, _ := cli.GetSignaturesForAddressWithConfig(context.Background(), base58Owner, rpc.GetSignaturesForAddressConfig{Limit: limit})
	for _, sig := range walletAccSigs {
		uniSigSet[sig.Signature] = sig
	}
	for assTokAcc := range associatedAccounts {
		assTokAccSigs, _ := cli.GetSignaturesForAddressWithConfig(context.Background(), assTokAcc.ToBase58(), rpc.GetSignaturesForAddressConfig{Limit: limit})
		for _, sig := range assTokAccSigs {
			uniSigSet[sig.Signature] = sig
		}
	}

	// 3. create slice for map & sort desc
	uniSigSlice := make([]rpc.SignatureWithStatus, 0, len(uniSigSet))
	for _, val := range uniSigSet {
		uniSigSlice = append(uniSigSlice, val)
	}
	sort.Slice(uniSigSlice, func(i, j int) bool {
		// reverse, bigger order first
		return *uniSigSlice[i].BlockTime > *uniSigSlice[j].BlockTime
	})

	// 4. get all txn full details
	txnSlice := make([]client.GetTransactionResponse, 0, limit)
	for _, sigStatus := range uniSigSlice[:limit] {
		transaction, _ := cli.GetTransaction(context.Background(), sigStatus.Signature)
		txnSlice = append(txnSlice, *transaction)
	}

	spew.Dump(txnSlice)
	return txnSlice
}
