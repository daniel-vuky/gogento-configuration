package storage

import (
	"context"
	"github.com/daniel-vuky/gogento-configuration/core_config/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

// PostgresCoreConfigRepository provides all functions to execute db queries
type PostgresCoreConfigRepository struct {
	*Queries
	connPool *pgxpool.Pool
}

// NewRepository creates a new service
func NewRepository(connPool *pgxpool.Pool) *PostgresCoreConfigRepository {
	return &PostgresCoreConfigRepository{
		Queries:  New(connPool),
		connPool: connPool,
	}
}

const createConfig = `-- name: CreateConfig :one
INSERT
INTO core_config_data (path, value)
VALUES ($1, $2)
RETURNING config_id, path, value, created_at, updated_at
`

func (r *PostgresCoreConfigRepository) Create(ctx context.Context, arg *models.CreateConfigParams) (models.CoreConfigModel, error) {
	row := r.connPool.QueryRow(ctx, createConfig, arg.Path, arg.Value)
	var i models.CoreConfigModel
	err := row.Scan(
		&i.ConfigID,
		&i.Path,
		&i.Value,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteConfig = `-- name: DeleteConfig :one
DELETE
FROM core_config_data
Where path = $1
RETURNING config_id, path, value, created_at, updated_at
`

func (r *PostgresCoreConfigRepository) Delete(ctx context.Context, path string) (models.CoreConfigModel, error) {
	row := r.connPool.QueryRow(ctx, deleteConfig, path)
	var i models.CoreConfigModel
	err := row.Scan(
		&i.ConfigID,
		&i.Path,
		&i.Value,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getConfig = `-- name: GetConfig :one
SELECT config_id, path, value, created_at, updated_at
FROM core_config_data
Where path = $1
`

func (r *PostgresCoreConfigRepository) Get(ctx context.Context, path string) (models.CoreConfigModel, error) {
	row := r.connPool.QueryRow(ctx, getConfig, path)
	var i models.CoreConfigModel
	err := row.Scan(
		&i.ConfigID,
		&i.Path,
		&i.Value,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateConfig = `-- name: UpdateConfig :one
UPDATE core_config_data
SET value = $2
Where path = $1
RETURNING config_id, path, value, created_at, updated_at
`

func (r *PostgresCoreConfigRepository) Update(ctx context.Context, arg *models.UpdateConfigParams) (models.CoreConfigModel, error) {
	row := r.connPool.QueryRow(ctx, updateConfig, arg.Path, arg.Value)
	var i models.CoreConfigModel
	err := row.Scan(
		&i.ConfigID,
		&i.Path,
		&i.Value,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
