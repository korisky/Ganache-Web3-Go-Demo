package main

import (
	"context"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"math"
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

// Test_DecodingAmount
func Test_DecodingAmount(t *testing.T) {

	//transaction, _ := cli.GetTransaction(context.Background(), "41ZCAJiXCHmXjtGTm5VtB5q1F1eWoGa84xBBiErdDsqf7K3m2yn4GACjuXzimMg6zjtGA2MGuu9gPRmRqA6bV3ka") // sol-self
	//transaction, _ := cli.GetTransaction(context.Background(), "4JQVvVJ3QQaBosMDUB7S9D3xRujgcb47jpdyXiUt9us5j5YjWTg9x6sHhpaPagsETZ5hPMVJR3LQ1SBnp4YCyHti") // spl-self
	//transaction, _ := cli.GetTransaction(context.Background(), "5Q5upfHBHLNH7ScCwLfEFc8iXD2RKNJYmK2f8hg7eqxGBjtGLcgxQS3GrJhH8eYc3kuXY9X5VQEA9cvPxWs1RzZ4") // spl-dif
	transaction, _ := cli.GetTransaction(context.Background(), "5T2KkZ9fH9J1qbeTS1MhXVYjDfXoGVrAJrYrVEPmrLoSiUzn6xqGS264tphX6xCnN8zWF4ZrRkBphGfimExpEPQd") // spl-nft
	spew.Dump(transaction)

	for _, instruction := range transaction.Transaction.Message.Instructions {
		amount, err := transfer.DecodeNativeTransferAmount(instruction)
		if nil == err {

			var realAmount float64
			if instruction.ProgramIDIndex == 1 {
				// for SOL transfer
				realAmount = float64(amount) / math.Pow10(9)

			} else if instruction.ProgramIDIndex == 4 {

				// for SPL transfer
				preTokenBalances := transaction.Meta.PreTokenBalances
				postTokenBalances := transaction.Meta.PostTokenBalances

				var decimal uint8
				for _, balance := range preTokenBalances {
					if balance.Owner == "AnayTW335MabjhtXTJeBit5jdLhNeUVBVPXeRKCid79D" {
						decimal = balance.UITokenAmount.Decimals
						break
					}
				}
				if 0 == decimal {
					for _, balance := range postTokenBalances {
						if balance.Owner == "AnayTW335MabjhtXTJeBit5jdLhNeUVBVPXeRKCid79D" {
							decimal = balance.UITokenAmount.Decimals
							break
						}
					}
				}

				if 0 != decimal {
					realAmount = float64(amount) / math.Pow10(int(decimal))
				}
			}

			fmt.Printf("Decode transaction with amount: %v \n", realAmount)
		}
	}
}
