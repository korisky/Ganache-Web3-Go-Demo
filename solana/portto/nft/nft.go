package nft

import (
	"context"
	"github.com/davecgh/go-spew/spew"
	"github.com/portto/solana-go-sdk/client"
	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/program/metaplex/token_metadata"
)

// TryDecodeNft try to decode nft metadata
func TryDecodeNft(cli *client.Client, nftAccAddress string) {
	// recover meta account pub key
	nftMinter := common.PublicKeyFromString(nftAccAddress)
	nftMetaDataAccount, _ := token_metadata.GetTokenMetaPubkey(nftMinter)

	// get info for retrieving meta data
	nftMetaDataAccountInfo, _ := cli.GetAccountInfo(context.Background(), nftMetaDataAccount.ToBase58())
	metadata, _ := token_metadata.MetadataDeserialize(nftMetaDataAccountInfo.Data)

	// in Creators, the address who share more, that would be the owner
	spew.Dump(metadata)
}
