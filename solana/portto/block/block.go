package block

import (
	"context"
	"github.com/davecgh/go-spew/spew"
	"github.com/portto/solana-go-sdk/client"
)

func TryGetBlockByBlockHeight(cli *client.Client, height uint64) client.GetBlockResponse {
	block, _ := cli.GetBlock(context.Background(), height)
	for _, blockTxn := range block.Transactions {
		for _, sig := range blockTxn.Transaction.Signatures {
			spew.Dump(sig)
		}
	}
	return block
}
