package server

import (
	"go-starter/internal/config"
	"go-starter/internal/handler"
	"go-starter/internal/helper"
	"go-starter/internal/server/middleware"
	"net/http"

	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func registerRest(e *gin.Engine, h *handler.Handler, conf *config.Config) {
	// root Jump when domain is configured and it is not localhost
	e.GET("/", func(c *gin.Context) {
		if conf.Domain != "localhost" {
			url := helper.GetHost(conf, conf.Domain)
			c.Redirect(http.StatusMovedPermanently, url)
		} else {
			c.String(http.StatusOK, "It's working.")
		}
	})

	// Health
	e.GET("/health", h.Ping)

	// api prefix for v1 version
	v1 := e.Group("/v1")

	// Authorize
	v1.GET("/authorize", middleware.Authorized, nil)
	// Hello
	v1.GET("/sample/hello", h.Hello)
	// Swagger
	v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
