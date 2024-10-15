package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var random *rand.Rand

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const numbers = "0123456789"
const alphanumbet = alphabet + numbers

func init() {
	random = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func RandomInt(min, max int64) int64 {
	return min + random.Int63n(max-min+1)
}

func RandomString(length int) string {
	return randomFunc(length, alphabet)
}

func RandomAlphaNumString(length int) string {
	return randomFunc(length-1, alphanumbet) + randomFunc(1, numbers)
}

func RandomNumString(length int) string {
	return randomFunc(length, numbers)
}

func RandomEmail() string {
	return fmt.Sprintf("%s@%s.%s", RandomString(6), RandomString(4), RandomString(3))
}

func randomFunc(length int, alphabet string) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < length; i++ {
		c := alphabet[random.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}
