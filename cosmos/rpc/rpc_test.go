package rpc

import (
	"context"
	"fmt"
	"github.com/tendermint/tendermint/rpc/client/http"
	cosmos "own.cosmos.demo"
	"testing"
)

func Test_getStatus_Rpc(t *testing.T) {

	// request
	c, _ := http.New(cosmos.RpcUrl)

	// decoding
	resp, _ := c.Status(context.Background())

	fmt.Println(resp.SyncInfo.LatestBlockHeight)
}
