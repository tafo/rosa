package todo

import "github.com/gin-gonic/gin"

func (ih ItemHandler) MapRoutes(engine *gin.RouterGroup) {
	engine.GET("", ih.GetAll)
	engine.POST("", ih.Create)
}
