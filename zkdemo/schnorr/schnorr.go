package schnorr

import (
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"math/big"
)

// GenerateKeyPair is for generate random key-pairs under certain elliptic curve
// 使用ECDSA之前，在椭圆曲线上随机选择一个Generator, 也就是开始点,
// 并移动privateKey次, 得到publicKey的坐标x和y
func GenerateKeyPair(curve elliptic.Curve) ([]byte, *big.Int, *big.Int) {
	priKey, x, y, _ := elliptic.GenerateKey(curve, rand.Reader)
	return priKey, x, y
}

// SignSchnorr is for sign under schnorr algorithms' procedure
func SignSchnorr(curve elliptic.Curve, privateKey []byte, msg []byte) (r, s *big.Int) {

	// get hash bytes of the msg and convert to a fixed length digest
	// 对需要进行签名的数据进行hash, 并且将其hash出来的值在后续添加Signature的计算中
	// s = (privateKey * r + e) * k^-1 mod n
	hash := sha256.Sum256(msg)
	e := new(big.Int).SetBytes(hash[:])

	for {
		// Generate a random nonce 'k' in the range [1, n-1], where 'n' is the order of the curve
		// 定义域在有穷数的ECC有一个特殊的性质：有一个整数“n”（群的阶数）使得对于曲线上的所有点 P，n * P 等于无穷远点。
		// 这个无穷远点作为椭圆曲线群中的恒等元；它相当于常规整数算术中的零。
		// 而根据ECC的标量乘法Scalar multiplication -> nP = P+P+P+到n次, 可以计算出
		// (k+n)*G = k*G + n*G
		//		   = k*G + 0(infinity)
		//		   = k*G
		// 所以相当于一个循环, 为了避免循环, 我们需要确保随机出来的k在范围[1,n-1]
		k, _ := rand.Int(rand.Reader, curve.Params().N)

		// Compute the point (kGx, kGy) = k * G, where G is the base point of the curve
		// 对点进行k次移动计算, 得到k作为PriKey的endPoint, 只关心其x值
		kGx, _ := curve.ScalarBaseMult(k.Bytes())

		// Set 'r' as the x-coordinate of the point -> (kGx, kGy) modulo n
		// in some case r would result in 0, need to loop to try another random k
		// r 是从k的公钥mod而来 -> r = kGx mod n
		// 这里进行modulo的意义在于, ECC本身并不在finite field，通过module可以保持在有穷域中,
		// 并且添加安全性, 让接收方也不容易反向推导
		r = new(big.Int).Mod(kGx, curve.Params().N)
		if r.Sign() == 0 {
			continue
		}

		// Calculate the modular inverse of k
		kInv := new(big.Int).ModInverse(k, curve.Params().N)

		// Compute s = (privateKey * r + e) * k^-1 mod n
		// 由于后续的验证, 是通过验证 s * G = R + e * PublicKey 可以推导 ->
		// 							   	      s * G = R + e * PublicKey
		//          (privateKey * r + e)/k * G = R + e * PublicKey   -> 由于 R = k * G
		// privateKey * r * k^-1 * G + e * k * G = R + e * PublicKey   -> 由于 k^-1 * G = R 和 privateKey * G = PublicKey
		// 					  PublicKey * R + e * R = R + e * PublicKey

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
