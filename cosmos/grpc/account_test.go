package grpc

import (
	"encoding/hex"
	"fmt"
	"github.com/cosmos/btcutil/bech32"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/go-bip39"
	"github.com/tendermint/tendermint/crypto/secp256k1"
	"log"
	"testing"
)

func Test_AccountAgain(t *testing.T) {
	fromBech32, _ := sdk.AccAddressFromBech32("osmo1xq9dsa599jurv2c8kj06duussplu9qmz0cxzru")
	fmt.Println(fromBech32)
}

// Test_AccountRecover try recover account from mnemonic (not working, much more longer)
func Test_AccountRecover(t *testing.T) {

	// Replace this with your Mnemonic Phrase
	mnemonic := "window turn course input bag clog elbow witness globe switch latin gasp"

	// Generating the seed from Mnemonic
	seed := bip39.NewSeed(mnemonic, "Thunder123LYK")

	// Deriving the private key
	master, ch := hd.ComputeMastersFromSeed(seed)
	priv, err := hd.DerivePrivateKeyForPath(master, ch, "m/44'/118'/0'/0/0'")
	if err != nil {
		log.Fatal(err)
	}

	// Generating the public key
	privKey := secp256k1.GenPrivKeySecp256k1(priv)
	pubKey := privKey.PubKey()

	// Generating the address
	decodeString, _ := hex.DecodeString(fmt.Sprintf("04%x", pubKey.Bytes()))
	conv, err := bech32.ConvertBits(decodeString, 8, 5, true)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// osmo1xq9dsa599jurv2c8kj06duussplu9qmz0cxzru
	encoded, err := bech32.Encode("osmo", conv)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Printing the address
	fmt.Println("Atom address:", encoded)
}
