package server

import (
	"go-starter/internal/config"
	"go-starter/internal/controller"
	"go-starter/internal/server/middleware"
	"go-starter/internal/service"
	"go-starter/pkg/ecode"
	"go-starter/pkg/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	svc     *service.Service
	ctrl    *controller.Controller
	cleanup func()
	err     error
)

// New creates an HTTP server.
func New(cfg *config.Config) (*gin.Engine, func(), error) {
	// Initialize database / services / controllers
	ctrl, svc, cleanup, err = initialize(cfg)
	if err != nil {
		return nil, nil, err
	}

	gin.SetMode(cfg.RunMode)
	engine := gin.New()

	// Middleware
	engine.Use(middleware.Logger())
	engine.Use(middleware.CORSHandler())
	engine.Use(middleware.ConsumeUser())

	// Register REST router
	registerRestRouter(engine, ctrl)

	// Register GraphQL router
	registerGraphqlRouter(engine, cfg.RunMode)

	engine.NoRoute(notFound)
	engine.NoMethod()

	return engine, cleanup, nil
}

func notFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, types.JSON{"message": ecode.Text(http.StatusNotFound)})
}
