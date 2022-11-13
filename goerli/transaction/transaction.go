package transaction

import (
	"bufio"
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"os"
)

func QueryTransactions(ctx context.Context, client *ethclient.Client) {
	block, _ := client.BlockByNumber(ctx, nil)
	for _, txn := range block.Transactions() {
		fmt.Println("Single Txn")
		fmt.Println(txn.Value().String())
		fmt.Println(txn.Gas())
		fmt.Println(txn.GasPrice().Uint64())
		fmt.Println(txn.Nonce())
		fmt.Println(txn.To().Hex())
		fmt.Println()
		break
	}
}

func SendingRawTransfer(ctx context.Context, client *ethclient.Client) {

	// buffer reader for input priKey
	r := bufio.NewReader(os.Stdin)
	fmt.Fprint(os.Stderr, "Please input your account pri-key\n")
	hexPriKey, _ := r.ReadString('\n')
	hexPriKey = hexPriKey[:len(hexPriKey)-1] // \n only represent 1 character

	// hex to using
	priKey, err := crypto.HexToECDSA(hexPriKey)
	if err != nil {
		log.Fatal(err)
	}

	// get public address from priKey
	pubKey := priKey.Public()
	pubKeyEcdsa, ok := pubKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Error on casting public key to ECDSA")
	}
	address := crypto.PubkeyToAddress(*pubKeyEcdsa)

	// get the next nonce
	nonce, err := client.PendingNonceAt(ctx, address)
	if err != nil {
		log.Fatal(err)
	}

	// transfer number initialisation (in EIP-1559)
	toAddress := common.HexToAddress("0x36F0A040C8e60974d1F34b316B3e956f509Db7e5")
	value := big.NewInt(2456)
	gasLimit := uint64(21000)
	tips := big.NewInt(2_000_000_000)    // max fee per gas 2-GWei
	feeCap := big.NewInt(20_000_000_000) // 20-GWei
	var data []byte
	chainId, err := client.NetworkID(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// fill txn
	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   chainId,
		Nonce:     nonce,
		GasTipCap: tips,
		GasFeeCap: feeCap,
		Gas:       gasLimit,
		To:        &toAddress,
		Value:     value,
		Data:      data,
	})

	// sign txn
	signTx, err := types.SignTx(tx, types.LatestSignerForChainID(chainId), priKey)
	if err != nil {
		log.Fatal(err)
	}

	// send txn
	err = client.SendTransaction(ctx, signTx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Finished sent transaction, with txHash: %s\n", signTx.Hash().Hex())
}
