package middleware

import (
	"go-starter/pkg/ecode"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Authorized middleware verifies the existence of a user.
func Authorized(c *gin.Context) {
	// Retrieve user ID from the context, or return unauthorized error
	if userID, exists := c.Get("uid"); !exists || userID == "" {
		// Respond with unauthorized error
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code":    ecode.Unauthorized,
			"message": ecode.Text(ecode.Unauthorized),
		})
		return
	}

	// Continue to the next handler
	c.Next()
}
