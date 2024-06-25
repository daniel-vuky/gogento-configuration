package service

import (
	"context"
	"github.com/daniel-vuky/gogento-configuration/core_config/postgres/models"
	"github.com/daniel-vuky/gogento-configuration/core_config/postgres/repository"
)

type Service struct {
	Repo repository.CoreConfigRepository
}

// NewService create new service
func NewService(r repository.CoreConfigRepository) *Service {
	return &Service{
		Repo: r,
	}
}

func (s *Service) GetConfig(ctx context.Context, path string) (models.CoreConfigModel, error) {
	return s.Repo.Get(ctx, path)
}

func (s *Service) CreateConfig(ctx context.Context, arg *models.CreateConfigParams) (models.CoreConfigModel, error) {
	return s.Repo.Create(ctx, arg)
}

func (s *Service) DeleteConfig(ctx context.Context, path string) (models.CoreConfigModel, error) {
	return s.Repo.Delete(ctx, path)
}

func (s *Service) UpdateConfig(ctx context.Context, arg *models.UpdateConfigParams) (models.CoreConfigModel, error) {
	return s.Repo.Update(ctx, arg)
}
