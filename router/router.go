package router

import (
	"github.com/Lincyaw/PaperGraph-backend/handlers"
	"github.com/Lincyaw/PaperGraph-backend/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000", "http://localhost:5173"} // 允许来自前端服务器的请求
	// 中间件
	router.Use(middleware.Logging(), cors.New(config))

	api := router.Group("/api/v1")
	api.GET("/paper", handlers.Paper)
	api.GET("/", handlers.Home)
	api.GET("/ip", handlers.IP)
	return router
}
