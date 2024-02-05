package util

import (
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func RandomString(n int) string {
	var builder strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		builder.WriteByte(c)
	}
	return builder.String()
}
