package repository

import (
	"context"
	"github.com/daniel-vuky/gogento-configuration/core_config/models"
)

type Reader interface {
	Get(ctx context.Context, path string) (models.CoreConfigModel, error)
}

type Writer interface {
	Create(ctx context.Context, arg *models.CreateConfigParams) (models.CoreConfigModel, error)
	Delete(ctx context.Context, path string) (models.CoreConfigModel, error)
	Update(ctx context.Context, arg *models.UpdateConfigParams) (models.CoreConfigModel, error)
}

type CoreConfigRepository interface {
	Reader
	Writer
}
