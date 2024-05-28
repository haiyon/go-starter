package controller

import (
	"go-starter/internal/service"
)

// Controller represents a controller definition.
type Controller struct {
	Service *service.Service
}

// New creates a new Controller instance.
func New(svc *service.Service) *Controller {
	return &Controller{Service: svc}
}
