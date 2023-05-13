package utils

import (
	"crypto/rand"
	"math/big"
)

func GenerateRandomNumberString(length int) (string, error) {
	var numbers = []rune("1234567890")
	b := make([]rune, length)
	for i := range b {
		bigInt, errInt := rand.Int(rand.Reader, big.NewInt(int64(len(numbers))))
		if errInt != nil {
			return "", errInt
		}
		b[i] = numbers[bigInt.Int64()]
	}
	return string(b), nil
}
