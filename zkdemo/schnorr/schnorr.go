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
func SignSchnorr(curve elliptic.Curve, privateKey []byte, msg []byte) (r, s *big.Int) {
	hash := sha256.Sum256(msg)
	e := new(big.Int).SetBytes(hash[:])
	for {
		k, _ := rand.Int(rand.Reader, curve.Params().N)
		kGx, _ := curve.ScalarBaseMult(k.Bytes())
		r = new(big.Int).Mod(kGx, curve.Params().N)

		if r.Sign() == 0 {
			continue
		}

		kInv := new(big.Int).ModInverse(k, curve.Params().N)
		temp := new(big.Int).Mul(new(big.Int).SetBytes(privateKey), r)
		temp.Add(temp, e)
		s = temp.Mul(temp, kInv)
		s.Mod(s, curve.Params().N)

		if s.Sign() != 0 {
			break
		}
	}
	return r, s
}

// VerifySchnorr is for verifying
func VerifySchnorr(curve elliptic.Curve, publicKeyX, publicKeyY *big.Int, msg []byte, r, s *big.Int) bool {

	if r.Cmp(curve.Params().N) >= 0 || s.Cmp(curve.Params().N) >= 0 {
		return false
	}

	hash := sha256.Sum256(msg)
	e := new(big.Int).SetBytes(hash[:])

	w := new(big.Int).ModInverse(s, curve.Params().N)
	u1 := new(big.Int).Mul(e, w)
	u1.Mod(u1, curve.Params().N)
	u2 := new(big.Int).Mul(r, w)
	u2.Mod(u2, curve.Params().N)

	x1, y1 := curve.ScalarBaseMult(u1.Bytes())
	x2, y2 := curve.ScalarMult(publicKeyX, publicKeyY, u2.Bytes())
	x3, _ := curve.Add(x1, y1, x2, y2)

	return r.Cmp(x3) == 0
}
