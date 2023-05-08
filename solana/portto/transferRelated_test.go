package main

import (
	"context"
	"fmt"
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

	//transaction, _ := cli.GetTransaction(context.Background(), "41ZCAJiXCHmXjtGTm5VtB5q1F1eWoGa84xBBiErdDsqf7K3m2yn4GACjuXzimMg6zjtGA2MGuu9gPRmRqA6bV3ka")
	transaction, _ := cli.GetTransaction(context.Background(), "2BghzZxWRXR7zy9by3GHGhi32yYbrwY59ivvdNzfMCcuHJ3FyFVkLy7bxkxKqcdrKVpaszvdSoCrjPSyAgMQiTts")

	for _, instruction := range transaction.Transaction.Message.Instructions {
		amount, err := transfer.DecodeNativeTransferAmount(instruction.Data)
		if nil == err {
			fmt.Printf("Decode transaction with amount: %v \n", float64(amount)/math.Pow10(9))
		}
	}
}
