package nft

import (
	"context"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/portto/solana-go-sdk/client"
	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/program/metaplex/token_metadata"
)

// TryDecodeMetadata try to decode nft metadata
func TryDecodeMetadata(cli *client.Client, nftAccAddress string) {
	// recover owner's meta account pub key
	nftMinter := common.PublicKeyFromString(nftAccAddress)
	nftMetaDataAccPubKey, _ := token_metadata.GetTokenMetaPubkey(nftMinter)

	// if we put user's account + nft token's mint account -> we could figure out the result of them are same
	fmt.Println("Nft token owner pub key -> ", nftMetaDataAccPubKey.ToBase58())

	// get info for retrieving meta data
	nftMetaDataAccountInfo, _ := cli.GetAccountInfo(context.Background(), nftMetaDataAccPubKey.ToBase58())
	metadata, _ := token_metadata.MetadataDeserialize(nftMetaDataAccountInfo.Data)

	// in Creators, the address who share more, that would be the owner
	fmt.Println("\nMeta Data -> ")
	spew.Dump(metadata)
}
