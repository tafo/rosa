package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type AdminMiddleware struct{}

func NewAdminMiddleware() AdminMiddleware {
	return AdminMiddleware{}
}

func (m AdminMiddleware) Handler() gin.HandlerFunc {
	return func(context *gin.Context) {
		_, exists := context.Get("account")
		if !exists {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "You are not authorized.",
				"details": "The user must be authenticated",
			})
			return
		}

		context.Next()
	}
}