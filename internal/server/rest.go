package server

import (
	"github.com/gin-gonic/gin"
	"go-starter/internal/middleware"
)

func innerRest(e *gin.Engine) {
	// Health
	e.GET("/ping", ctrl.Ping)

	// controller prefix
	v1 := e.Group("/")
	// 授权
	v1.GET("/authorize", middleware.Authorized, ctrl.Ping)
	// Hello
	v1.GET("/hello", ctrl.Hello)
}
