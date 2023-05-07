package schnorr

import (
	"crypto/elliptic"
	"crypto/rand"
	"math/big"
)

func GenerateKeyPair(curve elliptic.Curve) ([]byte, *big.Int, *big.Int) {
	priKey, x, y, _ := elliptic.GenerateKey(curve, rand.Reader)
	return priKey, x, y
}
