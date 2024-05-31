package middleware

import (
	"go-starter/internal/config"
	"go-starter/internal/helper"

	"github.com/gin-gonic/gin"
)

// BindConfig binds config to gin.Context
func BindConfig(c *gin.Context) {
	ctx := helper.SetConfig(c.Request.Context(), config.GetConfig())
	c.Request = c.Request.WithContext(ctx)
	c.Next()
}

// BindGinContext binds gin.Context to context.Context
func BindGinContext(c *gin.Context) {
	ctx := helper.WithGinContext(c.Request.Context(), c)
	c.Request = c.Request.WithContext(ctx)
	c.Next()
}
