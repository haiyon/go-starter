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

// newHTTP creates an HTTP server.
func newHTTP(conf *config.Config, h *handler.Handler, svc *service.Service) (*gin.Engine, error) {

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
	registerGraphqlRouter(engine, svc, conf.RunMode)

	engine.NoRoute(notFound)
	engine.NoMethod()

	return engine, nil
}

func notFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, types.JSON{"message": ecode.Text(http.StatusNotFound)})
}
