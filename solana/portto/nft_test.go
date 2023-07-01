package main

import (
	"context"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/program/metaplex/token_metadata"
	"testing"
)

func Test_NFTRelated(t *testing.T) {
	// 1. get nft owner's metadata account public key
	nftMinter := common.PublicKeyFromString("7rYSA4K3GyZpC1vmMcSYsTq4DhM1GyXWDcHxm2UcA3eU")
	metaDataAccPubKey, _ := token_metadata.GetTokenMetaPubkey(nftMinter)
	fmt.Println("Nft token owner's public account -> ", metaDataAccPubKey.ToBase58())

	// 2. get info for retrieving metadata
	accountInfo, _ := cli.GetAccountInfo(context.Background(), metaDataAccPubKey.ToBase58())
	metadata, _ := token_metadata.MetadataDeserialize(accountInfo.Data)
	spew.Dump(metadata)
}
