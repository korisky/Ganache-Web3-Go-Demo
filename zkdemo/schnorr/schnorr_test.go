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
	privateKey, publicKeyX, publicKeyY := GenerateKeyPair(curve)

	msg := "test"
	signatureR, signatureS := SignSchnorr(curve, privateKey, []byte(msg))

	verify := VerifySchnorr(curve, publicKeyX, publicKeyY, []byte(msg), signatureR, signatureS)
	fmt.Println(verify)
}
