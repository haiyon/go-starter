package helper

import (
	"context"
	"go-starter/internal/config"

	"github.com/gin-gonic/gin"
)

// fromGinContext extracts the context.Context from *gin.Context
func fromGinContext(c *gin.Context) context.Context {
	return c.Request.Context()
}

// withGinContext returns a context.Context that embeds the *gin.Context
func withGinContext(ctx context.Context, c *gin.Context) context.Context {
	return context.WithValue(ctx, "ginContext", c)
}

// getGinContext extracts *gin.Context from context.Context if it exists
func getGinContext(ctx context.Context) (*gin.Context, bool) {
	c, ok := ctx.Value("ginContext").(*gin.Context)
	return c, ok
}

// GetConfig gets config from context
func GetConfig(ctx context.Context) *config.Config {
	if c, ok := getGinContext(ctx); ok {
		if conf, exists := c.Get("config"); exists {
			return conf.(*config.Config)
		}
	}
	if conf, exists := ctx.Value("config").(*config.Config); exists {
		return conf
	}
	return nil
}

// SetConfig sets config to context
func SetConfig(ctx context.Context, conf *config.Config) context.Context {
	if c, ok := getGinContext(ctx); ok {
		c.Set("config", conf)
	}
	return context.WithValue(ctx, "config", conf)
}
