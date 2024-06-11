package util

import "math/rand"

const Alphabet string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// RandomString
// Generates a random string of length n
func RandomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = Alphabet[rand.Intn(len(b))]
	}
	return string(b)
}
