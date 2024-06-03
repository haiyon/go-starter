package helper

import (
	"context"
	"go-starter/internal/config"
	"go-starter/pkg/nanoid"

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

// GetValue retrieves a value from the context.
func GetValue(ctx context.Context, key string) any {
	if c, ok := GetGinContext(ctx); ok {
		if val, exists := c.Get(key); exists {
			return val
		}
	}
	if val, exists := ctx.Value(key).(any); exists {
		return val
	}
	return nil
}

// SetValue sets a value to the context.
func SetValue(ctx context.Context, key string, val any) context.Context {
	if c, ok := GetGinContext(ctx); ok {
		c.Set(key, val)
	}
	return context.WithValue(ctx, key, val)
}

// GetConfig gets config from gin.Context or context.Context.
func GetConfig(ctx context.Context) *config.Config {
	if conf, ok := GetValue(ctx, "config").(*config.Config); ok {
		return conf
	}
	// Context does not contain config, load it from config.
	return config.GetConfig()
}

// SetConfig sets config to gin.Context or context.Context.
func SetConfig(ctx context.Context, conf *config.Config) context.Context {
	return SetValue(ctx, "config", conf)
}

// GetTraceID gets trace id from gin.Context
func GetTraceID(ctx context.Context) string {
	if traceID, ok := GetValue(ctx, "trace_id").(string); ok {
		return traceID
	}
	return ""
}

// SetTraceID sets trace id to gin.Context
func SetTraceID(ctx context.Context, traceID string) {
	SetValue(ctx, "trace_id", traceID)
}

// NewTraceID creates a new trace ID.
func NewTraceID() string {
	return nanoid.Must(16)
}
