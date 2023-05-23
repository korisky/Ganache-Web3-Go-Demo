package cosmos

import (
	"context"
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"google.golang.org/grpc"

	"testing"
)

func Test_queryState(t *testing.T) {

	address, _ := sdk.AccAddressFromHexUnsafe("cosmos1wrljuqeelfuksll8clpqz5nxsqhdr7jcp0psv9")

	// Create a connection to the gRPC server.
	grpcConn, _ := grpc.Dial(
		"127.0.0.1:9090",    // your gRPC server address.
		grpc.WithInsecure(), // The Cosmos SDK doesn't support any transport security mechanism.
		// This instantiates a general gRPC codec which handles proto bytes. We pass in a nil interface registry
		// if the request/response types contain interface instead of 'nil' you should pass the application specific codec.
		grpc.WithDefaultCallOptions(grpc.ForceCodec(codec.NewProtoCodec(nil).GRPCCodec())),
	)
	defer grpcConn.Close()

	// This creates a gRPC client to query the x/bank service.
	bankClient := banktypes.NewQueryClient(grpcConn)
	bankRes, _ := bankClient.Balance(
		context.Background(),
		&banktypes.QueryBalanceRequest{Address: address.String(), Denom: "atom"},
	)

	fmt.Println(bankRes.GetBalance()) // Prints the account balance

}
