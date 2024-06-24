package config

import (
	"github.com/stretchr/testify/require"
	"testing"
)

// TestLoadConfigSuccess
// Test the LoadConfig function
func TestLoadConfigSuccess(t *testing.T) {
	loadedConfig, err := LoadConfig("../")
	require.NoError(t, err)
	require.NotEmpty(t, loadedConfig)
}

// TestGetDatabaseSource
// Test the GetDatabaseSource method
func TestGetDatabaseSource(t *testing.T) {
	loadedConfig, err := LoadConfig("../")
	require.NoError(t, err)
	require.NotEmpty(t, loadedConfig)
	require.NotEmpty(t, loadedConfig.GetDatabaseSource())
}
