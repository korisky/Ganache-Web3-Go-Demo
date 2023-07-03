package grpc

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/protobuf/proto"
	"log"

	"testing"
)

var mnemonic = ""

// Test_RecoverAccountByMnemonic recover account from mnemonic -> keyring -> account could divide public address
func Test_RecoverAccountByMnemonic(*testing.T) {

	// recover account by using mnemonic
	kb := keyring.NewInMemory(Create_Cdc())
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
