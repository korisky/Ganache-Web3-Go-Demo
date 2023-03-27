package main

import (
	"github.com/portto/solana-go-sdk/client"
	"github.com/portto/solana-go-sdk/rpc"
	"self/solana.demo/m/httpProxy"
	"self/solana.demo/m/portto/accounts"
)

var (
	// own accountAddress account
	accountAddress           = "AnayTW335MabjhtXTJeBit5jdLhNeUVBVPXeRKCid79D"
	assTokenAccountAddress   = "Gd8nxWzbnJ2zwtn5TukvEMKKjjbFhdtqA1L67DgnRvXc"
	tokenMintAddress         = "Gh9ZwEmdLJ8DscKNTkTqPbNwLNNBjuSzaG9Vp2KGtKJr"
	nftMintAddress           = "EZqtsCxYpYtNaX1Pd2ep3ZUVxS6qHLVQriugvbKGEahk"
	nftCollectionMintAddress = "ARHE7qXefr79DqyApiEkZ2QwnyzfAnUew4jRXfkMBVT2"
	ownEndpoint              = "https://solana-devnet.g.alchemy.com/v2/On35d8LdFc1QGYD-wCporecGj359qian"
	// custom connection would be:
	cli = client.New(rpc.WithEndpoint(ownEndpoint), rpc.WithHTTPClient(httpProxy.GetHttpClient()))
)

// main are following stuff from https://portto.github.io/solana-go-sdk/tour/token-transfer.html
func main() {

	accounts.TryFindAssociatedTokenAddress("GCTFHbBcLs3vsNkCKLz2uCq5boZhMHbuawuFuhbYttc8", tokenMintAddress)

	//balance.TryGetBalance(cli, accountAddress)

	//tokenBalance, u := balance.TryGetTokenBalance(cli, "Gd8nxWzbnJ2zwtn5TukvEMKKjjbFhdtqA1L67DgnRvXc")
	//spew.Dump(tokenBalance, u)

	// get all accounts by this owner
	//accounts.TryGetTokenAccountsByOwner(cli, accountAddress)
	// decode nft mint account
	//nft.TryDecodeMetadata(cli, nftMintAddress)
	//nft.TryDecodeMetadata(cli, nftCollectionMintAddress)

	// decode token account
	//account, _ := cli.GetTokenAccount(context.Background(), assTokenAccountAddress)
	//spew.Dump(account)

	//accounts.TryFindAssociatedTokenAddress(accountAddress, nftMintAddress)

	//transfer.TryTransferToken(cli, "", assTokenAccountAddress, tokenMintAddress)

	// decode block
	//block := block.TryGetBlockByBlockHeight(cli, 202616989)
	//spew.Dump(block)

	// get txn
	//transaction, _ := cli.GetTransaction(context.Background(), "2FjiGVncyv1SWpGsYVx2yYegUdipTtgWFMnU6kfjZVZF69Y2afyh6GJ6eLofjhkUSxCpdudJiqdLJbU7haynyugC")
	//spew.Dump(transaction)R

	//accounts.TryFindTxnSigByOwnerAddress(cli, accountAddress, 10)

	//associatedAccounts, _ := cli.GetTokenAccountsByOwner(context.Background(), accountAddress)
	//spew.Dump(associatedAccounts)
	//config, _ := cli.GetSignaturesForAddressWithConfig(context.Background(), "AnayTW335MabjhtXTJeBit5jdLhNeUVBVPXeRKCid79D", rpc.GetSignaturesForAddressConfig{Limit: 10})
	//spew.Dump(config)
}
