package server

import (
	"go-starter/internal/config"
	"go-starter/internal/data"
	"go-starter/internal/handler"
	"go-starter/internal/service"
)

// initialize initializes the database, services, and handlers.
func initialize(conf *config.Config) (*handler.Handler, *service.Service, func(), error) {
	// Initialize database
	d, cleanup, err := data.New(&conf.Data)
	if err != nil {
		return nil, nil, nil, err
	}

	// Initialize services
	svc := service.New(conf, d)

	// Initialize handlers
	handlers := handler.New(svc)

	return handlers, svc, cleanup, nil
}
