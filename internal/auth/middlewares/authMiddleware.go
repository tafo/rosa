package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/tafo/rosa/internal/auth/models"
	"net/http"
	"regexp"
)


type AuthMiddleware struct {
	jwt models.JWTWrapper
}

func NewAuthMiddleware(jwt models.JWTWrapper) AuthMiddleware {
	return AuthMiddleware{jwt}
}

func (m AuthMiddleware) Handler() gin.HandlerFunc {
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

		token := header[7:]
		account, err := m.jwt.ExtractUserFromToken(token)
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