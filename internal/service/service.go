package service

import (
	"context"
	"go-starter/common/conf"
	"go-starter/internal/generated/ent"
	repo "go-starter/internal/repository"
)

// Service service def.
type Service struct {
	cfg  *conf.Config
	repo *repo.Repository
}

// New create a Service and return
func New(cfg *conf.Config) (s *Service) {
	return &Service{
		cfg:  cfg,
		repo: repo.New(cfg),
	}
}

// DBClose close the resource.
func (svc *Service) DBClose() error {
	return svc.GetClient().Close()
}

// GetClient get entgo client
func (svc Service) GetClient() *ent.Client {
	return svc.repo.Client
}

// Ping check server
func (svc *Service) Ping(ctx context.Context) (err error) {
	err = svc.repo.Ping(ctx)
	return
}
