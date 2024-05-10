package router

import (
	"github.com/Lincyaw/PaperGraph-backend/handlers"
	"github.com/Lincyaw/PaperGraph-backend/middleware"
	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	router := gin.Default()

	// 中间件
	router.Use(middleware.Logging())

	// 路由
	router.GET("/", handlers.Home)
	router.GET("/ip", handlers.IP)
	router.GET("/paper", handlers.Paper)

	return router
}
