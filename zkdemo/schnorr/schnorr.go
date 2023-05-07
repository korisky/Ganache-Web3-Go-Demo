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
	// get hash bytes of the msg and convert to a fixed length digest
	hash := sha256.Sum256(msg)
	e := new(big.Int).SetBytes(hash[:])

	for {
		// Generate a random nonce 'k' in the range [1, n-1], where 'n' is the order of the curve
		k, _ := rand.Int(rand.Reader, curve.Params().N)

		// Compute the point (kGx, kGy) = k * G, where G is the base point of the curve
		kGx, _ := curve.ScalarBaseMult(k.Bytes())

		// Set 'r' as the x-coordinate of the point -> (kGx, kGy) modulo n
		// in some case r would result in 0, need to loop to try another random k
		r = new(big.Int).Mod(kGx, curve.Params().N)
		if r.Sign() == 0 {
			continue
		}

		// Calculate the modular inverse of k
		kInv := new(big.Int).ModInverse(k, curve.Params().N)

		// Compute s = (privateKey * r + e) * k^-1 mod n
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

	// Check if r and s are in the valid range [1, n-1]
	if r.Cmp(curve.Params().N) >= 0 || s.Cmp(curve.Params().N) >= 0 {
		return false
	}

	// Hash the message to create a fixed-length digest
	hash := sha256.Sum256(msg)
	e := new(big.Int).SetBytes(hash[:])

	// Compute w = s^-1 mod n
	w := new(big.Int).ModInverse(s, curve.Params().N)

	// Compute u1 = e * w mod n
	u1 := new(big.Int).Mul(e, w)
	u1.Mod(u1, curve.Params().N)

	// Compute u2 = r * w mod n
	u2 := new(big.Int).Mul(r, w)
	u2.Mod(u2, curve.Params().N)

	// Calculate the point (x1, y1) = u1 * G, where G is the base point of the curve
	x1, y1 := curve.ScalarBaseMult(u1.Bytes())

	// Calculate the point (x2, y2) = u2 * (publicKeyX, publicKeyY)
	x2, y2 := curve.ScalarMult(publicKeyX, publicKeyY, u2.Bytes())

	// Calculate the point (x3, y3) = (x1, y1) + (x2, y2)
	x3, _ := curve.Add(x1, y1, x2, y2)

	// The signature is valid if r is equal to x3
	return r.Cmp(x3) == 0
}
