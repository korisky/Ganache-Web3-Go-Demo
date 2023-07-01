package main

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	"own.goerli.demo/event"
)

var (
	ctx = context.Background()
	//url = "https://goerli.infura.io/v3/3f0482cf4c3545dbabaeab75f414e467"
	url = ""
	// connect to own web3 node url
	client, _ = ethclient.DialContext(ctx, url)
)

// from https://goethereumbook.org/client/
func main() {

	//// get header number
	//header, _ := client.HeaderByNumber(ctx, nil)
	//fmt.Println("Latest Block Number:" + header.Number.String() + "\n")
	//fmt.Println("Latest Txn Hash:" + header.TxHash.String() + "\n")
	//
	//// call query balance
	//account.QueryBalance(ctx, client, "0xe10eE98bB84B2073B88353e3AB4433916205DF40", nil)
	//
	//// call transfer
	//transfer.QueryTransactions(ctx, client)
	//
	//// judge contract or address
	//account.JudgeAddressOrSmartContract(client, "0xe10eE98bB84B2073B88353e3AB4433916205DF40")
	//account.JudgeAddressOrSmartContract(client, "0xff02b7d59975E76F67B63b20b813a9Ec0f6AbD60")
	//
	//// send raw transfer
	//transfer.SendingRawTransfer(ctx, client)

	event.SubscribeErc1155SingleTransferEvent(ctx, client)

}
