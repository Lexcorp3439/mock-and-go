package testutils

import (
	"crypto/rand"
	"math/big"
	"strconv"
)

func GetRandomInt(max int64) int64 {
	n, _ := rand.Int(rand.Reader, big.NewInt(max))
	return n.Int64()
}

func GetRandomPhone() string {
	result := "+"
	for i := 0; i < 11; i++ {
		result += strconv.FormatInt(GetRandomInt(10), 10)
	}
	return result
}
