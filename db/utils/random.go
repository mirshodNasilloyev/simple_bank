package utils

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min)
}

// RandomString generates random string of length l
func RandomString(l int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < l; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// RandomOwner generates random owner name
func RandomOwner() string {
	return RandomString(8)
}

// RandomMoney generates a random account of money
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// RandomCurrency generates a random currency code
func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "GBP", "JPY"}
	return currencies[rand.Intn(len(currencies))]
}
