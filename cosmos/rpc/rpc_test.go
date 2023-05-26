package rpc

import (
	"context"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/tendermint/tendermint/rpc/client/http"
	cosmos "own.cosmos.demo"
	"testing"
)

func Test_getLatestBlockHeight_Rpc(t *testing.T) {
	// request
	c, _ := http.New(cosmos.RpcUrl)
	// decoding
	resp, _ := c.Status(context.Background())
	fmt.Println(resp.SyncInfo.LatestBlockHeight)
}

func Test_getLatestBlock_Rpc(t *testing.T) {

	// request
	c, _ := http.New(cosmos.RpcUrl)
	blockHeight := int64(848636)
	// decoding
	block, _ := c.Block(context.Background(), &blockHeight)
	spew.Dump(block)
}
