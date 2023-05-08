package main

import (
	"github.com/portto/solana-go-sdk/client"
	"github.com/portto/solana-go-sdk/rpc"
	"web3Demo/httpProxy"
)

var (
	// own accountAddress account
	accountAddress           = "AnayTW335MabjhtXTJeBit5jdLhNeUVBVPXeRKCid79D"
	assTokenAccountAddress   = "Gd8nxWzbnJ2zwtn5TukvEMKKjjbFhdtqA1L67DgnRvXc"
	usdcTokenMintAddress     = "Gh9ZwEmdLJ8DscKNTkTqPbNwLNNBjuSzaG9Vp2KGtKJr"
	nftMintAddress           = "EZqtsCxYpYtNaX1Pd2ep3ZUVxS6qHLVQriugvbKGEahk"
	nftCollectionMintAddress = "ARHE7qXefr79DqyApiEkZ2QwnyzfAnUew4jRXfkMBVT2"
	ownEndpoint              = "https://solana-devnet.g.alchemy.com/v2/On35d8LdFc1QGYD-wCporecGj359qian"
	//ownEndpoint = "https://solana-mainnet.g.alchemy.com/v2/LoSsmI5Mt0lWufIzKRRq5Ct8IVNUQvdE"
	// custom connection would be:
	cli = client.New(rpc.WithEndpoint(ownEndpoint), rpc.WithHTTPClient(httpProxy.GetHttpClient()))
)

// main are following stuff from https://portto.github.io/solana-go-sdk/tour/token-transfer.html
func main() {

	//balance.TryGetBalance(cli, accountAddress)

	//tokenBalance, u := balance.TryGetTokenBalance(cli, "Gd8nxWzbnJ2zwtn5TukvEMKKjjbFhdtqA1L67DgnRvXc")
	//spew.Dump(tokenBalance, u)

	// get all accounts by this owner
	//accounts.TryGetTokenAccountsByOwner(cli, "6JtKTjiumR3GiXp1fXgjDWgQZLHEfv3WTMuR9fNRLhun")

	// get account info
	//info, _ := cli.GetAccountInfo(context.Background(), "9sMYz9FYrzkpVKqG55yVproUdMmGNxAoM7z9637gwDuh")
	//spew.Dump(info)

	// decode nft mint account
	//nft.TryDecodeMetadata(cli, "9sMYz9FYrzkpVKqG55yVproUdMmGNxAoM7z9637gwDuh")
	//nft.TryDecodeMetadata(cli, nftCollectionMintAddress)

	// decode token account
	//account, _ := cli.GetTokenAccount(context.Background(), assTokenAccountAddress)
	//spew.Dump(account)

	//transfer.TryTransferToken(cli, "", assTokenAccountAddress, usdcTokenMintAddress)

	// decode block
	//block := block.TryGetBlockByBlockHeight(cli, 202616989)
	//spew.Dump(block)

	//accounts.TryFindTxnSigByOwnerAddress(cli, accountAddress, 10)

	//associatedAccounts, _ := cli.GetTokenAccountsByOwner(context.Background(), accountAddress)
	//spew.Dump(associatedAccounts)
	//config, _ := cli.GetSignaturesForAddressWithConfig(context.Background(), "AnayTW335MabjhtXTJeBit5jdLhNeUVBVPXeRKCid79D", rpc.GetSignaturesForAddressConfig{Limit: 10})
	//spew.Dump(config)

	// get txn
	//transaction, _ := cli.GetTransaction(context.Background(), "41ZCAJiXCHmXjtGTm5VtB5q1F1eWoGa84xBBiErdDsqf7K3m2yn4GACjuXzimMg6zjtGA2MGuu9gPRmRqA6bV3ka")
	////transaction, _ := cli.GetTransaction(context.Background(), "2BghzZxWRXR7zy9by3GHGhi32yYbrwY59ivvdNzfMCcuHJ3FyFVkLy7bxkxKqcdrKVpaszvdSoCrjPSyAgMQiTts")
	//for _, instruction := range transaction.Transaction.Message.Instructions {
	//	transferInstruction, _ := decodeTransferInstruction(instruction.Data)
	//	if nil != transferInstruction {
	//		fmt.Printf("Amount %v", transferInstruction.Amount)
	//	}
	//}
	//
	//spew.Dump(transaction)
}
