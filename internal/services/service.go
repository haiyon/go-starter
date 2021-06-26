package services

import (
	"context"
	repo "haiyon/go-starter/internal/repository"
	"haiyon/go-starter/pkg/conf"
)

// Service services def.
type Service struct {
	c *conf.Config
	r *repo.Repository
}

// New create a Service and return
func New(c *conf.Config) (s *Service) {
	s = &Service{
		c: c,
		r: repo.New(c),
	}
	return
}

// Ping check server
func (s *Service) Ping(ctx context.Context) (err error) {
	err = s.r.Ping(ctx)
	return
}
