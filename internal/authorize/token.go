package authorize

import (
	"errors"
	"go-starter/common/conf"
	"go-starter/common/types"
	"go-starter/common/util"
	"go-starter/internal/generated/ent"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// use config
var signingKey = []byte(conf.G.JWTSecret) // nanoid(35)

// generateToken Generate Token
func generateToken(payload types.JSON, subject string, expire time.Duration) (string, error) {
	atClaims := &jwt.MapClaims{
		"token_id": util.NanoID(),
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

// GenerateUserToken Generate User Token
func GenerateUserToken(user *ent.User, authToken *ent.AuthToken) (string, string) {
	accessPayload := types.JSON{
		"user_id": user.ID,
	}
	refreshPayload := types.JSON{
		"user_id":  user.ID,
		"token_id": authToken.ID,
	}

	accessToken, _ := GenerateAccessToken(accessPayload)
	refreshToken, _ := GenerateRefreshToken(refreshPayload)

	return accessToken, refreshToken
}

// RefreshUserToken Refresh User Token
func RefreshUserToken(user *ent.User, tokenID string, originalRefreshToken string, refreshTokenExp int64) (string, string) {
	now := time.Now().Unix()
	diff := refreshTokenExp - now

	refreshToken := originalRefreshToken
	accessPayload := types.JSON{
		"user_id": user.ID,
	}
	accessToken, _ := GenerateAccessToken(accessPayload)
	if diff < 60*60*24*15 {
		// log.Printf( "refreshing....")
		refreshPayload := types.JSON{
			"user_id":  user.ID,
			"token_id": tokenID,
		}

		refreshToken, _ = GenerateRefreshToken(refreshPayload)
	}

	return accessToken, refreshToken
}

// GenerateAccessToken 生成访问令牌, 360 分钟
func GenerateAccessToken(payload types.JSON, subject ...string) (string, error) {
	defaultSubject := "access_token"
	if len(subject) > 0 {
		defaultSubject = subject[0]
	}
	return generateToken(payload, defaultSubject, time.Hour*6)
}

// GenerateRegisterToken 生成注册令牌, 60 分钟
func GenerateRegisterToken(payload types.JSON, subject ...string) (string, error) {
	defaultSubject := "register_token"
	if len(subject) > 0 {
		defaultSubject = subject[0]
	}
	return generateToken(payload, defaultSubject, time.Hour*1)
}

// GenerateRefreshToken 生成刷新令牌, 10080 分钟, 7 天
func GenerateRefreshToken(payload types.JSON, subject ...string) (string, error) {
	defaultSubject := "refresh_token"
	if len(subject) > 0 {
		defaultSubject = subject[0]
	}
	return generateToken(payload, defaultSubject, time.Hour*24*7)
}
