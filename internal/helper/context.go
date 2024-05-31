package helper

import (
	"context"
	"go-starter/internal/config"

	"github.com/gin-gonic/gin"
)

// FromGinContext extracts the context.Context from *gin.Context.
func FromGinContext(c *gin.Context) context.Context {
	return c.Request.Context()
}

// WithGinContext returns a context.Context that embeds the *gin.Context.
func WithGinContext(ctx context.Context, c *gin.Context) context.Context {
	return context.WithValue(ctx, "ginContext", c)
}

// GetGinContext extracts *gin.Context from context.Context if it exists.
func GetGinContext(ctx context.Context) (*gin.Context, bool) {
	c, ok := ctx.Value("ginContext").(*gin.Context)
	return c, ok
}

// GetConfig gets config from gin.Context or context.Context.
func GetConfig(ctx context.Context) *config.Config {
	if c, ok := GetGinContext(ctx); ok {
		if conf, exists := c.Get("config"); exists {
			return conf.(*config.Config)
		}
	}
	if conf, exists := ctx.Value("config").(*config.Config); exists {
		return conf
	}
	// Context does not contain config, load it from config.
	return config.GetConfig()
}

// SetConfig sets config to gin.Context or context.Context.
func SetConfig(ctx context.Context, conf *config.Config) context.Context {
	if c, ok := GetGinContext(ctx); ok {
		c.Set("config", conf)
	}
	return context.WithValue(ctx, "config", conf)
}
