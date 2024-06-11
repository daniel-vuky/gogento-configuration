package util

import (
	"github.com/daniel-vuky/gogento-configuration/util"
	"github.com/stretchr/testify/require"
	"testing"
)

// TestLoadConfig
// Test the LoadConfig function
func TestLoadConfig(t *testing.T) {
	// Fail case
	config, err := util.LoadConfig("../")
	require.Error(t, err)
	require.Empty(t, config)

	// Success case
	config, err = util.LoadConfig("../../")
	require.NoError(t, err)
	require.NotEmpty(t, config)
}
