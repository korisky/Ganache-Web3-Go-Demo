package eth2

import (
	"context"
	"fmt"
	eth2client "github.com/attestantio/go-eth2-client"
	"github.com/attestantio/go-eth2-client/http"
	"github.com/davecgh/go-spew/spew"
	"github.com/rs/zerolog"
	"testing"
)

var (
	nodeUrl = "https://weathered-restless-general.ethereum-goerli.discover.quiknode.pro/acb676ab0fbacff1f79f83f792c78b20839e720f/"
)

func Test_Eth2Connection(t *testing.T) {

	// create client
	ctx, cancel := context.WithCancel(context.Background())
	client, err := http.New(ctx, http.WithAddress(nodeUrl), http.WithLogLevel(zerolog.DebugLevel))
	if err != nil {
		panic(err)
	}
	fmt.Printf("\nConnected to %s\n", client.Name())

	// client functions have their own interfaces, not all functions are supported by all clients
	if provider, isProvider := client.(eth2client.GenesisProvider); isProvider {
		// request for genesis time
		genesis, err := provider.Genesis(ctx)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Genesis time is %v\n\n", genesis.GenesisTime)
	}

	if provider, isProvider := client.(eth2client.BeaconBlockHeadersProvider); isProvider {
		// request for block
		header, err := provider.BeaconBlockHeader(ctx, "202980")
		if err != nil {
			panic(err)
		}
		fmt.Println("Beacon Block header: ")
		spew.Dump(header)
	}

	// cancelling the context passed to New() frees up resources held by the
	// client, closes connections, clears handlers, etc.
	cancel()
}
