package util

import (
	"github.com/daniel-vuky/gogento-configuration/util"
	"github.com/stretchr/testify/require"
	"testing"
)

// TestRandomString
// Test the RandomString function
func TestRandomString(t *testing.T) {
	randomString := util.RandomString(10)
	require.NotEmpty(t, randomString)
	require.Equal(t, 10, len(randomString))
}
