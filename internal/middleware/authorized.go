package middleware

import (
	"go-starter/common/ecode"
	"go-starter/common/xhttp"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Authorized 验证用户是否存在
func Authorized(ctx *gin.Context) {
	if _, exists := ctx.Get("uid"); !exists {
		exception := &xhttp.ResponseException{
			Status:  http.StatusUnauthorized,
			Code:    ecode.Unauthorized,
			Message: "请求没有得到授权",
		}
		xhttp.Fail(ctx, exception)
		return
	}
	ctx.Next()
}
