package cosmos

import (
	"fmt"
	"io"
	"net/http"

	"testing"
)

func Test_getStatus(t *testing.T) {

	// Osmosis Testnet Rpc
	rpcUrl := "https://rpc.osmotest5.osmosis.zone"

	resp, err := http.Get(rpcUrl + "/status")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(body)

}
