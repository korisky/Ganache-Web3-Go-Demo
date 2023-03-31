package main

import (
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
