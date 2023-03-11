package balance

import (
	"context"
	"github.com/davecgh/go-spew/spew"
	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/token"
	"github.com/gagliardetto/solana-go/rpc"
)

func TryGetTokenAmount(cli *rpc.Client, accountAddress, tokenMintAddress string) {
	// keys
	accPublicKey := solana.MustPublicKeyFromBase58(accountAddress)
	tokenPubKey := solana.MustPublicKeyFromBase58(tokenMintAddress)
	// req
	resp, _ := cli.GetTokenAccountsByOwner(
		context.TODO(), accPublicKey,
		&rpc.GetTokenAccountsConfig{Mint: &tokenPubKey},
		&rpc.GetTokenAccountsOpts{Encoding: solana.EncodingBase64Zstd})
	// decoding
	for _, rawAccount := range resp.Value {
		// declare new variable for storing
		var tokAcc token.Account
		// try decoding
		binaryData := rawAccount.Account.Data.GetBinary()
		decodedData := bin.NewBinDecoder(binaryData)
		_ = decodedData.Decode(&tokAcc)
		// print
		spew.Dump(tokAcc)
	}
}
