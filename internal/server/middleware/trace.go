package middleware

import (
	"context"
	"go-starter/internal/helper"
	"go-starter/pkg/log"

	"github.com/gin-gonic/gin"
)

func Trace(c *gin.Context) {
	// Get the trace ID from the request
	traceID := helper.GetTraceID(c)

	// If trace ID is not present in the request, generate a new one
	if traceID == "" {
		traceID = helper.NewTraceID()
		// Set the trace ID in the request context
		ctx := log.NewTraceIDContext(context.Background(), traceID)
		c.Request = c.Request.WithContext(ctx)
	}

	// Set trace header in the response
	c.Writer.Header().Set("X-Trace-ID", traceID)

	c.Next()
}
