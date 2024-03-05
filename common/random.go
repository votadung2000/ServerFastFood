package common

import (
	"math/rand"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSequence(n int) string {
	b := make([]rune, n)

	for i := range b {
		source := rand.NewSource(time.Now().UnixNano())
		randSource := rand.New(source)

		b[i] = letters[randSource.Intn(999999)%len(letters)]
	}

	return string(b)
}

func GenSalt(length int) string {
	if length < 0 {
		length = 50
	}

	return randSequence(length)
}
