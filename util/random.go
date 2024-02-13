package util

import (
	"math/rand"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

func RandomString(n int) string {
	var result string
	for i := 0; i < n; i++ {
		result += string(alphabet[rand.Intn(len(alphabet))])
	}
	return result
}

func RandomBool() bool {
	return rand.Intn(2) == 1
}
