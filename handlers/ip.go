package handlers

import (
	"github.com/gin-gonic/gin"
)

func IP(c *gin.Context) {
	clientIP := c.ClientIP()
	c.JSON(200, gin.H{
		"ip": clientIP,
	})
}
