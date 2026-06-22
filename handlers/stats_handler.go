package handlers

import "github.com/gin-gonic/gin"

func StatsHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Stats route",
	})
}
