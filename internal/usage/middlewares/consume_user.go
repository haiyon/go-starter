package middlewares

import (
	"net/http"
	"haiyon/go-starter/pkg/ecode"
	"haiyon/go-starter/pkg/utils"
	"haiyon/go-starter/pkg/xhttp"
	"strings"

	"github.com/gin-gonic/gin"
)

// refreshToken TODO 刷新 Token
// func refreshToken() {}

// ConsumeUser 处理当前用户
func ConsumeUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("Authorization")
		// Check format
		// ie Bearer: eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9
		b := "Bearer: "
		if !strings.Contains(token, b) {
			ctx.Next()
			return
		}
		t := strings.Split(token, b)
		if len(t) < 2 {
			ctx.Next()
			return
		}
		// 解密令牌
		tokenData, err := utils.DecodeToken(t[1])
		if err != nil {
			exception := &xhttp.ResponseException{
				Status:  http.StatusForbidden,
				Code:    ecode.AccessDenied,
				Message: err.Error(),
			}
			xhttp.Fail(ctx, exception)
			return
		}

		// 设置当前用户 ID
		userID := tokenData["user_id"].(string)
		ctx.Set("uid", userID)
		ctx.Next()

	}
}
