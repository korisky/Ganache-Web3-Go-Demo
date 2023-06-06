package main

import (
	"context"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/portto/solana-go-sdk/program/token"
	"log"
	"testing"
	"web3Demo/portto/accounts"
	"web3Demo/portto/nft"
)

// Test_GetAllAssociatedTokenAccountByOwner get all associated token account by owner -> input correct owner
func Test_GetAllAssociatedTokenAccountByOwner(t *testing.T) {
	// could get all associated accounts & their: <SPL OwnAddress, SPL Mint Address, SPL Amount, SPL Owner>
	accounts.TryGetTokenAccountsByOwner(cli, accountAddress)
}

// Test_TryFindAssociatedTokenAddress is about derived associated token address by input owner address & token mint address
func Test_TryFindAssociatedTokenAddress(t *testing.T) {
	accounts.TryFindAssociatedTokenAddress("GQ6V9ZLVibN7eAtxEQxLJjXX8L9RybMJPpUCwi16vVgL", usdcTokenMintAddress)
}

// Test_AccountInfo get token info -> could find out -> token & nft data length different
func Test_AccountInfo(t *testing.T) {
	token_info, _ := cli.GetAccountInfo(context.Background(), usdcTokenMintAddress)
	spew.Dump(token_info)
	fmt.Println()
	nft_info, _ := cli.GetAccountInfo(context.Background(), nftMintAddress)
	spew.Dump(nft_info)
}

// Test_AccountInfoDecode -> here is the correct way to decode token account's account info
func Test_AccountInfoDecode(t *testing.T) {
	tokenInfo, _ := cli.GetAccountInfo(context.Background(), "4jfNnrE97f4CzUt6z9C36NqpFxpv5T9erKU9JsG6Kr5N")
	tokenAccount, err := token.TokenAccountFromData(tokenInfo.Data)
	if err != nil {
		log.Fatalf("Error on decoding token account info, %v", err)
	}
	spew.Dump(tokenAccount)
}

// Test_MetaplexNft decode nft-meta data -> input correct mint account for that specific nft/token
func Test_MetaplexNft(t *testing.T) {
	nft.TryDecodeMetadata(cli, "EZqtsCxYpYtNaX1Pd2ep3ZUVxS6qHLVQriugvbKGEahk") // is nft
	nft.TryDecodeMetadata(cli, "Gd8nxWzbnJ2zwtn5TukvEMKKjjbFhdtqA1L67DgnRvXc") // is not nft
}
