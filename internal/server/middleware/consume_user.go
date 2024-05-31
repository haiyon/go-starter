package middleware

import (
	"go-starter/pkg/ecode"
	"go-starter/pkg/jwt"
	"go-starter/pkg/resp"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// refreshToken TODO 刷新 Token
func refreshToken(oldToken string) (string, error) {
	return oldToken, nil
}

// isTokenExpiring 检查令牌是否即将过期
func isTokenExpiring(tokenData map[string]any) bool {
	exp, ok := tokenData["exp"].(int64)
	if !ok {
		return false
	}
	expirationTime := time.Unix(exp, 0)
	return time.Until(expirationTime) < 10*time.Minute // 假设如果令牌在 10 分钟内过期，则刷新
}

// ConsumeUser 处理当前用户
func ConsumeUser(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	// Check format
	// ie Bearer: eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9
	b := "Bearer: "
	if !strings.Contains(token, b) {
		c.Next()
		return
	}
	t := strings.Split(token, b)
	if len(t) < 2 {
		c.Next()
		return
	}
	// 解密令牌
	tokenData, err := jwt.DecodeToken(signingKey, t[1])
	if err != nil {
		exception := &resp.Exception{
			Status:  http.StatusForbidden,
			Code:    ecode.AccessDenied,
			Message: err.Error(),
		}
		resp.Fail(c.Writer, exception)
		return
	}

	// 设置当前用户 ID
	userID := tokenData["user_id"].(string)
	c.Set("uid", userID)

	// 检查令牌是否即将过期并刷新令牌
	if isTokenExpiring(tokenData) {
		newToken, err := refreshToken(t[1])
		if err == nil {
			c.Header("Authorization", "Bearer "+newToken)
		}
	}

	c.Next()

}
