package storage

import (
	"context"
	"github.com/daniel-vuky/gogento-configuration/core_config/postgres/models"
	"github.com/daniel-vuky/gogento-configuration/pkg/utils"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
	"testing"
)

// compareConfig
// Compare two CoreConfigDatum
func compareConfig(
	t *testing.T,
	err error,
	fromConfig *models.CoreConfigModel,
	targetConfig *models.CoreConfigModel,
) {
	require.NoError(t, err)
	require.NotEmpty(t, fromConfig)
	require.NotEmpty(t, targetConfig)
	require.Equal(t, fromConfig.Path, targetConfig.Path)
	require.Equal(t, fromConfig.Value, targetConfig.Value)
}

// CreateRandomConfig
// Test the CreateConfig function
func CreateRandomConfig(t *testing.T) models.CoreConfigModel {
	arg := &models.CreateConfigParams{
		Path: utils.RandomString(10),
		Value: pgtype.Text{
			String: utils.RandomString(10),
			Valid:  true,
		},
	}
	config, err := postgresRepository.Create(context.Background(), arg)
	compareConfig(t, err, &models.CoreConfigModel{
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
	arg := &models.CreateConfigParams{
		Path: "",
		Value: pgtype.Text{
			String: utils.RandomString(10),
			Valid:  true,
		},
	}
	config, err := postgresRepository.Create(context.Background(), arg)
	require.Error(t, err)
	require.Empty(t, config)
}

// TestCreateConfig
// Test the CreateConfig function
func TestDeleteConfig(t *testing.T) {
	randomConfig := CreateRandomConfig(t)
	deletedConfig, err := postgresRepository.Delete(context.Background(), randomConfig.Path)
	compareConfig(t, err, &randomConfig, &deletedConfig)
}

// TestGetConfig
// Test the GetConfig function
func TestGetConfig(t *testing.T) {
	randomConfig := CreateRandomConfig(t)
	config, err := postgresRepository.Get(context.Background(), randomConfig.Path)
	compareConfig(t, err, &randomConfig, &config)
}

// TestUpdateConfig
// Test the UpdateConfig function
func TestUpdateConfig(t *testing.T) {
	randomConfig := CreateRandomConfig(t)
	arg := &models.UpdateConfigParams{
		Path: randomConfig.Path,
		Value: pgtype.Text{
			String: utils.RandomString(12),
			Valid:  true,
		},
	}
	config, err := postgresRepository.Update(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, config)
	require.Equal(t, arg.Path, config.Path)
	require.NotEqual(t, randomConfig.Value, config.Value)
}
