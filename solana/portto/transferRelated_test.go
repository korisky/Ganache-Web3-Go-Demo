package main

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/portto/solana-go-sdk/client"
	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/types"
	"strconv"
	"strings"
	"testing"
	"web3Demo/portto/transfer"
)

// Test_TransferSol
func Test_TransferSol(t *testing.T) {
	fromPriKey := ""
	toPubKey := "GQ6V9ZLVibN7eAtxEQxLJjXX8L9RybMJPpUCwi16vVgL"
	for i := 0; i < 5; i++ {
		transfer.TryTransferSol(cli, fromPriKey, toPubKey)
	}
}

func decodeTxnInput(transaction *client.GetTransactionResponse) (uint64, uint8) {

	message := transaction.Transaction.Message
	meta := transaction.Meta
	accounts := message.Accounts

	// 1. for SOL transfer checking
	if len(message.Instructions) == 1 {
		data := message.Instructions[0].Data
		return binary.LittleEndian.Uint64(data[4:]), 9
	}

	// 2. for SPL simple transfer checking
	// filter for all tokenProgram related instructions
	tokenProgramInstructions := make([]types.CompiledInstruction, len(message.Instructions))
	for _, instruction := range message.Instructions {

		if strings.EqualFold(accounts[instruction.ProgramIDIndex].ToBase58(), common.TokenProgramID.ToBase58()) {
			tokenProgramInstructions = append(tokenProgramInstructions, instruction)
		}

		if bytes.Equal(accounts[instruction.ProgramIDIndex].Bytes(), common.MetaplexTokenMetaProgramID.Bytes()) {
			// no-fungible NFT token mint checking
			if len(meta.PreTokenBalances) == 0 && len(meta.PostTokenBalances) == 1 {
				balance := meta.PostTokenBalances[0]
				amt, _ := strconv.ParseUint(balance.UITokenAmount.Amount, 10, 64)
				return amt, balance.UITokenAmount.Decimals
			}
		}
	}
	// only one token program instruction -> must be normal SPL transfer
	if len(tokenProgramInstructions) == 1 {
		data := tokenProgramInstructions[0].Data
		return binary.LittleEndian.Uint64(data[1:]), meta.PostTokenBalances[0].UITokenAmount.Decimals
	}

	return 0, 0
}

// Test_DecodingAmount
func Test_DecodingAmount(t *testing.T) {

	//transaction, _ := cli.GetTransaction(context.Background(), "41ZCAJiXCHmXjtGTm5VtB5q1F1eWoGa84xBBiErdDsqf7K3m2yn4GACjuXzimMg6zjtGA2MGuu9gPRmRqA6bV3ka") // sol-self
	transaction, _ := cli.GetTransaction(context.Background(), "4JQVvVJ3QQaBosMDUB7S9D3xRujgcb47jpdyXiUt9us5j5YjWTg9x6sHhpaPagsETZ5hPMVJR3LQ1SBnp4YCyHti") // spl-self
	//transaction, _ := cli.GetTransaction(context.Background(), "5Q5upfHBHLNH7ScCwLfEFc8iXD2RKNJYmK2f8hg7eqxGBjtGLcgxQS3GrJhH8eYc3kuXY9X5VQEA9cvPxWs1RzZ4") // spl-dif
	//transaction, _ := cli.GetTransaction(context.Background(), "5T2KkZ9fH9J1qbeTS1MhXVYjDfXoGVrAJrYrVEPmrLoSiUzn6xqGS264tphX6xCnN8zWF4ZrRkBphGfimExpEPQd") // spl-nft
	spew.Dump(transaction)

	amt, decimal := decodeTxnInput(transaction)
	fmt.Printf("\n\n Got amt:%v, decimal:%v\n\n", amt, decimal)

	//
	//for _, instruction := range transaction.Transaction.Message.Instructions {
	//	amount, err := transfer.DecodeNativeTransferAmount(instruction)
	//	if nil == err {
	//
	//		var realAmount float64
	//		if instruction.ProgramIDIndex == 1 {
	//			// for SOL transfer
	//			realAmount = float64(amount) / math.Pow10(9)
	//
	//		} else if instruction.ProgramIDIndex == 4 {
	//
	//			// for SPL transfer
	//			preTokenBalances := transaction.Meta.PreTokenBalances
	//			postTokenBalances := transaction.Meta.PostTokenBalances
	//
	//			var decimal uint8
	//			for _, balance := range preTokenBalances {
	//				if balance.Owner == "AnayTW335MabjhtXTJeBit5jdLhNeUVBVPXeRKCid79D" {
	//					decimal = balance.UITokenAmount.Decimals
	//					break
	//				}
	//			}
	//			if 0 == decimal {
	//				for _, balance := range postTokenBalances {
	//					if balance.Owner == "AnayTW335MabjhtXTJeBit5jdLhNeUVBVPXeRKCid79D" {
	//						decimal = balance.UITokenAmount.Decimals
	//						break
	//					}
	//				}
	//			}
	//
	//			if 0 != decimal {
	//				realAmount = float64(amount) / math.Pow10(int(decimal))
	//			}
	//		}
	//
	//		fmt.Printf("Decode transaction with amount: %v \n", realAmount)
	//	}
	//}
}
