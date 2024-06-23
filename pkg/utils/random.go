package utils

import (
	"math/rand"
	"time"
)

const Alphabet string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

// RandomString
// Generates a random string of length n
func RandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = Alphabet[rand.Intn(len(Alphabet))]
	}
	return string(b)
}

// RandomInt
// Generates a random integer between min and max
func RandomInt(min, max int) int {
	return rand.Intn(max-min) + min
}

// RandomEmail
// Generates a random email
func RandomEmail() string {
	return RandomString(5) + "@unit_test.com"
}

// RandomDate
// Generates a random date
func RandomDate() time.Time {
	return time.Now().AddDate(-RandomInt(18, 50), 0, 0)
}
