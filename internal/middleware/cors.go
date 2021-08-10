package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CORSHandler 自定义 CORS
func CORSHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// HTTP Headers, ref: https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Headers
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "content-type, content-length, accept-encoding, x-csrf-token, authorization, accept, origin, cache-control, x-requested-with")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "post, get, options, put, delete")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
