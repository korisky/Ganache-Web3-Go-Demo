package transfer

import (
	"context"
	"fmt"
	"github.com/portto/solana-go-sdk/client"
	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/program/sysprog"
	"github.com/portto/solana-go-sdk/program/token"
	"github.com/portto/solana-go-sdk/types"
	"log"
	"web3Demo/portto/accounts"
)

// TryTransferSol is for pure SOL transfer
func TryTransferSol(cli *client.Client, base58priKey, toBase58PubKey string) string {
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
				To:     common.PublicKeyFromString(toBase58PubKey),
				Amount: 3500,
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
	return txHash
}

// TryTransferToken is for transfer the token with instruction in msg
func TryTransferToken(cli *client.Client, accountPriKey, accountAssTokenPubKey, assTokenMintPubKey string) {
	account, _ := accounts.TryRecoverAccount(accountPriKey)
	assTokenAccount := common.PublicKeyFromString(accountAssTokenPubKey)

	// 1. fetch recent block-hash
	resp, _ := cli.GetLatestBlockhash(context.Background())
	// 2. create msg
	msg := types.NewMessage(types.NewMessageParam{
		RecentBlockhash: resp.Blockhash,
		FeePayer:        account.PublicKey,
		Instructions: []types.Instruction{
			token.TransferChecked(
				token.TransferCheckedParam{
					From:     assTokenAccount,
					To:       assTokenAccount,
					Mint:     common.PublicKeyFromString(assTokenMintPubKey),
					Auth:     account.PublicKey,
					Signers:  []common.PublicKey{},
					Amount:   3e5,
					Decimals: 6,
				}),
		}})
	// 3. create tx by msg + signer
	tx, err := types.NewTransaction(
		types.NewTransactionParam{
			Message: msg,
			Signers: []types.Account{account},
		})
	if err != nil {
		log.Fatalf("\nFailed to new tx, err:%v", err)
	}

	txHash, err := cli.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Fatalf("\nFailed to send tx, err:%v", err)
	}

	log.Println("txHash ", txHash)
}
