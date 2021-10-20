package helper

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
	"github.com/tafo/rosa/internal/auth/models"
)

type HMACSecret []byte

func NewHMACSecret() HMACSecret {
	return []byte(viper.GetString("JWT_SECRET"))
}

type JWTWrapper struct {
	secret HMACSecret
}

func NewJWTWrapper(secret HMACSecret) JWTWrapper {
	return JWTWrapper{secret: secret}
}

func (jwtWrapper JWTWrapper) GenerateTokenForUser(account models.Account) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    account.Id,
		"email": account.Email,
	})

	return token.SignedString([]byte(jwtWrapper.secret))
}

func (jwtWrapper JWTWrapper) ExtractUserFromToken(tokenString string) (account models.Account, err error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method: %v", token.Header["alg"])
		}

		return []byte(jwtWrapper.secret), nil
	})

	if err != nil {
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		account.Id = claims["id"].(int)
		account.Email = claims["email"].(string)
	}
	return
}
