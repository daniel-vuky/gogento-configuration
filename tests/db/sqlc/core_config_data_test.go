package sqlc

import (
	"context"
	db "github.com/daniel-vuky/gogento-configuration/db/sqlc"
	"github.com/daniel-vuky/gogento-configuration/util"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
	"testing"
)

// compareConfig
// Compare two CoreConfigDatum
func compareConfig(
	t *testing.T,
	err error,
	fromConfig *db.CoreConfigDatum,
	targetConfig *db.CoreConfigDatum,
) {
	require.NoError(t, err)
	require.NotEmpty(t, fromConfig)
	require.NotEmpty(t, targetConfig)
	require.Equal(t, fromConfig.Path, targetConfig.Path)
	require.Equal(t, fromConfig.Value, targetConfig.Value)
}

// CreateRandomConfig
// Test the CreateConfig function
func CreateRandomConfig(t *testing.T) db.CoreConfigDatum {
	arg := &db.CreateConfigParams{
		Path: util.RandomString(10),
		Value: pgtype.Text{
			String: util.RandomString(10),
			Valid:  true,
		},
	}
	config, err := testStore.CreateConfig(context.Background(), arg)
	compareConfig(t, err, &db.CoreConfigDatum{
		Path:  arg.Path,
		Value: arg.Value,
	}, &config)

	return config
}

// TestCreateConfig
// Test the CreateConfig function
func TestCreateConfig(t *testing.T) {
	// Success case
	CreateRandomConfig(t)
}

// TestCreateConfigValueEmpty
// Test the CreateConfig function
func TestCreateConfigValueEmpty(t *testing.T) {
	// Path empty case
	arg := &db.CreateConfigParams{
		Path: "",
		Value: pgtype.Text{
			String: util.RandomString(10),
			Valid:  true,
		},
	}
	config, err := testStore.CreateConfig(context.Background(), arg)
	require.Error(t, err)
	require.Empty(t, config)
}

// TestCreateConfig
// Test the CreateConfig function
func TestDeleteConfig(t *testing.T) {
	randomConfig := CreateRandomConfig(t)
	deletedConfig, err := testStore.DeleteConfig(context.Background(), randomConfig.Path)
	compareConfig(t, err, &randomConfig, &deletedConfig)
}

// TestGetConfig
// Test the GetConfig function
func TestGetConfig(t *testing.T) {
	randomConfig := CreateRandomConfig(t)
	config, err := testStore.GetConfig(context.Background(), randomConfig.Path)
	compareConfig(t, err, &randomConfig, &config)
}

// TestUpdateConfig
// Test the UpdateConfig function
func TestUpdateConfig(t *testing.T) {
	randomConfig := CreateRandomConfig(t)
	arg := &db.UpdateConfigParams{
		Path: randomConfig.Path,
		Value: pgtype.Text{
			String: util.RandomString(12),
			Valid:  true,
		},
	}
	config, err := testStore.UpdateConfig(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, config)
	require.Equal(t, arg.Path, config.Path)
	require.NotEqual(t, randomConfig.Value, config.Value)
}
