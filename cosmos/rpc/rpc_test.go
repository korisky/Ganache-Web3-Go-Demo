package rpc

import (
	"context"
	"fmt"
	"github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/davecgh/go-spew/spew"
	"github.com/golang/protobuf/proto"
	"github.com/tendermint/tendermint/rpc/client/http"
	cosmos "own.cosmos.demo"
	"runtime"
	"testing"
)

func Test_getLatestBlockHeight_Rpc(t *testing.T) {
	// request
	c, _ := http.New(cosmos.RpcUrl)
	// decoding
	resp, _ := c.Status(context.Background())
	fmt.Println(resp.SyncInfo.LatestBlockHeight)
}

func Test_decodeBlock_Rpc(t *testing.T) {

	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Println("Init heap allocation: ", m.HeapAlloc)

	// request
	c, _ := http.New(cosmos.RpcUrl)
	blockHeight := int64(2124985)
	// decoding
	block, _ := c.Block(context.Background(), &blockHeight)
	// unmarshal all proto tx
	for _, txBytes := range block.Block.Data.Txs {
		var tx tx.Tx
		err := proto.Unmarshal(txBytes, &tx)
		if err != nil {
			t.Fatal(err)
		}
		spew.Dump(tx)
	}

	runtime.GC()
	runtime.ReadMemStats(&m)
	fmt.Println("After heap allocation: ", m.HeapAlloc)

}
