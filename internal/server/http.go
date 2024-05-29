package server

import (
	"go-starter/internal/config"
	"go-starter/internal/handler"
	"go-starter/internal/server/middleware"
	"go-starter/internal/service"
	"go-starter/pkg/ecode"
	"go-starter/pkg/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	svc     *service.Service
	h       *handler.Handler
	cleanup func()
	err     error
)

// New creates an HTTP server.
func New(conf *config.Config) (*gin.Engine, func(), error) {
	// Initialize database / services / handlers
	h, svc, cleanup, err = initialize(conf)
	if err != nil {
		return nil, nil, err
	}

	gin.SetMode(conf.RunMode)
	engine := gin.New()

	// Middleware
	middleware.Init(conf)
	engine.Use(middleware.Logger())
	engine.Use(middleware.CORSHandler())
	engine.Use(middleware.ConsumeUser())

	// Register REST router
	registerRestRouter(engine, h)

	// Register GraphQL router
	registerGraphqlRouter(engine, conf.RunMode)

	engine.NoRoute(notFound)
	engine.NoMethod()

	return engine, cleanup, nil
}

func notFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, types.JSON{"message": ecode.Text(http.StatusNotFound)})
}
