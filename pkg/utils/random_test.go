package utils

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

// TestRandomString
// Test RandomString function
func TestRandomString(t *testing.T) {
	randomString := RandomString(5)
	require.NotEmpty(t, randomString)
	require.Equal(t, 5, len(randomString))
}

// TestRandomInt
// Test RandomInt function
func TestRandomInt(t *testing.T) {
	randomInt := RandomInt(1, 10)
	require.GreaterOrEqual(t, randomInt, 1)
	require.LessOrEqual(t, randomInt, 10)
}

// TestRandomEmail
// Test RandomEmail function
func TestRandomEmail(t *testing.T) {
	randomEmail := RandomEmail()
	require.NotEmpty(t, randomEmail)
	require.Contains(t, randomEmail, "@unit_test.com")
}

// TestRandomDate
// Test RandomDate function
func TestRandomDate(t *testing.T) {
	randomDate := RandomDate()
	require.NotEmpty(t, randomDate)
	require.IsType(t, time.Time{}, randomDate)
}
