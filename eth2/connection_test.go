package eth2

import (
	"context"
	"fmt"
	"github.com/attestantio/go-eth2-client/http"
	"github.com/rs/zerolog"
	"testing"
)

var (
	nodeUrl = "https://dry-special-paper.ethereum-goerli.discover.quiknode.pro/2b769a89d976b40e56a1f38ebac60e6a04bb28d2"
)

func Test_Eth2Connection(t *testing.T) {

	// create client
	ctx, _ := context.WithCancel(context.Background())
	client, err := http.New(ctx, http.WithAddress(nodeUrl), http.WithLogLevel(zerolog.DebugLevel))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Connected to %s\n", client.Name())

	//

}
