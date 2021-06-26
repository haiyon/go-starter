package http

import (
	"haiyon/go-starter/internal/usage/middlewares"

	"github.com/gin-gonic/gin"
)

func innerRest(e *gin.Engine) {
	// Health
	e.GET("/ping", ctrl.Ping)

	// controllers prefix
	v1 := e.Group("/")
	// 授权
	v1.GET("/authorize", middlewares.Authorized, ctrl.Ping)
	// Hello
	v1.GET("/hello", ctrl.Hello)
}
