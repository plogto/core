package util

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const defaultLetters = "abcdefghijklmnopqrstuvwxyz1234567890"

func RandomString(n int) string {
	return generateRandomString(n, nil)
}

func RandomHexString(n int) string {
	return RandomStringByLetters(n, "abcdef0123456789")
}

func RandomStringByLetters(n int, letters string) string {
	return generateRandomString(n, &letters)
}

func generateRandomString(n int, letters *string) string {
	b := make([]byte, n)
	var letterBytes string

	if letters != nil {
		letterBytes = *letters
	} else {
		letterBytes = defaultLetters
	}

	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
