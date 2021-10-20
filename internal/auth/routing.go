package auth

import "github.com/gin-gonic/gin"

func (ah AccountHandler) Routes(engine *gin.Engine) {
	engine.POST("/register", ah.register)
}