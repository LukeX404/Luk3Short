package utils

import (
	"math/rand"
	"time"
)

func RandomString() string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	rand.Seed(time.Now().UnixNano())

	size := rand.Intn(6) + 5

	b := make([]byte, size)

	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}

	return string(b)
}