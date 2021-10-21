package auth

import "github.com/gin-gonic/gin"

func (ah AccountHandler) MapRoutes(engine *gin.RouterGroup) {
	engine.POST("/register", ah.Register)
	engine.POST("/login", ah.Login)
}