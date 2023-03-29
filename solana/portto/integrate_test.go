package main

import (
	"testing"
	"web3Demo/portto/accounts"
	"web3Demo/portto/nft"
)

// Test_GetAllAssociatedTokenAccountByOwner get all associated token account by owner -> input correct owner
func Test_GetAllAssociatedTokenAccountByOwner(t *testing.T) {
	// could get all associated accounts & their: <SPL Mint Address, SPL Amount, SPL Owner>
	accounts.TryGetTokenAccountsByOwner(cli, "6JtKTjiumR3GiXp1fXgjDWgQZLHEfv3WTMuR9fNRLhun")
}

func Test_DecodeMetaplexNft(t *testing.T) {
	nft.TryDecodeMetadata(cli, "HtLbwcFnQD5J2RFoSK2uJPuGbJ1bgPbPpZEEE2WY7Q5c") // is nft
	nft.TryDecodeMetadata(cli, "4k3Dyjzvzp8eMZWUXbBCjEvwSkkk59S5iCNLY3QrkX6R") // is not nft
}
