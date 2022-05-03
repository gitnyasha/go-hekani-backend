package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghkmnpqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomString generates a random string of the given length.
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString() string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < 8; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// random user email
func RandomEmail() string {
	return fmt.Sprintf("%s@mail.com", RandomString())
}

// random year
func RandomDateYear() int {
	return rand.Intn(1000)
}
