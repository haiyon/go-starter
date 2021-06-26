package controllers

import (
	"context"
	"net/http"
	"haiyon/go-starter/internal/services"
	"haiyon/go-starter/pkg/log"

	"github.com/gin-gonic/gin"
)

// Controller REST Struct
type Controller struct {
	s *services.Service
}

// New create a API and return
func New(svc *services.Service) (r *Controller) {
	return &Controller{
		svc,
	}
}

// Ping health status
func (ctrl *Controller) Ping(ctx *gin.Context) {
	if err := ctrl.s.Ping(ctx); err != nil {
		log.Fatalf(context.Background(), "ping error: %+v", err)
	}
	ctx.Status(http.StatusOK)
}
