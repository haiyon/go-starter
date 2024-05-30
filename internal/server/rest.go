package server

import (
	"go-starter/internal/config"
	"go-starter/internal/handler"
	"go-starter/internal/server/middleware"
	"go-starter/pkg/util"
	"net/http"

	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func registerRest(e *gin.Engine, h *handler.Handler, conf *config.Config) {
	// root Jump when domain is configured and it is not localhost
	e.GET("/", func(ctx *gin.Context) {
		if conf.Domain != "localhost" {
			url := util.GetDomain(conf, conf.Domain)
			ctx.Redirect(http.StatusMovedPermanently, url)
		} else {
			ctx.String(http.StatusOK, "It's working.")
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
