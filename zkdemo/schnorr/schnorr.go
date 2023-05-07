package schnorr

import (
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"math/big"
)

// GenerateKeyPair is for generate random key-pairs under certain elliptic curve
func GenerateKeyPair(curve elliptic.Curve) ([]byte, *big.Int, *big.Int) {
	priKey, x, y, _ := elliptic.GenerateKey(curve, rand.Reader)
	return priKey, x, y
}

// SignSchnorr is for sign under schnorr algorithms' procedure
func SignSchnorr(curve elliptic.Curve, priKey []byte, x, y *big.Int, msg []byte) (r, s *big.Int) {
	hash := sha256.Sum256(msg)
	e := new(big.Int).SetBytes(hash[:])
	for {
		k, _ := rand.Int(rand.Reader, curve.Params().N)
		r, _ := curve.ScalarBaseMult(k.Bytes())
		s = new(big.Int).Mul(new(big.Int).SetBytes(priKey), r)
		s.Add(s, new(big.Int).Mul(k, e))
		s.Mod(s, curve.Params().N)

		if r.Sign() != 0 && s.Sign() != 0 {
			break
		}
	}
	return r, s
}
