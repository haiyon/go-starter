package server

import (
	"go-starter/internal/config"
	"go-starter/internal/controller"
	"go-starter/internal/data"
	"go-starter/internal/service"
)

// initialize initializes the database, services, and controllers.
func initialize(cfg *config.Config) (*controller.Controller, *service.Service, func(), error) {
	// Initialize database
	d, cleanup, err := data.New(&cfg.Data)
	if err != nil {
		return nil, nil, nil, err
	}

	// Initialize services
	svc := service.New(cfg, d)

	// Initialize controllers
	ctrl := controller.New(svc)

	return ctrl, svc, cleanup, nil
}
