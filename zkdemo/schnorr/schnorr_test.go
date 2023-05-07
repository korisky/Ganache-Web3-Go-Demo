package schnorr

import (
	"crypto/elliptic"
	"fmt"
	"testing"
)

// Test_SchnorrDemonstration is a demonstration for Schnorr Algorithm
func Test_SchnorrDemonstration(t *testing.T) {
	// choose a curve
	curve := elliptic.P256()
	// random private & public key pair
	privateKey, publicKeyX, publicKeyY := GenerateKeyPair(curve)

	fmt.Println(privateKey)
	fmt.Println(publicKeyX)
	fmt.Println(publicKeyY)
}
