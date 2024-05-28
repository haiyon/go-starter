package service

import (
	"go-starter/internal/config"
	"go-starter/internal/data"
)

// Service represents a service definition.
type Service struct {
	cfg    *config.Config
	sample data.ISample
}

// New creates a Service instance and returns it.
func New(cfg *config.Config, d *data.Data) *Service {
	return &Service{
		cfg:    cfg,
		sample: data.NewSample(d),
	}
}
