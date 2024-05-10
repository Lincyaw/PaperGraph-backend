package handlers

import (
	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello, World!",
	})
}
func IP(c *gin.Context) {
	clientIP := c.ClientIP()
	c.JSON(200, gin.H{
		"ip": clientIP,
	})
}
