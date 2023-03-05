package transfer

import (
	"context"
	"fmt"
	"github.com/portto/solana-go-sdk/client"
	"github.com/portto/solana-go-sdk/program/sysprog"
	"github.com/portto/solana-go-sdk/types"
	"log"
	"quickNode/solana/accounts"
)

// TryTransferSimple is for pure SOL transfer
func TryTransferSimple(cli *client.Client, base58priKey string) {
	account, _ := accounts.TryRecoverAccount(base58priKey)
	// 1. fetch recent block-hash
	resp, err := cli.GetLatestBlockhash(context.Background())
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}
	// 2. create msg
	msg := types.NewMessage(types.NewMessageParam{
		FeePayer:        account.PublicKey,
		RecentBlockhash: resp.Blockhash,
		Instructions: []types.Instruction{
			sysprog.Transfer(sysprog.TransferParam{
				From:   account.PublicKey,
				To:     account.PublicKey,
				Amount: 5000,
			}),
		},
	})

	// 3. create tx by msg + signer
	tx, err := types.NewTransaction(types.NewTransactionParam{
		Message: msg,
		Signers: []types.Account{account},
	})
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}

	// 4. send txn
	txHash, err := cli.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}

	fmt.Printf("On-chain txHash: %v\n", txHash)
}
