package main

import (
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"
	"testing"
)

// Test_SchnorrDemonstration is a demonstration for Schnorr Algorithm
func Test_SchnorrDemonstration(t *testing.T) {

	// Generate a random Schnorr key pair
	curve := elliptic.P256() // Choose a curve
	privKey, _ := rand.Int(rand.Reader, curve.Params().N)
	pubKeyX, pubKeyY := curve.ScalarBaseMult(privKey.Bytes())
	pubKey := elliptic.Marshal(curve, pubKeyX, pubKeyY)

	// Compute the message hash
	message := []byte("hello, world")
	h := sha256.Sum256(message)

	// Generate a random nonce
	randNonce, _ := rand.Int(rand.Reader, curve.Params().N)

	// Compute the challenge
	nonceX, nonceY := curve.ScalarBaseMult(randNonce.Bytes())
	nonceBytes := elliptic.Marshal(curve, nonceX, nonceY)
	hasher := sha256.New()
	hasher.Write(pubKey)
	hasher.Write(h[:])
	hasher.Write(nonceBytes)
	cBytes := hasher.Sum(nil)
	c := new(big.Int).SetBytes(cBytes)

	// Compute the response
	s := new(big.Int).Mul(privKey, c)
	s.Sub(randNonce, s.Mod(s, curve.Params().N))

	// Verify the response
	responseX, responseY := curve.ScalarBaseMult(s.Bytes())
	responseBytes := elliptic.Marshal(curve, responseX, responseY)
	hasher.Reset()
	hasher.Write(pubKey)
	hasher.Write(h[:])
	hasher.Write(responseBytes)
	c1Bytes := hasher.Sum(nil)
	c1 := new(big.Int).SetBytes(c1Bytes)
	if c.Cmp(c1) == 0 {
		fmt.Println("Verification successful!")
	} else {
		fmt.Println("Verification failed!")
	}

}
