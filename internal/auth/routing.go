package auth

import "github.com/gin-gonic/gin"

func (ah AccountHandler) MapRoutes(engine *gin.Engine) {
	engine.POST("/register", ah.Register)
}