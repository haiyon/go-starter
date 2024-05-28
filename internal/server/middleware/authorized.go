package middleware

import (
	"go-starter/pkg/ecode"
	"go-starter/pkg/resp"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Authorized is a middleware for verifying the existence of a user.
func Authorized(ctx *gin.Context) {
	if _, exists := ctx.Get("uid"); !exists {
		exception := &resp.Exception{
			Status:  http.StatusUnauthorized,
			Code:    ecode.Unauthorized,
			Message: ecode.Text(ecode.Unauthorized),
		}
		resp.Fail(ctx, exception)
		return
	}
	ctx.Next()
}
