package controller

import (
	"context"
	"go-starter/common/log"
	"go-starter/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Controller REST Struct
type Controller struct {
	s *service.Service
}

// New create a API and return
func New(svc *service.Service) (r *Controller) {
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
