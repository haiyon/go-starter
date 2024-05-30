package server

import (
	"go-starter/internal/config"
	"go-starter/internal/handler"
	"go-starter/internal/helper"
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

	// Register REST
	registerRest(engine, h, conf)

	// Register GraphQL
	registerGraphql(engine, svc, conf.RunMode)

	// Register config to Context
	engine.Use(func(c *gin.Context) {
		c.Set("config", conf)
		c.Request = c.Request.WithContext(helper.SetConfig(c.Request.Context(), conf))
		c.Next()
	})

	engine.NoRoute(notFound)
	engine.NoMethod()

	return engine, nil
}

func notFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, types.JSON{"message": ecode.Text(http.StatusNotFound)})
}
