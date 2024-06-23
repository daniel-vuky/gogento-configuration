package models

import (
	"github.com/jackc/pgx/v5/pgtype"
	"time"
)

type CoreConfigModel struct {
	ConfigID  int64       `json:"config_id"`
	Path      string      `json:"path"`
	Value     pgtype.Text `json:"value"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}

type CreateConfigParams struct {
	Path  string      `json:"path"`
	Value pgtype.Text `json:"value"`
}

type UpdateConfigParams struct {
	Path  string      `json:"path"`
	Value pgtype.Text `json:"value"`
}
