package util

import (
	"errors"
	"go-starter/common/types"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var signingKey = []byte("YwnIF-8T_Dkfp7dCVlTt0VoBwYM_E31IXh") // utils.DefaultNanoID(34)

// GenerateToken Generate Token
func GenerateToken(payload types.JSON, subject string, expire time.Duration) (string, error) {
	atClaims := &jwt.MapClaims{
		"token_id": NanoID(),
		"subject":  subject,
		"payload":  payload,
		"exp":      time.Now().Add(expire).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		return "", errors.New("令牌签名错误")
	}
	return tokenString, nil
}

// ValidateToken Validate Token
func ValidateToken(t string) (*jwt.Token, error) {
	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		return nil, errors.New("授权令牌验证失败")
	}
	return token, nil
}

// DecodeToken Decode Token
func DecodeToken(t string) (types.JSON, error) {
	token, err := ValidateToken(t)
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("无效的授权令牌")
	}
	return token.Claims.(jwt.MapClaims), nil
}

// GenerateAccessToken 生成访问令牌, 180 分钟
func GenerateAccessToken(payload types.JSON) (string, error) {
	return GenerateToken(payload, "access_token", time.Hour*3)
}

// GenerateRegisterToken 生成注册令牌, 15 分钟
func GenerateRegisterToken(payload types.JSON) (string, error) {
	return GenerateToken(payload, "register_token", time.Minute*15)
}

// GenerateRefreshToken 生成刷新令牌, 7 天
func GenerateRefreshToken(payload types.JSON) (string, error) {
	return GenerateToken(payload, "refresh_token", time.Hour*24*7)
}
