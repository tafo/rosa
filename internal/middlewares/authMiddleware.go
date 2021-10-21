package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
	"github.com/tafo/rosa/internal/auth"
	"net/http"
	"regexp"
)


func AuthHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		header := context.GetHeader("Authorization")
		if header == "" {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "You are not authorized.",
				"details": "The 'Authorization' header must be provided",
			})
			return
		}

		if match, err := regexp.MatchString("Bearer .+", header); !match || err != nil {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "You are not authorized.",
				"details": "The 'Authorization' header must in be in the format 'Bearer token'",
			})
			return
		}

		tokenString := header[7:]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("invalid signing method: %v", token.Header["alg"])
			}

			return []byte(viper.GetString("jwt_secret")), nil
		})

		if err != nil {
			return
		}

		var account auth.Account
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			account.Id = claims["id"].(int)
			account.Email = claims["email"].(string)
		}

		if err != nil {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "You are not authorized.",
				"details": "The Bearer token is not valid",
			})
			return
		}

		context.Set("account", account)
		context.Next()
	}
}