package grpc

import (
	"context"
	"fmt"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/golang/protobuf/proto"

	cosmos "own.cosmos.demo"

	"log"
	"own.cosmos.demo/grpc/account"

	"testing"
)

// Test_RecoverAccountByMnemonic recover account from mnemonic -> keyring -> account could divide public address
func Test_RecoverAccountByMnemonic(*testing.T) {

	mnemonic := ""

	// recover account by using mnemonic
	kb := keyring.NewInMemory(account.CreateCdc())
	record, _ := kb.NewAccount("osmo", mnemonic, keyring.DefaultBIP39Passphrase, sdk.FullFundraiserPath, hd.Secp256k1)

	// Set the config for the Cosmos SDK.
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount("osmo", "osmo")

	// Extract the public key bytes from the record.
	pubKeyBytes := record.PubKey.Value

	// Unmarshal the bytes into the secp256k1.PubKey structure.
	var pubKey secp256k1.PubKey
	if err := proto.Unmarshal(pubKeyBytes, &pubKey); err != nil {
		log.Fatalf("Error unmarshalling public key: %v", err)
	}

	// Convert the secp256k1 public key to a Cosmos SDK AccAddress.
	addressBytes := sdk.AccAddress(pubKey.Address().Bytes())

	// Convert the AccAddress bytes to a Bech32 encoded string.
	addressString, err := sdk.Bech32ifyAddressBytes(config.GetBech32AccountAddrPrefix(), addressBytes)
	if err != nil {
		log.Fatalf("Error converting address bytes to bech32: %v", err)
	}

	fmt.Printf("Cosmos address (Bech32): %s\n", addressString)
}

// Test_KeyRecover try recover private key from mnemonic
func Test_KeyRecover(t *testing.T) {

	//theMnemonic := ""
	//
	//seed, err := bip39.NewSeedWithErrorChecking(theMnemonic, "")
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	//masterKey, chainCode := hd.ComputeMastersFromSeed(seed)
	//
	//config := sdk.GetConfig()
	//config.SetBech32PrefixForAccount("osmo", "osmo")
	//
	//priKey, err := hd.DerivePrivateKeyForPath(masterKey, chainCode, sdk.GetConfig().GetFullBIP44Path())
	//
	//privKeySecp256k1 := secp256k1.PrivKey{Key: priKey}
	//fmt.Println(privKeySecp256k1)
	//
	////client := tmservice.NewServiceClient(cosmos.Conn)
	////client.
	//
	//priv1, _, address1 := testdata.KeyTestPubAddr()
	//priv2, _, address2 := testdata.KeyTestPubAddr()
	//msg1 := banktypes.NewMsgSend(address1, address2, types.NewCoins(types.NewInt64Coin("atom", 1)))
}

// Test_AddressBalance is for query all balances
func Test_AddressBalance(t *testing.T) {
	defer cosmos.Conn.Close()
	queryClient := banktypes.NewQueryClient(cosmos.Conn)
	request := banktypes.QueryAllBalancesRequest{Address: "fx1xmc2qsxgucyhf50nfvckk054dagfmdl9ddkjlp"}

	balances, _ := queryClient.AllBalances(context.Background(), &request)
	fmt.Println(balances)
}

func Test_AccountSequence(t *testing.T) {
	defer cosmos.Conn.Close()
	client := authtypes.NewQueryClient(cosmos.Conn)
	request := authtypes.QueryAccountRequest{Address: "fx1xmc2qsxgucyhf50nfvckk054dagfmdl9ddkjlp"}

	response, _ := client.Account(context.Background(), &request)
	fmt.Println(response)
}

// prepareTx is for preparing a need txn before sign it
//func Test_SendTxn(t *testing.T) {
//
//
//	fromAddress, _ := sdk.AccAddressFromBech32("")
//	toAddress, _ := sdk.AccAddressFromBech32("")
//
//	sendMsg := banktypes.NewMsgSend(
//		fromAddress,
//		toAddress,
//		sdk.NewCoins(sdk.NewInt64Coin("FX", 3000000)))
//
//	// Create a transaction builder
//	txf := tx.Factory{}.
//		WithChainID(chainID).
//		WithGas(200000).
//		WithFees("1000FX"). // Set the fee here
//		WithMemo("").
//		WithAccountNumber(0).
//		WithSequence(0)
//}
