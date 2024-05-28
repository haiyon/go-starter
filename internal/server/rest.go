package server

import (
	"go-starter/internal/controller"
	"go-starter/internal/server/middleware"

	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func registerRestRouter(e *gin.Engine, ctrl *controller.Controller) {
	// Controller prefix
	v1 := e.Group("/")
	// Authorize
	v1.GET("/authorize", middleware.Authorized, nil)
	// Hello
	v1.GET("/sample/hello", ctrl.Hello)
	// Swagger
	v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
