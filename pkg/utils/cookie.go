package utils

import (
	"github.com/gin-gonic/gin"
)

// SetCookie 设置 Cookies
func SetCookie(ctx *gin.Context, accessToken, refreshToken, domain string) {
	ctx.SetCookie("access_token", accessToken, 60*60*24, "/", domain, true, true)
	ctx.SetCookie("refresh_token", refreshToken, 60*60*24*30, "/", domain, true, true)
}

// SetRegisterCookie 设置注册 Cookies
func SetRegisterCookie(ctx *gin.Context, registerToken, domain string) {
	ctx.SetCookie("register_token", registerToken, 60*60, "/", domain, true, true)
}
