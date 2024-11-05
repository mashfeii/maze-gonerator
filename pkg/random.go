package pkg

import (
	"crypto/rand"
	"math/big"
)

func CryptoRandSecure(upperBound int64) int64 {
	nBig, _ := rand.Int(rand.Reader, big.NewInt(upperBound))
	return nBig.Int64()
}
