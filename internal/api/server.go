package api

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/tafo/rosa/internal/auth"
	"github.com/tafo/rosa/internal/auth/middlewares"
	"net/http"
	"time"
)

func NewHttpServer(router *gin.Engine, accountHandler auth.AccountHandler, authMiddleware middlewares.AuthMiddleware) *http.Server {
	router.Use(gin.Recovery())
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	accountHandler.Routes(router)
	{
		authorized := router.Group("/")
		authorized.Use(authMiddleware.Handler())
	}
	{
		//admin := router.Group("/")
		//admin.Use(authMiddleware.Handler(), adminMiddleware.Handler())
	}

	port := viper.GetString("server_port")
	if port == "" {
		port = "5001"
	}

	return &http.Server{
		Handler:      router,
		Addr:         ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
}

func NewRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	return gin.New()
}
