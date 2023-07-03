package grpc

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
)

// CreateCdc is for getting simple crypto code compiler
func CreateCdc() *codec.ProtoCodec {
	registry := codectypes.NewInterfaceRegistry()
	cryptocodec.RegisterInterfaces(registry)
	return codec.NewProtoCodec(registry)
}

// SendTransaction is for simple transfer
func SendTransaction() {

}
