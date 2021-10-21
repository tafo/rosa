package api

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/tafo/rosa/internal/auth"
	"github.com/tafo/rosa/internal/middlewares"
	"github.com/tafo/rosa/internal/todo"
	"net/http"
	"time"
)

func NewHttpServer() *http.Server {
	var router = NewRouter()

	public := router.Group("/")
	auth.Handler.MapRoutes(public)

	private := router.Group("/items")
	private.Use(middlewares.AuthHandler())
	todo.Handler.MapRoutes(private)

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
	switch viper.GetString("server_mode") {
	case gin.ReleaseMode:
		gin.SetMode(gin.ReleaseMode)
	case gin.TestMode:
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.DebugMode)
	}

	return gin.Default()
}
