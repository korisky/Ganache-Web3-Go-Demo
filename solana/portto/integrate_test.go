package main

import (
	"testing"
	"web3Demo/portto/accounts"
	"web3Demo/portto/nft"
)

// Test_GetAllAssociatedTokenAccountByOwner get all associated token account by owner -> input correct owner
func Test_GetAllAssociatedTokenAccountByOwner(t *testing.T) {
	// could get all associated accounts & their: <SPL OwnAddress, SPL Mint Address, SPL Amount, SPL Owner>
	accounts.TryGetTokenAccountsByOwner(cli, accountAddress)
}

// Test_DecodeMetaplexNft decode nft-meta data -> input correct mint account for that specific nft/token
func Test_DecodeMetaplexNft(t *testing.T) {
	nft.TryDecodeMetadata(cli, "EZqtsCxYpYtNaX1Pd2ep3ZUVxS6qHLVQriugvbKGEahk") // is nft
	nft.TryDecodeMetadata(cli, "Gd8nxWzbnJ2zwtn5TukvEMKKjjbFhdtqA1L67DgnRvXc") // is not nft
}
