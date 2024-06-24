package usecase

import (
	"context"
	"github.com/daniel-vuky/gogento-configuration/core_config/postgres/models"
)

type Reader interface {
	GetConfig(ctx context.Context, path string) (models.CoreConfigModel, error)
}

type Writer interface {
	CreateConfig(ctx context.Context, arg *models.CreateConfigParams) (models.CoreConfigModel, error)
	DeleteConfig(ctx context.Context, path string) (models.CoreConfigModel, error)
	UpdateConfig(ctx context.Context, arg *models.UpdateConfigParams) (models.CoreConfigModel, error)
}

type CoreConfigUseCase interface {
	Reader
	Writer
}
